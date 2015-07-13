export GOPATH:=$(shell pwd)

default: dev

backend: deps-backend
	go install ddesktop/main
	mv bin/main bin/ddesktop

frontend: deps-frontend
	rm -rf bin/webroot
	cp -r src/ddesktop/webroot bin

deps: deps-frontend deps-backend

deps-backend:
	go get -d -v ddesktop/...

deps-frontend:
	bower install

dev: backend frontend
	cp config.yml bin
	cd bin && sudo ./ddesktop