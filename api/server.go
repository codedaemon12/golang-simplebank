package api

import (
	db "gopractice/simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//server serves http requests for the banking servcices
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.GET("/entries/:id", server.getEntry)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
