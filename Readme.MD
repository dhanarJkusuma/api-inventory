
# api-inventory
REST API using `Golang` + `Sqlite3` + `Echo Framework` + `GORM`.
ref : [Trying Clean Architecure on Golang](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047) .

## How To Run
Execute command :
```
./start.sh
```
if it doesn't work, try to put this app directory to the $GOPATH/src directory

If you want to rebuild this app, you can run this command :
```
./bin/install.sh
```
For migration, use form-data with attribute-name `file` and type `file`


## Documentation
Documentation api using swagger.json/swagger.yaml in the root directory of this app.

## Library
Library that I've used for coding
- [validator](https://github.com/go-playground/validator)
- [GORM](https://github.com/jinzhu/gorm)
- [ECHO Framework](https://github.com/labstack/echo)
