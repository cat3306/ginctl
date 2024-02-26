name: {{.serviceName}}
host: {{.host}}
port: {{.port}}
release: {{.release}}

{{if .mysql}}
mysqlConfig:
  ip: 127.0.0.1
  port: 3306
  pwd: "12345678"
  user: root
  connectPoolSize: 100
  setLog: true
{{end}}

{{if .redis}}
redisConfig:
  ip: 127.0.0.1
  port: 6379
  pwd: "redis-hahah@123"
  db:
    - 0
{{end}}