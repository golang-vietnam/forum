## Simple golang programing language forum in VietNam 

 ![alt golang](https://raw.githubusercontent.com/golang-vietnam/forum/master/public/img/isomorphic_code_share.png "golang programing language")

### Core dependencies
- Web framework: [Gin](https://github.com/gin-gonic/gin)
- NoSql-MongoDb: [mgo](http://labix.org/mgo)
- Testing: [GoConvey](http://goconvey.co/)

### Install dependencies
    
    make install

### Run project
On `development` mode we use [Gin](https://github.com/codegangsta/gin) to live reload. Server run on port 3000

    make

### Test server
	
	make test