package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"theaveasso.bab/internal/db"
	"theaveasso.bab/internal/utility"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=8,max=72"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
}

type createUserResponse struct {
    Username string `json:"username"`
    Fullname string `json:"fullname"`
    Email string `json:"email"`
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required,min=1"`
    Password string `uri:"password" binding:"required,min=8"`
}

type listUsersRequest struct {
	PageID   int32 `form:"page_id"   binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utility.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.Fullname,
		Email:          req.Email,
	}
	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		switch {
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

    resp := createUserResponse{
        Email: user.Email,
        Fullname: user.FullName,
        Username: req.Username,
    }

	ctx.JSON(http.StatusOK, resp)
}

func (s *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := s.store.GetUser(ctx, req.Username)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
	}

	ctx.JSON(http.StatusOK, account)
}

func (s *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := s.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, accounts)
}
