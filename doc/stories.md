story\故事
==========

### about module \ 关于模块
I used to designed dependency for biz.Module,thus i just need specify direct module used then the system works, just like : `NewEnv(XXXModule)`, but then i got dependency hell :-D, so now use something like `NewEnv(AModule, BModule, CModule, FModule)`,verbose but simple.

zh_CN:我原来给模块设计了依赖系统，后来搞成依赖地狱了（笑，就现在这样也好，冗长但简单。


---

### Access Control \ 访问控制
I tried to use casbin,it's not hard to understand, but not that easy to use.The role based access control model is easy. So,several lines of code werer writen to to that.

zh_CN: 原本想用casbin的，好理解但觉得不好用。基于角色权限的模型其实挺容易理解的，所以就自己写了几行代码实现。


---