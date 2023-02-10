# go_rent

租房小程序服务端

使用框架: beego

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
