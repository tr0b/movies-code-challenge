package api

import (
	"database/sql"
	// "errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	// "github.com/tr0b/movies-code-challenge/backend/token"
)

type createMovieRequest struct {
	Title string `json:"title" binding:"required"`
}

func (server *Server) createMovie(ctx *gin.Context) {
	var req createMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	movie, err := server.store.CreateMovie(ctx, req.Title)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movie)

}

type getMovieRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getMovie(ctx *gin.Context) {
	var req getMovieRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	movie, err := server.store.GetMovie(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Adds authorization based on user in token payload [not for this project :)]
	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// if movie.Owner != authPayload.Username {
	// 	err := errors.New("movie does not belong to user")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK, movie)

}

func (server *Server) listMovies(ctx *gin.Context) {
	movies, err := server.store.ListMovies(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movies)

}

type updateMovieRequestJSON struct {
	Action string `json:"action" binding:"required,oneof=UPVOTE"`
}
type updateMovieRequestURI struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) updateMovie(ctx *gin.Context) {
	var json updateMovieRequestJSON
	var uri updateMovieRequestURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	movies, err := server.store.UpdateMovie(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, movies)

}

