### 练习项目 go-gin-example

* [源项目地址](https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md)
* [文档地址](https://eddycjy.com/posts/go/gin/2018-02-11-api-01/)

### 项目初始化
- 安装 mysql，新建 blog 数据库，导入 sql 文件，修改数据库配置
- 安装 swagger， `swag init`，更新 swagger 文档
- 启动项目

### CHANGE_LOG

- 22-6-28
  - 添加部分关于 redis 的配置，没有实际使用
  - 提取处理错误和响应方法
  - 重构获取单个文章方法
  - 整理项目结构

- 22-06-27
  - 添加更新时间戳回调
  - 更新 swagger
  - 添加硬删除方法
  - 添加定时任务
  - 重构配置获取方式，统一配置
  - 重构各相关文件获取配置方式，初始化方式
  - 提取操作文件方法，并重构日志相关文件
  - 添加上传图片逻辑
  - 添加获取图片路由
  - 添加 md5 函数
  - 添加新建、修改文章支持 cover_image_url 字段

- 22-06-26
  - 添加 docker 配置（先用镜像内打包，后改为打包后放入镜像内）
  - 安装 docker，熟悉基本操作
    - 启动、停止、删除、打包、链接、查看、数据卷
  - docker 内数据库导入 sql

- 22-06-25
  - 新增标签操作
  - 新增 v1 api
  - 新增 auth 中间件
  - 新增登录方法
  - 新增鉴权方法
  - 添加日志记录
  - 添加热重启，彻底结束进程可以使用 `lsof -i:port`, `kill -3 PID`
  - 新增 swagger

- 22-06-24  
  - 创建基本目录
  - 替换引用本地包 - replace
  - 使用 ini 配置项目
  - 封装数据库基本初始化方式，还未使用
  - 获取初始化 ini 配置
  - 封装全局工具函数
  - 提取测试路由