package jwt_ex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func validateToken(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("x-Myheader", "token")
    res.Write([]byte("validate"))
}

func handleToken(res http.ResponseWriter, req *http.Request) {
    type MyCustomClaims struct {
        Admin bool   `json:"admin"`
        Name  string `json:"name"`
        jwt.StandardClaims
    }

    claims := MyCustomClaims{
        true,
        "John",
        jwt.StandardClaims{
            ExpiresAt: 15000,
            Issuer:    "website",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, _ := token.SignedString([]byte("secret"))

    res.Header().Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
	//
	cookie := http.Cookie{}
	cookie.Name = "name"
	cookie.Value = "value"
	cookie.Expires = time.Now().Add(356 * 24 * time.Hour)
	cookie.Secure = true 
	//test HttpOnly in web console -> document.cookie = ''
	cookie.HttpOnly = true 
	http.SetCookie(res, &cookie)
	//
    tok, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(tok *jwt.Token) (interface{}, error) {
        if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", tok.Header["alg"])
        }
        return []byte("secret"), nil
    })

    if claims, ok := tok.Claims.(*MyCustomClaims); ok && tok.Valid {
        fmt.Println(claims.Admin, claims.Name)
    } else {
        fmt.Println(err)
    }

    res.Write([]byte(tokenString))
}

func handleHome(res http.ResponseWriter, req *http.Request) {
    type Product struct {
        Name string
    }
    prod := Product{
        Name: "john",
    }
    payload, _ := json.Marshal(prod)
    res.Write([]byte(payload))
}

func RunServer() {
    http.HandleFunc("/validate", validateToken)
    http.HandleFunc("/token", handleToken)
    http.HandleFunc("/", handleHome)
	full_adress := "127.0.0.1:8080"
	fmt.Printf("http://%v/",full_adress)
	// fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
    http.ListenAndServe(full_adress, nil)
}
