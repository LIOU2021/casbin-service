# casbin service

## todo
1. golang load conf.yml
2. config.yml 寫在`data\src`，dockerfile要copy這個檔案到第二個stage。复制example变成config.yml
3. 将logger配置读取config.yml
## start
```
# 1. 處理docker-compose env
cp .env.example .env

# 2. 處理 golang server config
cp data/src/config.yml.example data/src/config.yml

# 3. 起docker-compose容器
make up
```
## tip
- 日誌。支援輪播功能
    - data\log\api\access.log
    - data\log\server\server.log
## Ref
- [Go 每日一库之 casbin](https://darjun.github.io/2020/06/12/godailylib/casbin/)
- [casbin official github](https://github.com/casbin/casbin#installation)
- [casbin.org](https://casbin.org/)
- [casbin online editor](https://casbin.org/editor/)
- [Use your own storage adapter](https://casbin.org/docs/adapters/#use-your-own-storage-adapter)
- [Database Storage Format](https://casbin.org/docs/policy-storage/#database-storage-format)