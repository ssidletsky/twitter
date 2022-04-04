build:
	docker-compose build

up:
	. .env
	docker-compose up -d

down:
	docker-compose down
