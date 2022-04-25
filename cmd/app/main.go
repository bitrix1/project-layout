package main

import (
	"fmt"

	"github.com/bitrix1/go-examples/internal/app/http_jwt"
	"github.com/bitrix1/go-examples/internal/app/jwt_ex"
)

//"github.com/bitrix1/go-examples/internal/app/http_jwt"
func main() {
	// run()
	http_jwt.RunServer()
	jwt_ex.Test()
	fmt.Printf("Done.")
}