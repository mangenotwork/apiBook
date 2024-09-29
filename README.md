# apiBook
接口文档管理工具，私有化部署，本地数据存储，无需数据库。

### 介绍
私有化部署接口文档管理工具，todo....

### 里程碑
- v0.0.1 (20240920) 登录功能，文档与文档目录增删改功能，镜像功能，项目管理功能，用户管理功能，分享功能。
- v0.0.2 生成请求代码功能，文档移动、目录与文档的排序功能。
- v0.0.3 倒排实现, 搜索功能，
- v0.0.4 导出导入实现，支持json ( OpenAPI 3.1, OpenAPI 3.0, Swagger 2.0 )
- v0.0.5 交互优化，查看项目成员，增加缓存，md编辑器上传图片。
- v0.0.6 本地数据库操作工具
- v0.0.7 生成pdf,生成word文档
- v0.0.8 导出实现, 这个版本支持 yaml(OpenAPI 3.1, OpenAPI 3.0, Swagger 2.0), markdown
- v0.1.0 所有基础模块开发完成
- v0.1.1 第一个可用版本含核心基础功能

### todo list
- 文档详情页增加复制url功能
- 私有项目需要查看项目成员
- 增加内存缓存提示接口响应性能
- md编辑器上传图片实现
- 这个版本功能测试+冒烟测试
- v0.0.4 <发布tag>
- 导入实现,这个版本支持json基础的,如yapi，OpenAPI 3.1, OpenAPI 3.0, Swagger 2.0
- 导出实现,这个版本支持json基础的导出: OpenAPI 3.1, OpenAPI 3.0, Swagger 2.0
- v0.0.5 <发布tag>
- 本地数据库视图操作页(可查看与删除操作 - 管理员可见)
- 这个版本功能测试+冒烟测试
- v0.0.6 <发布tag>
- 生成pdf
- 生成word文档
- v0.0.7 <发布tag>
- 导出实现, 这个版本支持  yaml, markdown
- v0.0.8 <发布tag>
- 改bug
- v0.1.0 所有基础模块开发完成 <发布tag>
- 冒烟测试
- 优化
- v0.1.1 <发布tag>
- mock实现
- v0.2.1 <发布tag>
- 快速请求页面
- 快速请求业务实现
- v0.3.1 <发布tag>
- 代理功能
- http抓包功能
- v0.4.1 <发布tag>
- 站内通知功能,同时修改提示
- v0.5.1 <发布tag>
- 冒烟测试
- 优化
- v1.0.1 第一阶段完成 <发布tag>
- 推广使用
- v1.0.2 <发布tag>
- 接口测试流程（需要分化）
- v1.0.3 <发布tag>
- 支持 swagger 并对接 swagger

## bug
- 

## 优化
- [数据] 将倒排索引数据从业务数据中拆分出来，新的db
- [文档页面] 创建了文档成功后应该默认打开
- [文档页面] 从目录选择创建文档 目录需要一样
