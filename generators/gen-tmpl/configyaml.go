package gen_tmpl

var ConfigYamlContent = `# 使用yaml做配置项
# 数据库配置项
db:
  mysql:
    dbname: "example"
    password: "root1234"
    username: "root"
    port: 3306
    host: "127.0.0.1"
log:
  # 默认路径是运行程序的目录
  #  logDirector: ./logging
  logAutoFile: log.middleware
  logInfoFile: log.manual
`