golang webFrame

使用内容框架
1)gin框架
2)xorm框架

目录结构介绍

1.config 目录
		1）全局项目启动端口
		2）全局jwt是否开启验证
		3）验证jwt的过期时间
		4）配置是否验证防止sql注入
		5）配置系统数据库相关连接
2.controllers
	书写服务
3.logfile目录
	logfile主要系统运行记录日志
4.middleware 目录
	书写系统重要的中间件插件
5.models
	记录系统所需要的对象模型
6.routers
	记录系统配置的路由信息，具体路由信息已经迁移到Controller层
7.tools/generateCode
	配套系统的代码生成器
8.utils
	公共类库
main.go
程序启动方法

===========未完待续==========



