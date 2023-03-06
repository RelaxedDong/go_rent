# go_rent

# 租房小程序服务端

使用框架: `beego`

项目结构：

<img width="1100" alt="image" src="https://user-images.githubusercontent.com/38744096/218026368-264c75e5-930d-4af3-81d9-9b3ccd822034.png">


# 小程序端：
https://github.com/RelaxedDong/rent_mini

# Appconf 配置
```
# 放通用配置
httpaddr=127.0.0.1
appname=蚁租房
httpport=8000
runmode=dev
copyrequestbody=true
logfilepath=logs/service.log

oss_upload_path=https://bucket_name.oss-cn-shenzhen.aliyuncs.com
oss_region_host=https://bucket_name.oss-cn-shenzhen.aliyuncs.com

#邮件服务器
email_host="smtp.qq.com"
#服务端口
email_port=587
```

`conf/dev.env or conf/pro.env`

这里根据conf/app.conf runmode指定的模式动态获取环境变量文件，然后load到AppConfig里面，参考：<a href="init/init_env.go">初始化i配置</a>

环境变量配置文件：
```dotenv
oss_access_key=xxxx
oss_secret_key=xxx

db_port=xxx
db_database=xxx
db_host=xxx
db_user=xxxx
db_password=xxx

#发件人昵称
email_senderName="xx"
#发件人邮箱
email_user="xx@qq.com"
#发件人授权码
email_password="xxx"

# 小程序AppID
APPID=xxx
APP_SECRET=xxxx

```

# 部署


### k8s
...
### docker-compose、docker-swarm
...
### supervisor+nginx

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
