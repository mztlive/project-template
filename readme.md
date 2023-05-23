# Project Template

可以基于这个模板创建基本的项目结构，内置以下模块

1. viper (配置)
2. sqlx (数据库)
3. zap (日志)
4. jwt (鉴权)
5. cors (跨域)
6. snowflake(唯一id生成)
7. 一些常用的数据结构


```sh
工程根目录
.
├── Makefile    # 编译脚本
|-- bin         # 编译后的产物
├── cmd         # 项目的入口文件
├── internal    # 内部代码，比如说一些service
├── pkg         # 一些公共代码，随模板一起发布的
├── http        # web项目的逻辑入口，控制器、中间件、路由都在这里面定义
├── go.mod
├── go.sum
├── readme.md
.
.
.
```