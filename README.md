# go_rent

租房小程序服务端

使用框架: beego

# 小程序端：
https://github.com/RelaxedDong/rent_mini

# 配置文件
```
appname = xxx
httpport = 8000
runmode = dev
copyrequestbody = true
logfilepath=logs/service.log
db_user=root
db_password=root
db_host=localhost
db_port=3306
db_database=rent
[dev]
httpaddr = 127.0.0.1
APPID=xxx # 小程序appid
APP_SECRET=xxx # 小程序 secret
[pro]
httpaddr = 127.0.0.1
APPID=xxx
APP_SECRET=xxx
```
