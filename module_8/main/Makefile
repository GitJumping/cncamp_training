#export tag=v1.0
#
#build:
#	echo "building main binary"
#	mkdir -p bin/amd64
#	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/amd64 main.go
#
#release: build
#	echo "building main container"
#	docker build -t module_2/main:${tag} .
#
#push: release
#	echo "pushing module_2/main"
#	docker push module_2/main:${tag}

export tag=v4.0
build:
	echo "building main binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/amd64 main.go

docker-build:
	docker build -t cncamp/golang-http:${tag} .
	docker push cncamp/golang-http:${tag}