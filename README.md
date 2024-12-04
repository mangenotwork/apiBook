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
- v0.0.5 导出导入实现（全量），支持json ( OpenAPI 3.1, OpenAPI 3.0, Swagger 2.0 )
- v0.0.6 本地数据库操作工具,系统信息
- 
### 里程碑 todo
- v0.0.7 现有功能的一些优化
- v0.1.0 所有基础模块开发完成
- v0.1.1 第一个可用版本含核心基础功能
- v0.2.4 在线请求功能
- v0.3.-  mock功能

### todo list
- [UI] 标准化设计和调优
- 在项目下拉菜单新增一个切换项目的选项  
- 页面刷新请求代码单独请求接口
- 导入提交后按钮置灰
- 导入增加进度条
- 导入可选择私有还是公有
- [UI] 目录的操作按钮收起来，不单独显示一行， 目录字小两号
- 搜索权重规划一下
- v0.0.7 <发布tag>

- 
- 这个版本功能测试+冒烟测试
- 改bug
- v0.1.0 所有基础模块开发完成 <发布tag>
- 
- 冒烟测试
- 优化
- v0.1.1 <发布tag>
- 
- 在线请求实现,暂定方案 l4进行转发 或 http请求器进行请求
- v0.2.1 <发布tag>
- 
- 在线请求页面
- 在线请求业务实现
- v0.2.2 <发布tag>
-
- 文档页面与在线请求衔接
- 将在线请求的内容回写到文档
- v0.2.3 <发布tag>
-
- 通过在线请求创建文档
- v0.2.4 <发布tag>


- mock实现
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
- [分享]打开分享文档第一白屏
```azure
Uncaught TypeError: $(...).tab is not a function
    at loadReqCode (doc.js?v=1732500772:215:29)
    at tnWiMry:2166:17
    at Object.success (pub.js?v=1732500772:330:13)
    at l (jquery.min.js?v=1732500772:4:24584)
    at Object.fireWith [as resolveWith] (jquery.min.js?v=1732500772:4:25405)
    at k (jquery.min.js?v=1732500772:6:4694)
    at XMLHttpRequest.<anonymous> (jquery.min.js?v=1732500772:6:8498)
    at Object.send (jquery.min.js?v=1732500772:6:8690)
    at Function.ajax (jquery.min.js?v=1732500772:6:4167)
    at AjaxPostNotAsync (pub.js?v=1732500772:323:7)
```
- [导出]swagger请求内容不全
- 

## 优化

- 导入需要对md格式转html，引入md库
- 搜索支持检索式（and  or  not  ..... ） <需要调研>
- 每个文档增加 md5, 用于导入时间检查重复

## 需求池
- 导入实现 Postman
- 导入实现 HAR
- 导入实现 RAP2
- 导入实现 JMeter
- 导入实现 Eolinker
- 导入实现 NEI
- 导入实现 RAML
- 导入实现 DOClever
- 导入实现 DOCWAY
- 导入实现 ShowDoc
- 导入实现 apiDoc
- 导入实现 I/O Docs
- 导入实现 WADL
- 导入实现 Google Discovery
- 生成pdf
- 生成word文档
- 导出实现, 这个版本支持  yaml, markdown
- 代码(go, java, c++, php...)的结构体转字段


## 预留


## debug
- /debug/sysInfo    查看系统信息-总览  项目总数量，用户总数量，db文件大小,图片存储大小及数量，运行时间
- /debug/projectInfo   查看系统信息-项目  项目信息，接口数量，用户数量，操作日志   
```azure
参数
type : list（查列表）  ""(查详情)
pid : 项目id, 查详情不能为空
```
- /debug/sysLog    查看系统日志  (登录，操作，报错)
- /debug/conf    查看配置文件
- /debug/db/search/bucket    搜索 db 的bucket
```azure
参数
dbName : ""(默认)  invertIndexDB(倒排索引)
search :  搜索词
```
- /debug/db/select    查看 db 对应的key的数据
```azure
参数
dbName: ""(默认)  invertIndexDB(倒排索引)
bucket: 桶
selectType： 查询类型   allKey（获取所有的key）  searchKey（搜索key） ""(搜索数据)
key： 搜索key
search: 搜索词
```


