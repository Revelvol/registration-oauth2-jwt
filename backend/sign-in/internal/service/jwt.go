package jwtService

import {
	"github.com/dgrijalva/jwt-go"
}


type AuthClaims {
	UserName string `json:"foo"`
	jwt.RegisteredClaims
}

func GenerateTokenFromUserExpireInEpoch(int expiredAt) string {

	hmacSampleSecret := "someValueFromDotENV"

	claims := AuthClaims {
		"julius",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(expiredAt),
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
	return tokenString; 
}


func IsTokenValid(tokenString string) bool {
	// parse token 
	hmacSampleSecret := "someValueFromDotENV"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate something here like is the code correct 
		if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != null {
		return false 
	}

	// validate token behaviour 
	switch {
		case token.Valid:
			fmt.Println("json token is valid")
			return true
		case errors.Is(err, jwt.ErrTokenMalformed):
			fmt.Println("That's not even a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			fmt.Println("Invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		default:
			fmt.Println("Couldn't handle this token:", err)
			return false 
	}
}

func ExtractValidTokenClaims(tokenString string, authClaims *AuthClaims)  {
	// extract claims 
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(authClaims); ok {
		fmt.Println(claims.UserName, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
}


