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

	> curl -X GET 'http://127.0.0.1:18080/api/v1/tags?page=1&page_size=2'

- test post auth

	> curl -X POST http://127.0.0.1:18080/auth -F 'app_key=hong' -F 'app_secret=blog-service'

- add token

	> -H token:<token>

## auth

### test data

> INSERT INTO `BlogService`.`blog_auth` (`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `is_del`) VALUES ('1', 'hong', 'blog-service', '0', 'hong', '0', '', '0', '0');


# Tracing
## jaeger

### docker

'''
docker run -d --name jaeger \
 -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.27
'''

## pkg

go get -u github.com/opentracing/opentracing-go@v1.1.0
go get -u github.com/uber/jaeger-client-go@v2.22.1

go get -u github.com/eddycjy/opentracing-gorm

# hotfix

## pkg

go get -u github.com/fsnotify/fsnotify
