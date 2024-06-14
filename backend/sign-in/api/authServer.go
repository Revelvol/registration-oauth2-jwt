package api

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"revelvoler/registration-service/internal/model"
	"revelvoler/registration-service/internal/service"
	"time"

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
		v1.POST("/login", server.loginEndpoint)
		v1.POST("/email/auth", server.generateEmailCodeEndpoint)
		v1.POST("/email/validate", server.validateEmailEndpoint)
		v1.POST("/refresh", server.refreshEndpoint)
	}

	//return server
	server.Router = router
	return server
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type loginRequest struct {
	Code    string `json:"code"`
	Channel string `json:"channel" binding:"required"`
}
type googleDetailResponse struct {
    ID            string `json:"id"`
    Email         string `json:"email"`
    VerifiedEmail bool   `json:"verified_email"`
    Name          string `json:"name"`
    GivenName     string `json:"given_name"`
    FamilyName    string `json:"family_name"`
    Picture       string `json:"picture"`
    Locale        string `json:"locale"`
}

func (server *Server)loginEndpoint(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var user model.User
	var userToken model.UserToken
	// based on channel do login flow (ex google)
	if req.Channel == "google" {
		// use code required from frontend to exchange access_token with google
		token, err := googleOauthConfig.Exchange(c, req.Code)
		if err != nil {
			// todo return context gin error here
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		userToken = model.UserToken{
			AccessToken:  sql.NullString{String:token.AccessToken, Valid: true},
			RefreshToken: sql.NullString{String:token.RefreshToken, Valid: true},
			ExpitedAt:        token.Expiry,
			Channel: sql.NullString{String:req.Channel, Valid: true},
		}

		// GET user detail from google 
		response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		// decode json into struct
		var googleUser googleDetailResponse
		if err = json.Unmarshal(contents, &googleUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		// put detail response into user if posible 
		user.Email = googleUser.Email
	} else {
		// do normal validation 
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsuported login type, still in development"})
	}
	// get or save user TBL_USER 
	service.SaveOrGetUserData(server.DB, &user) 
	// save token and channel to TBL_TOKEN
	service.SaveOrUpdateToken(server.DB, &userToken)
	expireAt := userToken.ExpitedAt.Minute() - time.Now().Minute()
	// generate JWT Token and response to frontend
	jwtToken, _ :=  service.GenerateTokenFromUserExpireInEpoch(user.Email, req.Channel, int(userToken.ExpitedAt.Unix()))

	// return token 
	c.JSON(http.StatusOK, gin.H{
		"token_type": "Bearer", 
		"access_token" : jwtToken,
		"expire_in" : expireAt,
	})
}

func (server *Server) refreshEndpoint(c *gin.Context) {
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

func (server *Server) generateEmailCodeEndpoint(c *gin.Context) {
	// try send email

	// if failed return email not valid, or something went wrong please try again later

	// if sucess save to db TBL_USER_EMAIL_VALIDATION

	// return expired date and email sended verification code

}

func (server *Server) validateEmailEndpoint(c *gin.Context) {

	// get email validation detail, check exist or expired

	// validate code

	// return success, update TBL_USER validated

}
