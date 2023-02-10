# go_rent

# 租房小程序服务端

使用框架: `beego`

项目结构：

<img width="1100" alt="image" src="https://user-images.githubusercontent.com/38744096/218026368-264c75e5-930d-4af3-81d9-9b3ccd822034.png">


# 小程序端：
https://github.com/RelaxedDong/rent_mini

# 配置文件
```
appname = 蚁租房
httpport = 8000
runmode = dev
copyrequestbody = true
logfilepath=logs/service.log

oss_access_key=xxx
oss_secret_key=xxxx
oss_upload_path=https://bucket_name.oss-cn-shenzhen.aliyuncs.com
oss_region_host=https://bucket_name.oss-cn-shenzhen.aliyuncs.com


db_port=3306
db_database=xxx

[dev]
httpaddr = 127.0.0.1
APPID=xxx # 小程序appid
APP_SECRET=xxx # 小程序 secret
db_host=xxx
db_user=xxx
db_password=xxx


[pro]
httpaddr = 127.0.0.1
APPID=xxx
APP_SECRET=xxx
db_host=localhost
db_user=xxx
db_password=xxx
```

# 部署
1. 使用bee打包：`bee pack`
2. 解压压缩包： `tar -zxvf go_rent.tar.gz`
3. 使用supervisor管理进程
```
[program:rent_backend]
#supervisor执行命令
command = /home/go_rent/go_rent
#项目目录
derectory = /home/go_rent
#开始的时候等待多少秒
startsecs = 0
#停止时等待多少秒
stopawitsecs = 0
#自动开始
autostart = true
autorestart = true
#输出的log文件
stdout_logfile = /home/go_rent/logs/supervisor.log
#输出的错误文件
redirect_stderr = true
# stderr_logfile = /home/go_rent/logs/logs/supervisor.err
redirect_stderr = true

[supervisord]
loglevel = info

#使用 supervisorctl 配置
[inet_http_server]
port = :9001

[supervisorctl]
#使用supervisorctl的登录地址和端口号
serverurl = http://127.0.0.1:9001

#不定义命名空间
[rpcinterface:supervisor]
supervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface
```
