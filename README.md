# Go-Vue-Blog
A Blog System Project base on Go (backend) and Vue.js ( frontend )  

一个基于 Go 的 Gin、GORM 和 JS 的 Vue 2、Ant Design Vue 的全栈在线博客项目。旨在快速熟悉 Go 的工程化实践及相关技术栈的理解和运用。

### 一、项目结构说明

1. `config/`   
   保存项目的配置信息。  
   

2. `model/`   
   编写 Model 层，是系统的数据库模型。在此使用了 GORM 作为对象关系映射。
   

3. `api/`  
   编写 Controller 层，并针对可能的版本更新设置`v1/`、`v2/`...等子文件夹。  
   

4. `router/`  
    路由层，来自前端的请求在此被分配对应的 API 去执行。  
   
   
5. `middleware/`  
    中间件层，在此处实现跨域拦截、日志记录等中间件。  
   
   
6. `utils/`  
    工具类库，在此实现一些通用的工具。  
   
   
7. `upload/`  
    调用 七牛云存储 API ,实现图床功能。  
   
   
8. `web/`  
    前端 Vue 项目，实现前后端分离。