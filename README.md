## 介绍
apiBook是接口文档管理工具，私有化部署，本地数据存储，部署和使用门槛低，支持在线请求，mock，兼容各种接口文档导入，多种导出方式，抓取注释生成接口文档等等。

## 支持功能

- 接口文档增删改查，目录归类，多项目接口文档管理，用户管理等基础功能
- markdown说明文档
- json阅读工具（在json中显示字段注释）
- [todo]代码(go, java, c++, php...)的结构体转字段
- 私有化部署 
- 无需依赖三方数据库，本地db文件进行数据存储
- 自动生成代码 
- 文档编辑记录，存档每一次操作
- 加密分享
- 内嵌在线工具(一二三在线工具 https://www.zhaozhongtian.top/)
- [todo]自定义数据实体
- [todo]模拟请求 
- [todo]导入导出 支持OpenAPI,Swagger标准
- [todo]Mock
- [todo]生成pdf,word文件
- [todo]代理功能（http/https）支持http抓包(应用于app开发中的抓包场景)
- [todo]多人修改，多人协作
- [todo]支持 Swagger 注释语法生成接口文档
- [todo]接口测试，自动化测试
- [todo]websocket客户端，调试连接
- [todo]sse客户端，调试连接
- [todo]tcp客户端，调试连接
- [todo]udp客户端，调试连接
- [todo]支持数据库连接进行接口断言测试
- [todo]支持grpc
- 

### 里程碑
- v0.0.1 (20240920) 登录功能，文档与文档目录增删改功能，镜像功能，项目管理功能，用户管理功能，分享功能。
- v0.0.2 生成请求代码功能，文档移动、目录与文档的排序功能。
- v0.0.3 倒排实现, 搜索功能。
- v0.0.4 交互优化，查看项目成员，增加缓存，优化存储，md编辑器上传图片。

### 里程碑 todo
- v0.0.5 导出导入实现，支持json ( OpenAPI 3.1, OpenAPI 3.0, Swagger 2.0 )
- v0.0.6 本地数据库操作工具,系统信息
- v0.0.7 生成pdf,生成word文档
- v0.0.8 导出实现, 这个版本支持 yaml(OpenAPI 3.1, OpenAPI 3.0, Swagger 2.0), markdown
- v0.1.0 所有基础模块开发完成
- v0.1.1 第一个可用版本含核心基础功能

### todo list
- 导入实现 yapi 
- 导出实现 OpenAPI 3.0.1 json
- 导出实现 OpenAPI 3.1 json
- 导出实现 Swagger 2.0 json
- 导出实现 ApiBook 1.0 json
- 导入实现 ApiBook 1.0 json
- 导入界面设计
- 对接导入
- 导出界面设计
- 在开始新增关于入口
- v0.0.5 <发布tag>
- 
- 本地数据库视图操作页(可查看与删除操作 - 管理员可见)
- 系统信息:db文件大小,图片存储大小及数量
- 这个版本功能测试+冒烟测试
- v0.0.6 <发布tag>
- 
- 生成pdf
- 生成word文档
- v0.0.7 <发布tag>
- 
- 导出实现, 这个版本支持  yaml, markdown
- v0.0.8 <发布tag>
- 
- 导入实现 Postman
- 导入实现 HAR
- 导入实现 RAP2
- 导入实现 JMeter
- 导入实现 Eolinker
- 导入实现 NEI
- 导入实现 RAML
- 导入实现 DOClever
- v0.0.9 <发布tag>
- 
- 导入实现 DOCWAY
- 导入实现 ShowDoc
- 导入实现 apiDoc
- 导入实现 I/O Docs
- 导入实现 WADL
- 导入实现 Google Discovery
- v0.0.10 <发布tag>
-
- 代码(go, java, c++, php...)的结构体转字段
- v0.0.11 <发布tag>
-
- 这个版本功能测试+冒烟测试
- 改bug
- v0.1.0 所有基础模块开发完成 <发布tag>
- 
- 冒烟测试
- 优化
- v0.1.1 <发布tag>
- 
- mock实现
- v0.2.1 <发布tag>
- 
- 快速请求页面
- 快速请求业务实现
- v0.3.1 <发布tag>
- 
- 代理功能
- v0.4.1 <发布tag>
- http抓包功能
- v0.4.2 <发布tag>
- websocket客户端，调试连接
- v0.4.3 <发布tag>
- sse客户端，调试连接
- v0.4.4 <发布tag>
- tcp客户端，调试连接
- v0.4.5 <发布tag>
- udp客户端，调试连接
- v0.4.6 <发布tag>
- 
- 站内通知功能,同时修改提示
- v0.5.1 <发布tag>
- 自定义数据实体
- v0.5.2 <发布tag>
-
- 接口测试流程（需要分化）
- v0.6.1 <发布tag>
- 
- 命令行工具抓取 swagger 注释生成接口文档
- v0.7.1 <发布tag>
- 命令行工具抓取注释生成接口文档
- v0.7.2 <发布tag>
- 
- 冒烟测试
- 优化
- v1.0.1 第一阶段完成 <发布tag>
-
- 推广使用
- v1.0.2 <发布tag>

## bug
- [文档页面] 页面刷新请求代码未显示，切换页签才显示
- [搜索] 接口名为:fffff; f进行搜索不到

## 优化


## 预留
支持导入 OpenAPI (Swagger)、Postman、HAR、RAP2、JMeter、YApi、Eolinker、NEI、RAML、DOClever 、Apizza 、DOCWAY、ShowDoc、apiDoc、I/O Docs、WADL、Google Discovery 等数据格式。

