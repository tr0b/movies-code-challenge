package api

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// "github.com/go-playground/validator/v10"

	db "github.com/tr0b/movies-code-challenge/backend/db/sqlc"
	// "github.com/tr0b/movies-code-challenge/backend/token"
	"github.com/tr0b/movies-code-challenge/backend/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config:     config,
		store:      store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/movie", server.createMovie)
	router.PATCH("/movie/:id", server.updateMovie)
	router.GET("/movie/", server.listMovies)
	router.GET("/movie/:id", server.getMovie)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
