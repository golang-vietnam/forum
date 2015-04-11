## Golang.vn golang programing language forum in VietNam 

 ![alt golang](https://raw.githubusercontent.com/golang-vietnam/forum/master/public/img/logo.jpg "golang programing language")

### Project structure

    |- YourGOPATH
        |-- bin
        |-- pkg
        |-- src
            |--github.com/golang-vietnam/forum

### Core dependencies
- Web framework: [Gin](https://github.com/gin-gonic/gin)
- ORM: [GORM](https://github.com/jinzhu/gorm)
- Template engine: [Pongo2](https://github.com/flosch/pongo2)
- Testing: [GoConvey](http://goconvey.co/)
- Front-end framework: [Materializecss](http://materializecss.com)

### Install dependencies
#### If you don't need use 'godep' to manage dependencies version (always lastest) that mean will break your project
    go get ./..
 
#### Use 'godep' recommend
    
    go get github.com/kr/godep
    export PATH=$PATH:$GOPATH/bin

Cd to forum package
    
    cd $GOPATH/src/github.com/golang-vietnam/forum
    godep restore

Install new dependencies then save

    godep save

More about [godep](https://github.com/tools/godep) 
