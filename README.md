## Simple golang programing language forum in VietNam 

 ![alt golang](https://raw.githubusercontent.com/golang-vietnam/forum/master/public/img/isomorphic_code_share.png "golang programing language")

### Core dependencies
- Web framework: [Gin](https://github.com/gin-gonic/gin)
- NoSql-MongoDb: [mgo](http://labix.org/mgo)
- Template engine: [Pongo2](https://github.com/flosch/pongo2)
- Testing: [GoConvey](http://goconvey.co/)
- Front-end framework: [Materializecss](http://materializecss.com)

### Install dependencies
    
    make install

### Run project
On `development` mode we use [Gin](https://github.com/codegangsta/gin) to live reload. Server run on port 3000

    make

### Test server
- Run server on test mode: `make runOnTest`
- In other terminal: `make test` or `make testOnWeb` then see test on `localhost:8080`

