# simple-tiktok

## 声明 

项目原地址 [地址](https://gitee.com/where-know-return/simple-tiktok)

## 编译运行项目
```shell
#编译
go build main.go router.go
#运行
./main.exe
```

**然后用手机连接电脑热点，手动代理IP即可进行调试** 

## 本地运行注意事项

1.  客户端下载 [地址](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7) 
2. PC端需要有 Mysql数据库，且已开启服务 （ps: 默认数据库账户和密码均为 root，如需个更改，请在db_connect .go 文件中修改）
3. 手机客户端与服务器通信办法（视频流需要的带宽挺大的，建议本地调试）[参考办法](https://juejin.cn/post/7096857967747661831)

## 鸣谢 

[字节跳动后端青训营](https://youthcamp.bytedance.com/) 