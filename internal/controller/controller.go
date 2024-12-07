package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mizmorr/wallet/internal/model"
	"github.com/mizmorr/wallet/internal/service"
	"github.com/mizmorr/wallet/pkg/apperror"
)

type Service interface {
	Deposit(ctx context.Context, wallet *model.WalletRequest) error
	Get(ctx context.Context, id uuid.UUID) (*model.WalletResponse, error)
	Withdraw(ctx context.Context, wallet *model.WalletRequest) error
}

var (
	_     Service = (*service.WalletService)(nil)
	valid         = validator.New()
)

type WalletController struct {
	ctx context.Context
	svc Service
}

func NewWalletController(svc Service, ctx context.Context) *WalletController {
	return &WalletController{
		svc: svc,
		ctx: ctx,
	}
}

func (c *WalletController) Get(g *gin.Context) {
	userid_raw, ok := g.Params.Get("id")
	if !ok {
		respondWithError(g, apperror.ErrBadRequest, "No id provided")
		return
	}

	userid, err := uuid.Parse(userid_raw)
	if err != nil {
		respondWithError(g, apperror.ErrValidation, "Could not parse userid")
		return
	}

	user, err := c.svc.Get(c.ctx, userid)
	if err != nil {
		respondWithError(g, apperror.ErrNotFound, "Could not find user with that id")
		return
	}
	g.JSON(http.StatusOK, user)
}

func (c *WalletController) Operate(g *gin.Context) {
	walletToOperate := &model.WalletRequest{}

	err := g.Bind(walletToOperate)
	if err != nil {
		log.Println(walletToOperate)
		respondWithError(g, apperror.ErrBadRequest, "Invalid ID!")
		return
	}
	err = valid.Struct(walletToOperate)
	if err != nil {
		respondWithError(g, apperror.ErrValidation, err.Error())
		return
	}

	switch walletToOperate.Operation {
	case "deposit":
		err = c.svc.Deposit(c.ctx, walletToOperate)
	case "withdraw":
		err = c.svc.Withdraw(c.ctx, walletToOperate)
	}
	if err != nil {
		respondWithError(g, apperror.ErrForbidden, err.Error())
		return
	}

	g.JSON(http.StatusOK, "Success!")
}

func respondWithError(g *gin.Context, err apperror.AppError, details string) {
	err.Details = details
	g.AbortWithStatusJSON(err.StatusCode, err)
}
