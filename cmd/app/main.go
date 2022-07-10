package main

import (
	"fmt"
	// C:\Program Files\Go\src\runtime\map.go
	"github.com/bitrix1/go-examples/internal/app/hash_map_ex"
	"github.com/bitrix1/go-examples/internal/app/leetcode_ex"
	// "github.com/bitrix1/go-examples/internal/app/http_jwt"
	// "github.com/bitrix1/go-examples/internal/app/cmd_win"
)

//"github.com/bitrix1/go-examples/internal/app/http_jwt"
func main() {
	hash_map_ex.Test()
	leetcode_ex.Test()
	// cmd_win.Main()
	// run()
	// http_jwt.Run()
	// jwt_ex.Test()
	fmt.Printf("Done.")
}
