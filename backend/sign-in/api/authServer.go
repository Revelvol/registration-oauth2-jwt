package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

// server serves Http request
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// OAuth configuration
var oauthConfig = &oauth2.Config{
	ClientID:     "YOUR_GOOGLE_CLIENT_ID",
	ClientSecret: "YOUR_GOOGLE_CLIENT_SECRET",
	RedirectURL:  "postmessage",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}


func AuthServer(db *gorm.DB) *Server {
	server := &Server{DB: db}
	router := gin.Default()

	// group api 
	v1 := router.Group("/api/auth")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/email/auth", generateEmailCodeEndpoint)
		v1.POST("/email/validate",validateEmailEndpoint)
		v1.POST("/refresh",refreshEndpoint)
	}

	//return server 
	server.Router = router; 
	return server
}


func loginEndpoint(c *gin.Context){

	// todo transform this into gin format form http
	// Exchange the authorization code for tokens
	token, err := oauthConfig.Exchange(c, "some CODE strint")
	if err != nil {
		return
	}
	// Return the token details to the frontend
	response := map[string]interface{}{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"expiry":        token.Expiry,
	}
	fmt.Printf("%s", response)
	// call google to get user resource 

	// get or save user TBL_USER (might want to use redis)

	// save token and channel to TBL_TOKEN 

	// generate JWT Token and response to frontend
}


func refreshEndpoint(c *gin.Context) {
	// extract user data from jwt 

	// check curent session login channel 

	// if from 3rd party get refresh token and access_token DB based on channel to TBL_TOKEN 

	// try refresh token to 3rd party 

	// if succesfull update TBL TOKEN with new token 

	// (optional) from access_token get resource from 3rd party resource server 

	// get user detail from TBL_USER based on extracted user data or 3rd party resource server 

	// generate JWT TOKEN and response to frontend 

	//negative will generate forbiden and frontend will redirect to login 


}

func generateEmailCodeEndpoint(c *gin.Context) {
	// try send email 

	// if failed return email not valid, or something went wrong please try again later 

	// if sucess save to db TBL_USER_EMAIL_VALIDATION 
	
	// return expired date and email sended verification code 

}

func validateEmailEndpoint(c *gin.Context){

	// get email validation detail, check exist or expired 

	// validate code 

	// return success, update TBL_USER validated 

}

