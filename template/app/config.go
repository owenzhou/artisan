package app

var ConfigTemplate = `
#生成代码的配置
controllerPath: app/http/controllers
modelPath: app/models
eventPath: app/events
listenerPath: app/listeners

#---------------------------
debug: true

timezone: "Asia/Shanghai"

appName: "My App"

template:
  -
    modulename: "fronend"
    layout: "views/layouts/layout.html"
    viewpath: "views/home/*.html"

#用户登录
auth:
  defaults:
    guard: web
  guards:
    web:
      driver: session
      provider: users
    api:
      driver: jwt
      provider: users
    admin:
      driver: session
      provider: adminusers
  providers:
    users:
      driver: database
      table: users
    adminusers:
      driver: database
      table: users
session:
  secret-key: "VMYHO4Z233HSIQTZFIJHDOU6O7XVPICDF6X2EY474MIB7UDE5NYA"
  lifetime: 7200
captcha:
  width: 96
  height: 38
  num: 5
  secret-key: mycaptcha
  lifetime: 60
mysql:
  host: 127.0.0.1
  port: 3306
  username: root
  password: 
  dbname: ginrbac
  charset: utf8mb4
  maxOpenConns: 25
  maxIdleConns: 25
  #以分为单位
  connMaxLifeTime: 5
logger:
  output-dir: logs
  filename: app.log
  encoding: json
  #文件最大保存时间，天为单位
  maxage: 30
  #文件最大大小，兆为单位
  maxsize: 10
  #文件保存数量，个为单位
  maxbackups: 30
  #是否压缩
  compress: false
jwt:
  sign-key: myjwt
  lifetime: 3600
redis:
  db: 0
  addr: 127.0.0.1:6379
  password: ""
`
