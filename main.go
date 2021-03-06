package main

import (
	"log"

	"github.com/takama/router"
	"github.com/tectiv3/standardfile/routing"
)

func logger(handle router.Handle) router.Handle {
	return func(c *router.Control) {
		log.Printf("%s\t%s", c.Request.Method, c.Request.RequestURI)
		handle(c)
	}
}

func panicHandler(c *router.Control, err interface{}) {
	log.Println(err)
	c.Code(500).Body("")
}

func main() {
	r := router.New()
	r.CustomHandler = logger
	r.PanicHandler = panicHandler

	r.GET("/", routing.Dashboard)

	r.POST("/api/items/sync", routing.SyncItems)
	r.POST("/api/items/backup", routing.BackupItems)
	// r.DELETE("/api/items", routing.DeleteItems)

	r.POST("/api/auth", routing.Registration)
	r.PATCH("/api/auth", routing.ChangePassword)
	r.POST("/api/auth/change_pw", routing.ChangePassword)
	r.POST("/api/auth/sign_in", routing.Login)
	r.POST("/api/auth/sign_in.json", routing.Login)
	r.GET("/api/auth/params", routing.GetParams)

	log.Println("Running on port 8888")
	r.Listen(":8888")
}
