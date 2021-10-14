package api

import (
	"database/sql"
	"log"
	"net/http"
	db "simple-bank/db/sqlc"

	"github.com/gin-gonic/gin"
)


type CreateAccountRequest struct {
	Owner string `json:"owner" binding:"required"` //here start with capital because it will accesed in another package
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

//create account handler
func (server *Server) CreateAccount(ctx *gin.Context){
	//here getting the account details
	var req CreateAccountRequest;
	err := ctx.ShouldBindJSON(&req)

	if err!=nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err!=nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return;
	}

	ctx.JSON(200,account)
}

//get account by ID, handler
type GetAccountRequest struct{
	ID int64 `uri:"id" binding:"required"`
}
func (server *Server) GetAccount(ctx *gin.Context){
	var req GetAccountRequest;

	err := ctx.ShouldBindUri(&req); if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	log.Print(req);

	account, err := server.store.GetAccount(ctx, req.ID);

	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(404, errorResponse(err));
			return;
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return;
	}

	ctx.JSON(200, account)
}


type ListAccountsRequest struct {
	PageID int32 `form:"pageId" binding:"required,min=1"`
	PageSize int32 `form:"pageSize" binding:"required,min=5,max=10"`
}
func (server *Server) ListAccounts(ctx *gin.Context){
	var req ListAccountsRequest;

	err := ctx.ShouldBindQuery(&req);
	log.Print(req);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams {
		Limit: req.PageSize,
		Offset: (req.PageID-1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return;
	}

	ctx.JSON(200, accounts)

}