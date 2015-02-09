#zBlog.Go

A fast and simple blog engine with [GoInk](https://github.com/fuxiaohei/GoInk) framework in Golang.

zGoBlog is forked from [GoBlog](https://github.com/fuxiaohei/GoBlog)

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

Then it will unzip static files in `new_dir` , initialize raw data and start server at `localhost:9001`.

##### Admin

Visit `localhost:9001/login/` to enter administrator with username `admin` and password `admin`. You'd better change them after installed successfully.

##### Deployment

I prefer to use nginx as proxy. The server section in `nginx.conf`:

        server {
                listen       80;
                server_name  your_domain;
                charset utf-8;
                access_log  /var/log/nginx/your_domain.access.log;

                location / {
                    proxy_pass http://127.0.0.1:9001;
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

### Products

* [Jack.zh's 一声叹息](http://link-pub.cn)

### Thanks

* [@FuXiaoHei](https://github.com/fuxiaohei)

### License

The MIT License

