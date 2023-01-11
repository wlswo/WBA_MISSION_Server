package controllers

import (
	"lecture/go-contracts/models"
	"lecture/go-contracts/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContractController struct {
	ContractService services.ContractService
}

func NewContractController(contractservice services.ContractService) (ContractController, error) {
	return ContractController{
		ContractService: contractservice,
	}, nil
}
func (cc *ContractController) GetTokenSymbol(c *gin.Context) {
	symbol, err := cc.ContractService.GetTokenSymbolByTokenName(c.Param("tokenName"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"symbol": symbol})
}

func (cc *ContractController) GetTokenBalanceByAddress(c *gin.Context) {
	balance, err := cc.ContractService.GetTokenBalanceByAddress(c.Param("address"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

/* 코인 전송 */
func (cc *ContractController) CoinTransferTo(c *gin.Context) {
	var value models.RequestValue
	/* BINDING */
	if err := c.ShouldBindJSON(&value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receipt, err := cc.ContractService.CoinTransferTo(c.Param("address"), value.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Tx receipt": receipt})
}

/* 개인키 받아 해당 계정으로부터 코인 전송 */
func (cc *ContractController) CoinTransferFrom(c *gin.Context) {
	var account models.RequestTransferFrom
	/* BINDING */
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	receipt, err := cc.ContractService.CoinTransferFrom(account.PrivateKey, c.Param("address"), account.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Tx receipt": receipt})
}

/* 토큰 전송 */
func (cc *ContractController) TokenTransferTo(c *gin.Context) {
	var value models.RequestValue
	/* BINDING */
	if err := c.ShouldBindJSON(&value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receipt, err := cc.ContractService.CoinTransferTo(c.Param("address"), value.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Tx receipt": receipt})
}

/* 개인키 받아 해당 계정으로부터 토큰 전송 */
func (cc *ContractController) TokenTransferFrom(c *gin.Context) {
	var account models.RequestTransferFrom
	/* BINDING */
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	receipt, err := cc.ContractService.CoinTransferFrom(account.PrivateKey, c.Param("address"), account.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Tx receipt": receipt})
}
