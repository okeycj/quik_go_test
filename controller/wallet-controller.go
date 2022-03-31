package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/okeycj/quik_go_test/dto"
	"github.com/okeycj/quik_go_test/helper"
	"github.com/okeycj/quik_go_test/services"
	"github.com/sirupsen/logrus"
)

var logg = logrus.New()

type WalletController interface {
	GetWalletBalance(ctx *gin.Context)
	CreditWallet(ctx *gin.Context)
	DebitWallet(ctx *gin.Context)
}

type walletController struct {
	walletService services.WalletService
	jwtService    services.JWTService
}

func NewWalletController(walletService services.WalletService, jwtService services.JWTService) WalletController {
	return &walletController{
		walletService: walletService,
		jwtService:    jwtService,
	}
}

func (c *walletController) GetWalletBalance(ctx *gin.Context) {
	id := ctx.Param("id")
	logg.WithFields(logrus.Fields{"id": id}).Info("id param from " + ctx.Request.URL.Path)
	if !c.walletService.WalletAssists(id) {
		response := helper.BuildErrorResponse("Unable to get wallet balance", "Wallet not found", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	} else {
		authHeader := ctx.GetHeader("Authorization")
		token, _ := c.jwtService.ValidateToken(authHeader)
		claims := token.Claims.(jwt.MapClaims)
		user_id := claims["user_id"].(string)
		u_id, _ := strconv.ParseUint(string(user_id), 0, 64)
		getWallet := c.walletService.GetWalletBalance(id)
		if getWallet.User.ID != u_id {
			response := helper.BuildErrorResponse("Unable to get wallet balance", "You are not authorized to view wallet", helper.EmptyObj{})
			ctx.JSON(http.StatusUnauthorized, response)
			return
		}
		response := helper.BuildResponse(true, "OK!", getWallet)
		ctx.JSON(http.StatusFound, response)
	}
}

func (c *walletController) CreditWallet(ctx *gin.Context) {
	id := ctx.Param("id")
	var creditWalletDTO dto.CreditOrDebitWalletDTO
	errDTO := ctx.ShouldBind(&creditWalletDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	logg.WithFields(logrus.Fields{"id": id, "body.amount": creditWalletDTO.Amount}).Info("id param and request body from " + ctx.Request.URL.Path)
	if creditWalletDTO.Amount.IsNegative() {
		response := helper.BuildErrorResponse("Failed to process request", "Amount must not be less than zero", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.walletService.WalletAssists(id) {
		response := helper.BuildErrorResponse("Unable to find wallet", "Wallet not found", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	} else {
		creditWallet := c.walletService.CreditWallet(id, creditWalletDTO.Amount)
		if creditWallet == nil {
			response := helper.BuildErrorResponse("Unable to credit wallet", "", helper.EmptyObj{})
			ctx.JSON(http.StatusInternalServerError, response)
		}
		response := helper.BuildResponse(true, "Successfully credited", creditWallet)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *walletController) DebitWallet(ctx *gin.Context) {
	id := ctx.Param("id")
	var debitWalletDTO dto.CreditOrDebitWalletDTO
	errDTO := ctx.ShouldBind(&debitWalletDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	logg.WithFields(logrus.Fields{"id": id, "body.amount": debitWalletDTO.Amount}).Info("id param and request body from " + ctx.Request.URL.Path)
	if debitWalletDTO.Amount.IsNegative() {
		response := helper.BuildErrorResponse("Failed to process request", "Amount must not be less than zero", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.walletService.WalletAssists(id) {
		response := helper.BuildErrorResponse("Unable to find wallet", "Wallet not found", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	} else {
		debitWallet := c.walletService.DebitWallet(id, debitWalletDTO.Amount)
		if debitWallet == nil {
			response := helper.BuildErrorResponse("Unable to debit wallet", "", helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.BuildResponse(true, "Successfully Debited", debitWallet)
		ctx.JSON(http.StatusOK, response)
		return
	}
}
