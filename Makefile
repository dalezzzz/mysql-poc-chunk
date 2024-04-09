db-run:
	docker compose up -d 
db-down:
	docker compose down --volumes
problem-run:
	go run problem/main.go
solution-run:
	go run solution/main.go