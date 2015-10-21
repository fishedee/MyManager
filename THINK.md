#实现登陆与用户模块

## 开发体验
php最爽，session想用就用，不需要将request传来传去，sha1与md5直接调用
nodejs次爽，session使用时需要先传入req，严格的包管理系统，使用很多基础函数都需要先包含才能用，错误处理比较麻烦。

## 性能
9线程 php
/login/islogin 299request/s
/user/search 283request/s

1线程 nodejs
/login/islogin 789request/s
/user/search 370request/s

4线程 go

## 总结
