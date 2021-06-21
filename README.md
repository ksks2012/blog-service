# BLOG service

## structure

- configs: setting file
- docs: document
- global: global variables
- internal: internal module
	<!-- TODO: -->
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
## Binaries

```sh
go build ./cmd/blog-service/
```
