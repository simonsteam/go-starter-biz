story\故事
==========

### about module \ 关于模块
I used to designed dependency for biz.Module,thus i just need specify direct module used then the system works, just like : `NewEnv(XXXModule)`, but then i got dependency hell :-D, so now use something like `NewEnv(AModule, BModule, CModule, FModule)`,verbose but simple.

zh_CN:我原来给模块设计了依赖系统，后来搞成依赖地狱了（笑，就现在这样也好，冗长但简单。