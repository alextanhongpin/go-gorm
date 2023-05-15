include .env
export

q := $(shell cat test.sql)


start:
	@go run main.go


up:
	@docker-compose up -d


down:
	@docker-compose down


sql: # make sql q='select 1 + 1'
	@docker-compose exec mysql mysql -u $(DB_USER) -p$(DB_PASS) -D$(DB_NAME) -e '$(q)'


sql-cli:
	@docker-compose exec mysql mysql -u $(DB_USER) -p$(DB_PASS) -D$(DB_NAME)
