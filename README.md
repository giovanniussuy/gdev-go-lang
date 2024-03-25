# gdev-go-lang
Project with live situations and learnings


go mod tidy

go build ./...
go run ./...

google.github.io/styleguide/go

# fiber + goccy



go get github.com/goccy/go-json

-import "encoding/json"
+import "github.com/goccy/go-json"


# deploy local



# go swagger 
para utilizar o swag do go precisa colocar a var SWAG em variaveis de ambiente

SWAG = C:\Users\giova\go\bin
Path = %SWAG%

instalar o swag
go get github.com/swaggo/swag/cmd/swag
swag init -g cmd/main.go
go get -u github.com/swaggo/fiber-swagger
import "github.com/swaggo/fiber-swagger" // fiber-swagger middleware 
// esse import não funfa, mas basta adicionar manualmente a rota do swagger no router

-- http://localhost:8080/swagger/index.html

