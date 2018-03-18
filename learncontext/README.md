# golang课程-深入context
----------------------------------
## 1. 课程介绍
学习golang中的context库的使用和原理，
掌握context库的用法，同时加深对golang并发编程的理解

## 2. 课程大纲
- 1. context解决了什么问题
- 1.0 回顾golang并发编程
- 1.1 取消机制：preContext解决方案
- 1.2 看看context需要解决的问题
- 2. 使用context 以gRPC为例
- 2.0 context使用介绍 
- 2.1 作为调用者使用context
- 2.2 作为实现者使用context
- 2.3 context使用规范
- 2.4 几个问题
- 2.4.1 为什么需要“defer cancel()”
- 2.4.2 为什么context没有一个Cancel函数
- 3. 深入context实现
- 4. 总结

## 3. 参考

- [Go语言并发模型：像Unix Pipe那样使用channel](https://segmentfault.com/a/1190000006261218)
- [Go语言并发模型：使用 context](https://segmentfault.com/a/1190000006744213)
- [golang中context包解读](https://studygolang.com/articles/9517)
- [go程序包源码解读——golang.org/x/net/contex](https://studygolang.com/articles/5131)

