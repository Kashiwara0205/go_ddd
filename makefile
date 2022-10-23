in:
	docker exec -ti go_ddd_go_ddd_1 /bin/bash
up:
	docker-compose up
down:
	docker-compose down
reboot:
	docker-compose down
	docker-compose up
go_test:
	docker exec -ti go_ddd_go_ddd_1 bash -c "go fmt ./..."
	docker exec -ti go_ddd_go_ddd_1 bash -c "go vet ./..."
	docker exec -ti go_ddd_go_ddd_1 bash -c "go test -v ./..."