###
# Targets for building and running the API & Webapp
#
.PHONY: build
build: clean
	docker-compose build

.PHONY: clean
clean:
	mkdir -p ./bin || true
	[[ -f ./bin/scans.db ]] || touch ./bin/scans.db


.PHONY: run
run:
	docker-compose up

###
# Targets for the Go API
#
.PHONY: api.build
api.build:
	go build -o ./bin/scanner ./cmd/scanner

.PHONY: api.run
api.run: clean
	PORT=8081 ./bin/scanner

.PHONY: api.build.docker
api.build.docker:
	docker-compose build api

.PHONY: api.test
api.test:
	go test -v $(shell go list ./...)

.PHONY: api.run.docker
api.run.docker:
	docker-compose up api

###
# Targets for the Angular Webapp
#
.PHONY: web.build
web.build:
	cd webapp && ng build

.PHONY: web.run
web.run:
	cd webapp && ng serve

.PHONY: web.build.docker
web.build.docker:
	docker-compose build web

.PHONY: web.run.docker
web.run.docker:
	docker-compose up web
