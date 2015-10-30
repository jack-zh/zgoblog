package handler

import (
	"github.com/jack-zh/zgoblog/app/model"
	"github.com/jack-zh/zgoblog/app/utils"
	"github.com/jack-zh/zgoblog/fweb"
	"net/url"
	"strconv"
	"strings"
)

func Login(context *fweb.Context) {
	if context.Method == "POST" {
		data := context.Input()
		user := model.GetUserByName(data["user"])
		if user == nil {
			Json(context, false).End()
			return
		}
		if !user.CheckPassword(data["password"]) {
			Json(context, false).End()
			return
		}
		exp := 3600 * 24 * 3
		expStr := strconv.Itoa(exp)
		s := model.CreateToken(user, context, int64(exp))
		context.Cookie("token-user", strconv.Itoa(s.UserId), expStr)
		context.Cookie("token-value", s.Value, expStr)
		Json(context, true).End()
		return
	}
	if context.Cookie("token-value") != "" {
		context.Redirect("/admin/")
		return
	}
	context.Render("admin/login", nil)
}

func Auth(context *fweb.Context) {
	tokenValue := context.Cookie("token-value")
	token := model.GetTokenByValue(tokenValue)
	if token == nil {
		context.Redirect("/logout/")
		context.End()
		return
	}
	if !token.IsValid() {
		context.Redirect("/logout/")
		context.End()
		return
	}
}

func Logout(context *fweb.Context) {
	context.Cookie("token-user", "", "-3600")
	context.Cookie("token-value", "", "-3600")
	context.Redirect("/login/")
}

func Geek(context *fweb.Context) {
	context.Layout("geek")
	Theme(context).Layout("geek").Render("geek", nil)
}

func TagArticles(ctx *fweb.Context) {
	ctx.Layout("home")
	page, _ := strconv.Atoi(ctx.Param("page"))
	tag, _ := url.QueryUnescape(ctx.Param("tag"))
	size := getArticleListSize()
	articles, pager := model.GetTaggedArticleList(tag, page, getArticleListSize())
	// fix dotted tag
	if len(articles) < 1 && strings.Contains(tag, "-") {
		articles, pager = model.GetTaggedArticleList(strings.Replace(tag, "-", ".", -1), page, size)
	}
	Theme(ctx).Layout("home").Render("index", map[string]interface{}{
		"Articles":    articles,
		"Pager":       pager,
		"SidebarHtml": SidebarHtml(ctx),
		"Tag":         tag,
		"Title":       tag,
	})
}

func Home(context *fweb.Context) {
	context.Layout("home")
	page, _ := strconv.Atoi(context.Param("page"))
	articles, pager := model.GetPublishArticleList(page, getArticleListSize())
	data := map[string]interface{}{
		"Articles":    articles,
		"Pager":       pager,
		"SidebarHtml": SidebarHtml(context),
	}
	if page > 1 {
		data["Title"] = "第 " + strconv.Itoa(page) + " 页"
	}
	Theme(context).Layout("home").Render("index", data)
}

func ListArticles(context *fweb.Context) {
	articles, pager := model.GetArticleList(context.Int("page"), 10)
	Theme(context).Layout("home").Render("listarticles", map[string]interface{}{
		"Title":    "文章列表",
		"Articles": articles,
		"Pager":    pager,
	})
}

func Article(context *fweb.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	slug := context.Param("slug")
	article := model.GetContentById(id)
	if article == nil {
		context.Redirect("/")
		return
	}
	if article.Slug != slug || article.Type != "article" {
		context.Redirect("/")
		return
	}
	article.Hits++
	Theme(context).Layout("home").Render("article", map[string]interface{}{
		"Title":       article.Title,
		"Article":     article,
		"CommentHtml": CommentHtml(context, article),
	})
}

func Page(context *fweb.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	slug := context.Param("slug")
	article := model.GetContentById(id)
	if article == nil || article.Status != "publish" {
		context.Redirect("/")
		return
	}
	if article.Slug != slug || article.Type != "page" {
		context.Redirect("/")
		return
	}
	article.Hits++
	Theme(context).Layout("home").Render("page", map[string]interface{}{
		"Title": article.Title,
		"Page":  article,
	})
}

func TopPage(context *fweb.Context) {
	slug := context.Param("slug")
	page := model.GetContentBySlug(slug)
	if page == nil || page.Status != "publish" {
		context.Redirect("/")
		return
	}
	if page.IsLinked && page.Type == "page" {
		Theme(context).Layout("home").Render("page", map[string]interface{}{
			"Title": page.Title,
			"Page":  page,
		})
		page.Hits++
		return
	}
	context.Redirect("/")
}

func Comment(context *fweb.Context) {
	cid, _ := strconv.Atoi(context.Param("id"))
	if cid < 1 {
		Json(context, false).End()
		return
	}
	if model.GetContentById(cid) == nil {
		Json(context, false).End()
		return
	}
	data := context.Input()
	msg := validateComment(data)
	if msg != "" {
		Json(context, false).Set("msg", msg).End()
		return
	}
	co := new(model.Comment)
	co.Author = data["user"]
	co.Email = data["email"]
	co.Url = data["url"]
	co.Content = data["content"]
	co.Avatar = utils.Gravatar(co.Email, "50")
	co.Pid, _ = strconv.Atoi(data["pid"])
	co.Ip = context.Ip
	co.UserAgent = context.UserAgent
	co.IsAdmin = false
	model.CreateComment(cid, co)
	Json(context, true).Set("comment", co.ToJson()).End()
	model.CreateMessage("comment", co)
	context.Do("comment_created", co)
}

func validateComment(data map[string]string) string {
	if utils.IsEmptyString(data["user"]) || utils.IsEmptyString(data["content"]) {
		return "称呼，邮箱，内容必填"
	}
	if !utils.IsEmail(data["email"]) {
		return "邮箱格式错误"
	}
	if !utils.IsEmptyString(data["url"]) && !utils.IsURL(data["url"]) {
		return "网址格式错误"
	}
	return ""
}

func getArticleListSize() int {
	size, _ := strconv.Atoi(model.GetSetting("article_size"))
	if size < 1 {
		size = 5
	}
	return size
}
