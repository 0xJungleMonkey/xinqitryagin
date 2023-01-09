[comment]: <> (This is a generated file please edit source in ./templates)
[comment]: <> (All modification will be lost, you have been warned)
[comment]: <> ()
### Sample CRUD API for the mysql database doadmin:AVNS_nQrjtn8ilVHqYs6xIim@tcp(dbaas-db-7154856-do-user-13260059-0.b.db.ondigitalocean.com:25060)/rocket_development?

## Example
The project is a RESTful api for accessing the mysql database doadmin:AVNS_nQrjtn8ilVHqYs6xIim@tcp(dbaas-db-7154856-do-user-13260059-0.b.db.ondigitalocean.com:25060)/rocket_development?.

## Project Files
The generated project will contain the following code under the `./example` directory.
* Makefile
  * useful Makefile for installing tools building project etc. Issue `make` to display help
* .gitignore
  * git ignore for go project
* go.mod
  * go module setup, pass `--module` flag for setting the project module default `example.com/example`
* README.md
  * Project readme
* app/server/main.go
  * Sample Gin Server, with swagger init and comments
* api/*.go
  * REST crud controllers
* dao/*.go
  * DAO functions providing CRUD access to database
* model/*.go
  * Structs representing a row for each database table

The REST api server utilizes the Gin framework, GORM db api and Swag for providing swagger documentation
* [Gin](https://github.com/gin-gonic/gin)
* [Swaggo](https://github.com/swaggo/swag)
* [Gorm](https://github.com/jinzhu/gorm)

## Building
```.bash
make example
```
Will create a binary `./bin/example`

## Running
```.bash
./bin/example
```
This will launch the web server on localhost:8080

## Swagger
The swagger web ui contains the documentation for the http server, it also provides an interactive interface to exercise the api and view results.
http://localhost:8080/swagger/index.html

## REST urls for fetching data


* http://localhost:8080/activeadmincomments
* http://localhost:8080/activestorageattachments
* http://localhost:8080/activestorageblobs
* http://localhost:8080/addresses
* http://localhost:8080/adminusers
* http://localhost:8080/arinternalmetadata
* http://localhost:8080/batteries
* http://localhost:8080/blazeraudits
* http://localhost:8080/blazerchecks
* http://localhost:8080/blazerdashboardqueries
* http://localhost:8080/blazerdashboards
* http://localhost:8080/blazerqueries
* http://localhost:8080/buildingdetails
* http://localhost:8080/buildings
* http://localhost:8080/columns
* http://localhost:8080/customers
* http://localhost:8080/elevators
* http://localhost:8080/employees
* http://localhost:8080/interventions
* http://localhost:8080/leads
* http://localhost:8080/maps
* http://localhost:8080/quotes
* http://localhost:8080/schemamigrations
* http://localhost:8080/users_

## Project Generated Details
```.bash
gen \
    --sqltype=mysql \
    --connstr \
    doadmin:AVNS_nQrjtn8ilVHqYs6xIim@tcp(dbaas-db-7154856-do-user-13260059-0.b.db.ondigitalocean.com:25060)/rocket_development? \
    --database \
    rocket_development \
    --json \
    --gorm \
    --guregu \
    --rest \
    --out \
    ./rocket \
    --module \
    rocket \
    --mod \
    --server \
    --makefile \
    --json-fmt=snake \
    --generate-dao \
    --generate-proj \
    --overwrite
```











