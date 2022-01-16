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

export VERSION=3.0.1

build:
	docker build -t cncamp/httpserver:${VERSION} .
build.push:build
	docker push cncamp/httpserver:${VERSION}

run:
	docker run -p 8080:8080 cncamp/httpserver:${VERSION}

deploy:
	kubectl apply -f kubernate.yaml

del:
	kubectl delete -f kubernate.yaml