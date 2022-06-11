# simple-tiktok

## 编译运行项目
```shell
#编译
go build main.go router.go
#运行
./main.exe
```

**然后用手机连接电脑热点，手动代理IP即可进行调试**

## 项目分工
- **侯凯恒** 视频Feed流，视频投稿信息，个人信息，
  *对应文件*

  > feed.go  publish.go user.go

- **张文鋆**  粉丝列表，关注列表，*对应文件*
    
    > relation.go
- **聂慈** 点赞列表，评论列表 *对应文件*
   
   >  favorite.go comment.go
## 数据库
项目使用 MySQL数据库
本地调试请修改 **db_connect.go** 文件的用户名和密码

文件中的用户名和密码均为 root