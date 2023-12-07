# basic_api

using Go and gin framework

### go mod setup (if needed)
initialize 1st time

    go mod init
add package

    go get github.com/githubnemo/CompileDaemon
  
 install package

     go install github.com/githubnemo/CompileDaemon
 vendorize

     go mod vendor
     go mod tidy


### Compile code

    CompileDaemon -command="./basic_api"