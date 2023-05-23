# Project Template

这是一个简单的golang web项目的目录结构，内置了一些常用的模块和标准了结构.   
如果你需要开启一个新的web项目，可以clone然后再修改

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
```

## Feature
1. viper (配置)
2. sqlx (数据库)
3. zap (日志)
4. jwt (鉴权)
5. cors (跨域)
6. snowflake(唯一id生成)
7. 一些常用的数据结构

## 如何开始
1. 复制本仓库中的setup_project.sh文件到你本地  
2. chmod +x setup_project.sh
3. ./setup_project.sh  
4. 根据提示操作

## 记录一些常用的包
golang版的loadsh (https://github.com/samber/lo)   
很方便的集合操作库 (https://github.com/deckarep/golang-set)   
微信接口封装 (https://github.com/silenceper/wechat)   

