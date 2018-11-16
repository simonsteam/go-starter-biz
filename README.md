# go-starter-biz
 Business code starter in golang.

[中文readme](README_CN.md)

## Status

In develop, **Not stable!**

## What is this project do?
Connect to postgresql, write your business code.

## How to use?
[TBD]
```bash
go mod download

go build ./...

#create postgresql user
psql < setup.sql
#create test template database
psql < db.sql

go test ./... -v -cover
```

## Features
- Database integration(postgresql `go-pg`)
- Validation
- Authorization (TBD)
- Database migrations.(TBD)
- Cache support.(TBD)

## Dependency
- github.com/go-pg/pg 
- github.com/stretchr/testify 
- go.uber.org/dig 
- gopkg.in/go-playground/validator.v9 

## QA
Q: Without web api here? <br>
A: Yes,i want business code seperate from web server.Of course business code finally should be exposed using rest-api or some other service.

Q: How to 

## Some abbr.
This help understanding codes, and save your life.
- biz: bussiness
- mdl: model
- svs: service
- repo: repository,db
- cfg: config
- ctx: context
- utl: util/utility
- vld: validate/validator


## How files organized
- In the root dir,some basic common files layout here.
    - types.go, structs
    - funcs.go, global func(s)
    - vars.go global vars

## Contribute
Well, this project is mainly for demo purpose,and use as starter.Showing the idea about orgnizing codes, interface design.

Code is not designed flexibly,you may need to do some modifications before start.E.g move from `go-pg` to other orm framework, or support for `mysql`.

If you like it, fork it.Or just clone to your machine and do what you want to it.

Im newbie in golang, if you have any idea about the design,**PLEASE** feel free to [create an issue.](https://github.com/sunsiansong/go-starter-biz/issues/new), PR is also good.

I may not pay attention to these kind of issues:
- Add support for another type of component that already exists, you can do it yourself as you need in your fork/clone.E.g, Add another orm framework(Cause `go-pg` exists in the project)

## TODO
- Integrate some authorization lib,like casbin.
- I18n support.
- better log

## License

This project is licensed under the [BSD 3-Clause "New" or "Revised" License](LICENSE).