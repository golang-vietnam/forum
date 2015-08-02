## Simple golang programing language forum in VietNam 

**We temporary paused this project in order to build a new link sharing site for developers. Please check it at [grox](https://github.com/golang-vietnam/grox) project.**

 ![alt golang](https://raw.githubusercontent.com/golang-vietnam/forum/master/isomorphic_code_share.png "golang programing language")

### Core dependencies
- Web framework: [Gin](https://github.com/gin-gonic/gin)
- NoSql-MongoDb: [mgo](http://labix.org/mgo)
- Testing: [GoConvey](http://goconvey.co/)

### Install dependencies
    
    go get github.com/golang-vietnam/forum
    cd $GOPATH/src/github.com/golang-vietnam/forum;make install
    
### Run project
On `development` mode we use [Gin](https://github.com/codegangsta/gin) to live reload. Server run on port 3000

    make

### Test server
	
	make test