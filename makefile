.PHONY: clear clear-log-server clear-mysql clear-mysql-conf clear-mysql-data

up:
	@docker-compose up -d	

down:
	@docker-compose down

build:
	@docker-compose build casbin

clear: clear-log-server clear-mysql

clear-log-server:
	rm -rf ./data/log/server

clear-mysql: clear-mysql-conf clear-mysql-data

clear-mysql-conf:
	rm -f ./data/mysql/my.cnf

clear-mysql-data:
	rm -rf ./data/mysql/data