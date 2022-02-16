package app

var ConfigTemplate = `
#生成代码的配置
controllerPath: app/http/controllers/backend
modelPath: app/models
eventPath: app/events
listenerPath: app/listeners

#---------------------------
debug: true

appName: "后台管理系统"

#模板,不使用map保证从第一个元素执行
template:
  -
    modulename: "backend"
    layout: "views/layouts/backend.html"
    viewpath: "views/backend/admin/home.html"
  -
    modulename: "block"
    layout: "views/layouts/block.html"
    viewpath: "views/fronend/auth/*.html;views/backend/auth/*.html;views/backend/common/*.html;views/backend/**/*.html"
  -
    modulename: "fronend"
    layout: "views/layouts/fronend.html"
    viewpath: "views/fronend/**/*.html"

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

captcha:
  width: 96
  height: 38
  num: 5
  secret-key: mycaptcha
  expired: 60
mysql:
  host: 127.0.0.1
  port: 3306
  username: root
  password: admin123
  dbname: ginrbac
  charset: utf8mb4
logger:
  link-name: lastest-log
  encoding: json
  output-dir: logs
  #文件最大保存时间，小时为单位
  max-age: 720
  #文件切割间隔，小时为单位
  rotation-time: 24
jwt:
  sign-key: myjwt
  expires-time: 3600
redis:
  db: 0
  addr: 127.0.0.1:6379
  password: ""
`
