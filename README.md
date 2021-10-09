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
go generate github.com/blog-service/internal/dao/dbversion/mysql
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
	> curl -X POST http://127.0.0.1:18080/upload/file -F file=@"./test.png" -F type=1

- test create tags

	> curl -X POST http://127.0.0.1:18080/api/v1/tags -F 'nane=RUST' -F 'create_by=hong'

- test get tags

	> curl -X GET http://127.0.0.1:18080/api/v1/tags?page=1&page_size=2

- test post auth

	> curl -X POST http://127.0.0.1:18080/auth -H 'app_key:eddycjy' -H 'app_secret: hong'

## auth

### test data

> INSERT INTO `BlogService`.`blog_auth` (`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `is_del`) VALUES ('1', 'hong', 'blog-service', '0', 'hong', '0', '', '0', '0');
