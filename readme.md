# ginctl

- 自动生成以`gin`为核心的`web`服务

- 完全兼容`goctl` api 语法
  
- 完全兼容 `goctl` 命令格式


- 严格以`api`文件生成代码即`api`没有描述的`handler`,`middleware`会被删除。这是`goctl` 不具备的
 # 生成gorm风格

 ginctl model mysql datasource --url "user:password@tcp(ip:port)/databasename" --dir ./model -t "cronjob,acl_groups" -m gorm