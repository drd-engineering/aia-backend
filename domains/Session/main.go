package session

import (
	//"time"

	"DRD/routes"
	//"github.com/gin-contrib/cache"
	//"github.com/gin-contrib/cache/persistence"
)

func InitiateRoutes() {
	r := routes.GetInstance()

	//store := persistence.NewInMemoryStore(time.Second)

	apiroutes := r.Group("/api/v1")
	{
		session := apiroutes.Group("/session")
		{
			session.GET("/get-user/:token", getUser)
		}
	}
}
