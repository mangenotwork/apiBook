# apiBook
接口文档管理工具，私有化部署，本地数据存储，无需数据库。

### 介绍
私有化部署接口文档管理工具，todo....

## todo list
- 修改bug
- v0.0.1
- 文档移动
- 目录排序
- 文档排序
- 倒排索引实现
- 搜索实现与对接
- 这个版本功能测试+冒烟测试
- v0.0.2
- 请求代码
- 这个版本功能测试+冒烟测试
- v0.0.3
- md编辑器上传图片实现
- 这个版本功能测试+冒烟测试
- v0.0.4
- 导入实现
- 这个版本功能测试+冒烟测试
- v0.1.0 所有基础模块开发完成
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


## bug

- [功能测试][首页] 首页分享没反应
- [功能测试][创建项目] 创建一个曾经的项目名字失败，提示项目名已存在
- [功能测试][创建用户] 创建一个曾经的用户账号或用户名创建成功了，然后把曾经的账户信息修改了
- [功能测试][文档页面] 参数说明表格示例字段字符长了没用弹出显示
- [功能测试][文档页面] http://127.0.0.1:17777/index/7b9cc815954c1f8b 响应显示错误，应该显示txt文本的信息，表格与参数说明一样
- [功能测试][文档页面] 类型为文本类型，没法显示
- [功能测试][文档页面] 镜像页面没有显示日志&镜像的数据
- [功能测试][文档页面] 多文档页签切换，json会留存上一个文档的内容
- [功能测试][分享] 分享项目首次打开是空白页面，需要有提示信息
- [功能测试][分享] 分享文档没法打开日志&镜像
- [功能测试][分享] 分享文档开日志&镜像左边没有竖杠
- [功能测试][分享] 分享项目右边锚点没有背景样式
- [功能测试][文档页面] 请求数据类型需要显示为 Content-Type:application/json; charset=utf-8 这种
- [代码走查][]