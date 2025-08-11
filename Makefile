 
build:
	go build -o bitcoinmonitor cmd/bitcoinmonitor/main.go

# run: build
# 	./bitcoinmonitor

run:
	go run cmd/bitcoinmonitor/main.go

docker-build:
	docker build -t bitcoinmonitor .

docker-run: docker-build
	docker run -p 8080:8080 bitcoinmonitor

swagger:
	swag init -g cmd/bitcoinmonitor/main.go --output ./docs

compose-up:
	docker-compose up -d

compose-build:
	docker-compose up -d --build

