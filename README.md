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
├── app-svr     // 应用接口服务
│   ├── app-admin   // 管理平台接口服务
│   ├── app-interface   // app接口服务
│   ├──
│
├── errno   // error number 定义
│
├── micro-svr   // 微服务
│   ├── oauth2  // OAuth2 服务
│   ├──
│   ├──
│
├── pkg     // 公共包
│   ├── errgroup    // WaiGroup的封装包
│   ├── orm     // SQL 初始化
│   ├── time    // 自定义 time
│   ├──
│   ├──
│
├── proto   // proto 文件
│   ├── oauth2
│
│
├── .gitignore  // 忽略文件
│
├── CHANGELOG.md    // change log
│   
├── go.mod      // mod 包管理文件
│    └── go.sum
└── README.md
```