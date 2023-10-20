.PHONY: clear clear-log-server clear-mysql-data

up:
	@docker-compose up -d	

down:
	@docker-compose down

build:
	@docker-compose build casbin

clear: clear-log clear-mysql-data

clear-log:
	rm -rf ./data/log/server
	rm -rf ./data/log/api
	rm -rf ./data/log/mysql

clear-mysql-data:
	rm -rf ./data/mysql/data