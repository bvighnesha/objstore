start:
	docker run --name redis-server -p 6379:6379 -d redis
	go run main.go
stop:
	docker stop redis-server
	docker rm redis-server