# goMonitor
analyse data from log and show in charts
## 实现
### 后端
go+gin框架
### 前端
html+原生js
## 需要的环境
### go
参考链接[ubuntu1604安装go1.13](https://blog.csdn.net/u013536232/article/details/104124423)
### gin
参考链接[go 安装gin(使用git clone)](https://blog.csdn.net/u013536232/article/details/104476356)
### crontab
如果需要进程CPU和内存监控，则需要系统开启crontab服务
## 部署方式
### 修改ip
`sh setIp.sh your_server_ip`
### 链接日志文件夹
`ln -s your_log_path log`
### 编译运行
`sh build.sh`<br>
`./monitor`