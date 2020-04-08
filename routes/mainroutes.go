package routes

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var instance *gin.Engine
var once sync.Once

func GetInstance() *gin.Engine {
	// Initiate value if there is no instance
	once.Do(func() {
		instance = gin.New()

		// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
		// By default gin.DefaultWriter = os.Stdout
		instance.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			// custom format Logging
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))

		instance.Use(gin.Recovery())
		instance.Use(CORSMiddleware())
	})
	return instance
}

func CORSMiddleware() gin.HandlerFunc {
	// Still need some improvements.
	// adding header to define the application detail
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
