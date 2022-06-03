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
- **钟昊** 点赞列表， _对应文件_
  
  > favorite.go
- **张文鋆**  粉丝列表，*对应文件*
    
    > relation.go
- **聂慈** 关注列表， *对应文件*
  
  > relation.go

- **王智宇** 评论列表， *对应文件*

  > comment.go

![结构](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/434/202206032317915.png)