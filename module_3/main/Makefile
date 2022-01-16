export tag=v1.0

build:
	echo "building main binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/amd64 main.go

release: build
	echo "building main container"
	docker build -t cncamp/main:${tag} .

push: release
	echo "pushing module_2/main"
	docker push cncamp/main:${tag}