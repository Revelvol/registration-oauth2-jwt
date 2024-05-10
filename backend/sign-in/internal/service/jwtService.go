package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type AuthClaims struct {
	UserName string `json:"foo"`
	jwt.RegisteredClaims
}

// Global variable declaration without semicolons
var hmacSampleSecret = []byte("someValueFromDotENV")

func GenerateTokenFromUserExpireInEpoch(expiredAtEpoch int) string {
	
	claims := AuthClaims {
		"julius",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add( time.Duration(expiredAtEpoch) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			// NotBefore: jwt.NewNumericDate(time.Now()), 
			Issuer:    "test",
			Subject:   "somebody",
			// ID:        "1"
			// Audience:  []string{"somebody_else"} 
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return "rtuern error early hare "
	}
	return tokenString; 
}


// IsTokenValid checks whether the given token string is valid
func IsTokenValid(tokenString string) bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if err != nil {
		// Handle token parsing errors
		if errors.Is(err, jwt.ErrTokenMalformed) {
			log.Fatal("That's not even a valid token")
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			log.Fatal("Invalid token signature")
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			log.Fatal("Token is expired or not active yet")
		} else {
			log.Fatal("Unexpected token error:", err)
		}
		return false
	}

	if token.Valid {
		fmt.Println("JWT token is valid")
		return true
	}

	// If the token is not valid for some other reason, handle accordingly
	fmt.Println("Token is invalid for an unspecified reason")
	return false
}
func ExtractValidTokenClaims(tokenString string)  {

	authClaims := AuthClaims{}
	// extract claims 
	token, err := jwt.ParseWithClaims(tokenString, &authClaims, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		log.Fatal(err)
	} 
	// No need for explicit type assertion; token.Claims already points to authClaims
	// else if claims, ok := token.Claims.(*AuthClaims); ok {
	// 	// extract the claims as normal 
	// 	fmt.Println(claims.UserName, claims.RegisteredClaims.Issuer)

	// } else {
	// 	log.Fatal("unknown claims type, cannot proceed")
	// }

	if token.Valid {
		fmt.Println("Token is valid")
		// Now you can access the fields in authClaims
		fmt.Println("UserName:", authClaims.UserName)
		fmt.Println("Issuer:", authClaims.RegisteredClaims.Issuer)
	} else {
		fmt.Println("Token is invalid") 
	}
}


