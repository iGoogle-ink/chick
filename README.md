## chick

### 技术栈

- [go-micro](https://github.com/micro/go-micro)
- [go-plugins](https://github.com/micro/go-plugins)
- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/jinzhu/gorm)
- [go-redis](https://github.com/go-redis/redis)
- [viper](https://github.com/spf13/viper)
- [errors](https://github.com/pkg/errors)


### 结构介绍
```
├── api   // proto service
│   ├── oauth2
│   ├── user
│   ├── 
│
├── app     // 业务网关服务
│   ├── account   // 账号相关
│   ├── app-svr   // app相关
│   ├── web-svr   // web相关
│
├── errno   // error number 定义
│
├── micro-svr   // 微服务
│   ├── user    // 用户相关
│   ├──
│
├── pkg     // 公共包
│   ├── config  // 公共config
│   ├── errgroup    // WaiGroup的封装包
│   ├── orm     // orm 初始化
│   ├── server  // micro 初始化 
│   ├── time    // 自定义 time
│   ├── util    // 基础工具
│   ├── web     // web 初始化
│
├── .gitignore  // 忽略文件
│
├── CHANGELOG.md    // change log
│   
├── go.mod      // mod 包管理文件
│    └── go.sum
└── README.md
```