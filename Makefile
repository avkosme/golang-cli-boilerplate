build:
	docker-compose run --rm -v `pwd`:`pwd` -w `pwd` app go build  -o ./bin ./...
start:
	docker-compose up -d
stop:
	docker-compose stop
restart:
	@make stop
	@make start
test:
	docker-compose run --rm -v `pwd`:`pwd` -w `pwd` app go build  -o ./bin ./...
	docker-compose up -d mongo-test
	docker-compose run --rm -v `pwd`:`pwd` -w `pwd` app go test -count=1 -v ./...
	docker-compose stop
