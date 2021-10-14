package api

import (
	"log"
	"simple-bank/db/sqlc"

	"github.com/gin-gonic/gin"
)


type Server struct  {
	store *db.Store
	router *gin.Engine
}


func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	//handling all the routes

	router.GET("/", func(c *gin.Context){
		log.Println("i am here")
		c.JSON(200, gin.H{
			"msg" :"Welcome",
		})
	})

	router.POST("/create", server.CreateAccount);

	router.GET("/get/:id", server.GetAccount)

	router.GET("/get", server.ListAccounts)
	
	server.router = router;
	return server;
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
