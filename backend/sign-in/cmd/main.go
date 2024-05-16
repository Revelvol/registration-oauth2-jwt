package main

import (
	"log"
	"os"
	"revelvoler/registration-service/internal/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// app holds the Cloud IAP certificates and audience field for this app, which
// are needed to verify authentication headers set by Cloud IAP.
type app struct {
	certs map[string]string
	aud   string
}

func main(){
	// gorm logger setting
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,             // Don't include params in the SQL log
			Colorful:                  true,
		},
	)
	// connect to database
	dsn := "host=127.0.0.1 user=admin password=password dbname=local-playground port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
            TablePrefix:   "core_login", // schema name
            SingularTable: false,
        },
	})
	if err != nil {
		panic("failed to connect to database")
	}
	db.Migrator().DropTable(&model.User{}, &model.UserDetail{}, &model.UserToken{}, &model.EmailValidation{})
	db.Migrator().AutoMigrate(&model.User{}, &model.UserDetail{}, &model.UserToken{}, &model.EmailValidation{})

	//setup gin
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3001") // listen and serve on 0.0.0.0:3001

}