#实现登陆与用户模块

## 开发体验
php最爽，session想用就用，不需要将request传来传去，sha1与md5直接调用
nodejs次爽，session使用时需要先传入req，严格的包管理系统，使用很多基础函数都需要先包含才能用，错误处理比较麻烦。
go次一点爽，包管理器没有版本控制，比较坑爹。开发效率上跟nodejs差不多，有了强类型系统，很多错误在编译时就确定下来。但是，编译器对代码格式要求比较多，刚开始比较不习惯（if必须带括号，分行参数末尾要带逗号等等）。另外Orm系统的侵入性很大，需要字段名都变首字母大写了，艹。

## 性能
9线程 php
/login/islogin 299request/s
/user/search 283request/s

1线程 nodejs
/login/islogin 789request/s
/user/search 370request/s

2线程 go
/login/islogin 4500request/s
/user/search 1500request/s

## 总结
go的性能太可怕了，基本上是php与nodejs的十倍左右。
加上严格的代码规范，与类型系统，写出来的代码都是一个样，特别规范。
区分Fatal与Error的异常机制刚开始觉得忧伤，用多了反而会喜欢的。
比较坑爹的是包管理器没有版本控制，而且也没有nodejs的直观。
追求稳定与性能的大团队应该选择go

nodejs的性能还是很好的，至少不会因为io阻塞而导致全站卡死。
有了ES7的async,await，基本上优雅地解决了回调地狱的问题，只是打代码多了很多async与await
丰富而且完善的包管理系统，让人舒服呀。
不区分Fatal与error，出错时统一打印堆栈，心碎了一地。
追求性能的小团队必备

php性能比较忧伤，严重的io阻塞而可能会导致全站卡死。
顺手的标准库函数用到让人爽到爆。
丰富而且完善的包管理系统composer，还是不错的。
不区分Fatal与error，出错时统一打印堆栈，心碎了一地。
追求快速上线的低成本小团队

