# gowebframework

从docker提取的web框架

已用其实现完整项目，删除了项目的其他部分，只保留一项account的基本操作作示例。

#作了部分扩展和修改
 ##静态页支持，目前只支持静态文件，需要更详细的支持，可以用go自带的模板。
 ##session 注释掉了，提醒，struct/point需要序列化再保存到session中。
 
dao为和manager为自动生成

封装了有关mongo和redist部分操作

web主要使用gorilla库

库的支持暂不完美

gorilla的schema库 解析提交数据,无法解析json中的arr。需要实现arr解析（或者修改库，因用的地点不多，部分array自已实现,account.roles便为自已解析）
mongodb的mgo库有缺陷，若实体中数据类型为Int(其他nubmer类型)，则不论查询还是更新，若传入的值为0，则无效


go各种库的文档很全面
即使没有文档，从库的test文件可以得到的信息足够满足90%以上的开发需求，还有疑问，可以深入其源码实现，再有疑问google。


tool/code.js 用来解析json数据格式为go的struct。

tool/mongocollection_go_opera_make.cst 根据"mongodb 表名"生成curl+分页的go代码,实现也有很多优化余地。
其他复杂的操作，实现在manager目录下。route-controller-manager/dao

tool/autobuild.js 用nodej和glub-watcher写的监听go 项目自动编译，很鸡肋,go太重了，改一小个部分，就编译的话，多数是编译失败,而且很消耗性能，还没想到完美的优化办法,暂时用个定时器 监听2秒，停止1秒，如此循环，会减少些 “无效”的编译,不过这又有了新的问题,还是不用的好。

-d 设置静态文档目录
-h host
-p port
-f pid路径
-mgop mongodb端口

项目启动
sudo go run ~/src/main/docker
