//https://github.com/dgrijalva/jwt-go/issues/141
package http_jwt

import (
	"fmt"
	"time"

	// "github.com/golang-jwt/jwt"
	jwt "github.com/golang-jwt/jwt/v4"
)
func Test() {

}

var hmacSampleSecret []byte

func RunServer() {
	ExampleNewWithClaims_customClaimsType()
}
func ExampleParseWithClaims_customClaimsType() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjoic2luZ2xlIn0.QAWg1vGvnqRuCFTMcPkjZljXHh8U3L_qUjszOtQbeaA"

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		// Admin bool   `json:"admin"`
        // Name  string `json:"name"`
		jwt.StandardClaims // StandardClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.Issuer)
	} else {
		fmt.Println(err)
	}
}

func ExampleNewWithClaims_customClaimsType() {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)).Unix() ,
			IssuedAt:  jwt.NewNumericDate(time.Now()).Unix(),
			NotBefore: jwt.NewNumericDate(time.Now()).Unix(),
			Issuer:    "test",
			Subject:   "somebody",
			Id: "123",
			// ID:        `json:"jti,1"`,
			// ID: "sdf",
			// Audience:  []string{"somebody_else"},
		},
	}

	// Create claims while leaving out some of the optional fields
	claims = MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	//Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiZXhwIjoxNTE2MjM5MDIyfQ.xVuY2FZ_MRXMIEgVQ7J-TFtaucVFRXUzHm9LmV41goM <nil>
}

func RunServer1() {

	
	// For HMAC signing method, the key can be any []byte. It is recommended to generate
	// a key using crypto/rand or something equivalent. You need the same key for signing
	// and validating.

	

		// sample token string taken from the New example
		tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"
	
		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
	
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})
	
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
		} else {
			fmt.Println(err)
		}
	
	
	fmt.Printf("http_jwt")
}
