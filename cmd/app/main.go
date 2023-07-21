package main

import (
	"fmt"

	"github.com/bitrix1/go-examples/internal/app/c_lib_connect"
)

func main() {
	c_lib_connect.Main()
	fmt.Printf("Done1.")
}

/*
import (
	"fmt"
	"github.com/bitrix1/go-examples/internal/app/hash_map_ex"
	"github.com/bitrix1/go-examples/internal/app/leetcode_ex"

)
func main() {
	hash_map_ex.Test()
	leetcode_ex.Test()
	fmt.Printf("Done.")
}
*/

/*
import (
	"fmt"
	// C:\Program Files\Go\src\runtime\map.go
	"github.com/bitrix1/go-examples/internal/app/cmd_win"
	"github.com/bitrix1/go-examples/internal/app/http_jwt"
)

func main() {
	cmd_win.Main()
	http_jwt.Run()
	// jwt_ex.Test()
	fmt.Printf("Done.")
}
*/
