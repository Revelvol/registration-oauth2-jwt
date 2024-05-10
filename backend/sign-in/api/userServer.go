package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func UserServer(db *gorm.DB) *Server {
	server := &Server{DB: db}
	router := gin.Default()

	// group api 
	v1 := router.Group("/api/user")
	{
		v1.GET("/detail/{user-id}", getUserDetailEndpoint)
		v1.POST("/detail/{user-id}",postUserDetailEndpoint)
	}

	//return server 
	server.Router = router; 
	return server
}


func getUserDetailEndpoint(c *gin.Context) {

}

func postUserDetailEndpoint(c *gin.Context) {

}