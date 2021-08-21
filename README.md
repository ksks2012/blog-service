# BLOG service

## structure

- etc: setting file
- docs: document
- global: global variables
- internal (internal module):
	<!-- TODO: -->
	- dao: data access object
	- middleware
	- model: database model control
	- routers: api routes
	- service: process business logic
- pkg: package
- storage: temp file
- scripts: build, install, analysis scripts
- third_party: third_party tools
## design

### DB version

#### Add Table in Schema

edit sql command in md file

add table function in UpgradeSchema

#### Go generate

```sh
go generate github.com/blog-service/internal/dao/mysql
```

#### swaggo init

```sh
swag init -g cmd/blog-service/main.go
```

## Binaries

```sh
go build ./cmd/blog-service/
```

## Run

./blog-service


## api test command

- test file upload
	> curl -X POST http://127.0.0.1:18080/upload/file -F file=test.png -F type=1
