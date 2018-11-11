# go-biz-starter
 golang 下的业务代码种子工程

 ## 这个工程做些什么？
 连接postgres数据库，写逻辑代码。

 ## 怎么使用？
 [TBD]
 ```bash
go mod download

go build ./...

go test ./... -v -cover
```

 ## 特性
 - 数据库集成(postgresl `go-pg`)
 - 数据校验
 - 授权(TBD)
 - 数据库迁移(TBD)
 - 缓存(TBD)

## 依赖
- github.com/go-pg/pg 
- github.com/stretchr/testify 
- go.uber.org/dig 
- gopkg.in/go-playground/validator.v9 

## QA
Q: 没有web接口？<br>
A: 是的，我想把业务代码和web服务器代码分隔。当然，最后肯定需要以某种形式暴露出去的。

## 缩略
为了少敲键盘，节约生命
- biz: bussiness
- mdl: model
- svs: service
- repo: repository,db
- cfg: config
- ctx: context
- utl: util/utility
- vld: validate/validator


## 代码怎么组织的
- 根目录下有一些通用的类型/函数/变量常量
    - types.go, structs
    - funcs.go, global func(s)
    - vars.go global vars

## 贡献代码
嗯，工程主要出于演示目的，可以作为种子使用。展示了我在代码组织、接口设计上的一些想法。

代码没有设计的很灵活，所以用之前可能需要做些改动以满足你的需求，比如：集成mysql或者换一个持久层框架。

喜欢的话start/fork/clone，爱怎么玩怎么玩。

我是golang新手，关于代码如果有什么想法的话请**务必**告诉我->[ [create an issue.](https://github.com/sunsiansong/go-starter-biz/issues/new), PR 也ok.

这些类型的问题应该不会关注
- 增加其他已存在的同类型组件，你完全可以自己搞。比如：增加其他数据库支持。

## TODO
- Integrate some authorization lib,like casbin.
- I18n support.
- better log

## License

许可证： [BSD 3-Clause "New" or "Revised" License](LICENSE).