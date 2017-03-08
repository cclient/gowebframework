# gowebframework

# go语言轻量的web框架,主要部分提取自docker项目

# 之前替代nodejs搭建高并发api server,数据格式为json

web核心为gorilla

beego太重,个人喜欢轻量的东西,也因后台主要为nosql,beego的orm等大量功能,完全用不上,因此采用该方案。

已以该项目为核心实现完整项目，只保留account的基本操作示例,遗憾的是项目完成较早,未采用RESTful

原项目后端mongo+elasticsearch+redis

前端angularjs,现已移除

#支持全局过滤器(有这个很多东西都很好实现了,类似流量统计)

#作了部分扩展和修改
 ##静态页支持，目前只支持静态文件，需要更详细的支持，可以用go自带的模板。
 ##mongo交互
 ##接口分页
 ##session 注释掉了，提醒，struct/point需要序列化再保存到session中。

dao、manager、error、controller、route基础功能,都可用codesmith模板生成,tool下为模板文件

tool/code.js 用来解析json数据格式为go的struct,数据用json交互,存储,根据数据定义,直接生成go类型,支持嵌套。

执行方式(依赖nodejs) node tool/code.js即可
例
var jdata={
    "_id" : "564d5162e54b3106fb7badea",
    "macs" : [
        "00-21-26-00-C8-B0"
    ],
    "time" : 1447907400,
    "timestr" : "2015-11-19 12:30",
    "shop":{
        "name":"shop1"
    }
};

生成结果

type Data struct {
	_id string `json:"_id" bson:"_id"`
	Macs []string `json:"macs" bson:"macs"`
	Time int `json:"time" bson:"time"`
	Timestr string `json:"timestr" bson:"timestr"`
	Shop Shop `json:"shop" bson:"shop"`
}
type Shop struct {
	Name string `json:"name" bson:"name"`
}

根类型,需要手动更改名称,子类型,自动命名。

tool/mongocollection_go_opera_make.cst 根据"mongodb 表名"生成基础的curl+分页的go。

稍复杂的操作，建议实现在manager目录下。

tool/autobuild.js 用nodej和glub-watcher写的监听go 项目自动编译,编译较耗时，改一小个部分，就编译的话，多数是编译失败,而且很消耗性能，还没想到完美的优化办法,暂时用个定时器 监听2秒，停止1秒，如此循环，会减少些 “无效”的编译。

-d 网站根目录对应本地目录路径
-h host
-p port
-f pid路径
-mgop mongodb端口

项目启动 默认端口9900

sudo go run ~/src/main/apiserver

启动成功后

默认主页
curl http://127.0.0.1:9900/

空对象
http://127.0.0.1:9900/api/v1/account/test

异常
http://127.0.0.1:9900/api/v1/account/test2