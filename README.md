# AUTO-SMS-GO

> 一个go的gin+grpc的练手项目,目的是熟悉使用golang编程
> 
> 将用到redis,mysql,rabbitmq等工具

## TODO LIST

- [x] add grpc and etcd dependencies
  - [x] use etcd&grpc implement register and discovery
- [x] add mysql support
  - [x] use sqlx implement basic sql support
- [x] add redis support
  - [x] use redis implement login token
  - [x] use redis implement crontab
- [x] add rabbitmq support
  - [x] use rabbitmq implement sms send
- [ ] add http gateway
