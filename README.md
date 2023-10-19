# casbin service

## todo
1. 先去go目錄增加yml讀取配置，分db、redis、log等大項。log參數要能提出日誌備份幾分跟幾天
2. config.yml 寫在`data\src`，dockerfile要copy這個檔案到第二個stage
## start
```
# 1. 處理docker-compose env
cp .env.example .env

# 2. 起docker-compose容器
make up
```
## Ref
- [Go 每日一库之 casbin](https://darjun.github.io/2020/06/12/godailylib/casbin/)
- [casbin official github](https://github.com/casbin/casbin#installation)
- [casbin.org](https://casbin.org/)
- [casbin online editor](https://casbin.org/editor/)
- [Use your own storage adapter](https://casbin.org/docs/adapters/#use-your-own-storage-adapter)
- [Database Storage Format](https://casbin.org/docs/policy-storage/#database-storage-format)