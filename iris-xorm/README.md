# iris+xorm Go语言开发球星库

https://www.imooc.com/learn/1066

## Install

```bash
$ go get -u github.com/kataras/iris
$ go get -u github.com/go-xorm/xorm
$ go get -u github.com/go-sql-driver/mysql
$ go get -u github.com/gorilla/securecookie
```

## 说明

运行路径在：`src/iris-xorm/web` 下，才能保证 public 和 view 的正常

- 首页： http://localhost:8080
- 管理页面：  http://localhost:8080/admin
    - 用户名：admin
    - 密码：password

## 压力测试

```
$ ab -n10000 -c10 http://localhost:8080/
$ wrk -c10 -t10 -d10 http://localhost:8080/
```
