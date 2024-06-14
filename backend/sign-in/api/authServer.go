package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"revelvoler/registration-service/internal/model"

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

// app holds the Cloud IAP certificates and audience field for this app, which
// are needed to verify authentication headers set by Cloud IAP.
type app struct {
	certs map[string]string
	aud   string
}

type loginRequest struct {
	code    string `json:"code"`
	channel string `json:"channel" binding:"required"`
}
type googleDetailResponse struct {
	Name  string
	Order string
}

// OAuth configuration
var googleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
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
		v1.POST("/email/validate", validateEmailEndpoint)
		v1.POST("/refresh", refreshEndpoint)
	}

	//return server
	server.Router = router
	return server
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func loginEndpoint(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var user model.User
	// based on channel do login flow (ex google)
	if req.channel == "google" {
		// use code required from frontend to exchange access_token with google
		token, err := googleOauthConfig.Exchange(c, req.code)
		if err != nil {
			// todo return context gin error here
			panic("Failed to get user info" + err.Error())
		}
		accessTokenResponse := map[string]interface{}{
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
			"expiry":        token.Expiry,
		}

		// GET user detail from google 
		response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
		if err != nil {
			panic(fmt.Errorf("failed getting user info: %s", err.Error()))
		}
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(fmt.Errorf("failed read response: %s", err.Error()))
		}
		// decode json into struct
		var googleUserDetail googleDetailResponse
		if err = json.Unmarshal(contents, &googleUserDetail); err != nil {
			fmt.Println("error:", err)
		}

	} else {
		// do normal validation 
	}
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

func validateEmailEndpoint(c *gin.Context) {

	// get email validation detail, check exist or expired

	// validate code

	// return success, update TBL_USER validated

}
