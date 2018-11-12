More QA
=====

Q: How do you deal with database in testing? <br>
A: There are stories.

Basicly,use dependency injection.`pg.DB` is injected to component when needed.

Too keep single test case env simple and clean, every unit test running in seperate database.

First, i considered using docker to create database per test case,but it is slow,i can not wait for that,and introducing docker also increse complexcity.

Then, i try to create a new database in the beginning of test, and create tables, and prepare datas, finally run test logic.It works, but i want to make it more simple.

Next, i try to use postgresql template database(i create a template database with preset tables) to create new test database,it works fine in daily test.I think i come to the final solution at first.When i runs `go test ./...` in the workspace root directory, tests `FAIL` with error: 
```
"source database \"biz_test_template\" is being accessed by other users", 0x44:"There is 1 other session using the database."
```
yes, i find that in postgresql document(chapter: managing database/template databases):
>It is possible to create additional template databases, and indeed one can copy any database in a cluster by specifying its name as the template for CREATE DATABASE. It is important to understand, how- ever, that this is not (yet) intended as a general-purpose “COPY DATABASE” facility. The principal limitation is that no other sessions can be connected to the source database while it is being copied. CREATE DATABASE will fail if any other connection exists when it starts; during the copy operation, new connections to the source database are prevented.

Also, refer to https://stackoverflow.com/questions/876522/creating-a-copy-of-a-database-in-postgresql

Fine,(´・＿・\`),`pg_dump` and `psql` would works,but it may sucks in windows.

**Finally**, i come back to exec create table sql line by line,several lines of code.

I also considerd using database migrate tools to construct new database,the solution sounds good,but i can not find a lib simple enugh(Especially, some lib rely on much dependencies,i feel bad). Though, execute some sqls in code is easy, i dont want hard code create table sql in go files, or parse a sql file and execute line by line.

Im not skilled at mock, if somebody has good idea,please tell me, thx.


-----

Q: Why you design file structure like this, any consideration?<br>
A: (TBD)