#zBlog.Go

zGoBlog is forked and edit from [GoBlog](https://github.com/fuxiaohei/GoBlog)

### Installation

`zBlog.Go` requires **Go 1.2** or above.

##### Manual

Use go get command:

    go get github.com/jack-zh/zGoBlog

Then you can find binary file `GoBlog(.exe)` in `$GOPATH/bin`.

### Run

Make a new dir to run `zBlog.Go`:

    cd new_dir
    zGoBlog

Then it will unzip static files in `new_dir` , initialize raw data and start server at `localhost:8888`.

##### Admin

Visit `localhost:8888/login/` to enter administrator with username `admin` and password `admin`. You'd better change them after installed successfully.

##### Deployment

I prefer to use nginx as proxy. The server section in `nginx.conf`:

        server {
                listen       80;
                server_name  your_domain;
                charset utf-8;
                access_log  /var/log/nginx/your_domain.access.log;

                location / {
                    proxy_pass http://127.0.0.1:8888;
                }

                location /static {
                    root            /var/www/your_domain;  # binary file is in this directory
                    expires         1d;
                    add_header      Cache-Control public;
                    access_log      off;
                }
        }

### Questions

Create issues or pull requests here.

### Note

You can edit `static/hsterm/config.js` to set your geek message from menu.

### Products

* [Jack.zh's 一声叹息](http://link-pub.cn)

### Thanks

* [@FuXiaoHei](https://github.com/fuxiaohei)

### Version
    0.3

### License

The MIT License
