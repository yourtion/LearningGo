# Go语言实战抽奖系统

https://coding.imooc.com/class/295.html

## Install

```bash
$ go get -u github.com/kataras/iris
$ go get -u github.com/iris-contrib/httpexpect
$ go get -u github.com/go-xorm/xorm
$ go get -u github.com/go-sql-driver/mysql
$ go get -u github.com/gorilla/securecookie

$ go get -u github.com/go-xorm/cmd/xorm
```

```
$ xorm reverse mysql "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8"  ./src/github.com/go-xorm/cmd/xorm/templates/goxorm
```

## _demos

- 年会抽奖程序：[1-annual-meeting](./_demo/1-annual-meeting/main.go)
- 彩票：[2-ticket](./_demo/2-ticket/main.go)
- 微信摇一摇：[3-wechat-shake](./_demo/3-wechat-shake/main.go)
- 支付宝集福卡：[4-alipay-fu](./_demo/4-alipay-fu/main.go)
- 微博抢红包：[5-weibo-redpacket](./_demo/5-weibo-redpacket/main.go)
- 抽奖大转盘：[6-wheel](./_demo/6-wheel/main.go)

### 需求整理提炼

#### 前端页面

- 交互效果、大转盘展示
- 用户登录，每日抽奖次数限制
- 获奖提示和中奖列表

#### 后端需求

- 登录、退出、奖品列表、抽奖、中奖列表
- 抽奖接口需要满足**高性能**和**高并发**要求
- 安全抽奖，奖品**不能超发**、合理均匀发放

#### 后台管理需求

- 基本数据管理：奖品、优惠券、用户、IP黑名单、中奖记录
- 试试更新奖品信息，更新奖品库存，奖品中奖周期等
- 后台定时任务，生成发奖计划，填充奖品池信息

#### 用户操作

登录 -> 抽奖页面（奖品列表、中奖记录、剩余抽奖次数）-> 抽奖（参数验证、中奖）-> 发奖（中奖记录、中奖提示）

#### 奖品状态变化

正常奖品 ->（奖品有效期、奖品库存）-> 奖品池

#### 抽奖业务流程

抽奖 -> 验证 -> 用户抽奖锁定 -> 验证用户今日参与次数 -> 验证IP今日最大限制次数 -> 验证IP今日抽奖次数 -> 验证IP黑名单 -> 验证用户黑名单 -> 获取抽奖编码 -> 匹配奖品 -> 验证奖品池库存 -> 发奖，更新奖品池库存 -> 发放优惠券 -> 中奖记录 -> 返回结果


#### 如何设计和利用缓存

- 目标：提高系统**性能**，减少数据库依赖
- 原则：**平衡**好“系统性能、开发时间、复杂度”
- 方向：数据**读多写少**，数据量**有限**，数据**分散**

#### 使用缓存的地方

- 奖品：数量少，更新频率低，最佳的全局缓存对象
- 优惠券：一次导入，优惠券编码缓存为 set 类型
- 中奖记录：读写差不多，可以缓存部分统计数据，如：最新中奖记录、最佳大奖发放记录等
- 用户黑名单：读多写少，可以按照 uid 散列
- IP黑名单：类似用户黑名单，可以按照 ip 散列

#### 系统设计与架构设计

- 需求：充分考虑运营需要，用户操作尽量简单
- 数据库设计：定义数据模型简单够用，留下扩展空间
- 架构设计：分层架构设计
    - 网络：负载均衡层/应用层/存储层
    - 应用：业务代码/框架代码/存储服务
