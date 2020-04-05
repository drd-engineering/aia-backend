package ESign

import (
	"fmt"
	//"time"

	"DRD/db"
	"DRD/routes"

	//"github.com/gin-contrib/cache"
	//"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func InitiateRoutes() {
	r := routes.GetInstance()

	//store := persistence.NewInMemoryStore(time.Second)

	apiroutes := r.Group("/api/v1")
	{
		eSignRoutes := apiroutes.Group("/e-sign")

		eSignRoutes.Use(ApplicationAuthorizationRequired())
		{
			eSignRoutes.POST("/save-user", saveUser)
			eSignRoutes.POST("/regenerate-link", regenerateLink)
		}

		eSignPublic := apiroutes.Group("/e-sign")
		{
			eSignPublic.POST("/save", saveSigns)
		}
	}
}

func ApplicationAuthorizationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header["Drd-Api-Key"]

		if len(apiKey) < 1 {
			fmt.Println(apiKey)
			c.AbortWithStatus(401)
			return
		}

		if apiKey[0] == "" {
			fmt.Println(apiKey[0])
			c.AbortWithStatus(401)
			return
		}

		dbInstance := db.GetDb()
		var token db.AppToken
		if err := dbInstance.Where("token = ? AND is_active", apiKey[0]).First(&token).Error; err != nil {
			c.AbortWithStatus(401)
			fmt.Println(err)
			return
		}
		fmt.Println(token)
		c.Next()
	}
}
