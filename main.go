package main

import (
	"DRD/db"
	"DRD/domains/ESign"
	drdsession "DRD/domains/Session"
	"DRD/routes"
)

func main() {

	db.GetDb()
	// Adding routing each domain has
	ESign.InitiateRoutes()
	drdsession.InitiateRoutes()

	// Start Server
	r := routes.GetInstance()
	r.Run(":8080")
}
