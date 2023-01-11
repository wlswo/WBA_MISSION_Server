package services

import (
	"context"
	conf "lecture/go-contracts/config"
	utils "lecture/go-contracts/utils"
	"math/big"
)

type ContractServiceImplement struct {
	ctx context.Context
	cfg *conf.Config
}

func NewContractService(ctx context.Context, cfg *conf.Config) (ContractService, error) {
	return &ContractServiceImplement{
		ctx: ctx,
		cfg: cfg,
	}, nil
}

/* 토큰 심볼 조회 */
func (o *ContractServiceImplement) GetTokenSymbolByTokenName(tokenName string) (symbol string, err error) {
	return utils.GetTokenSymbol(o.cfg.Contract.ContractAddress, tokenName)
}

/* 토큰 잔액 조회 */
func (o *ContractServiceImplement) GetTokenBalanceByAddress(from string) (balance *big.Int, err error) {
	return utils.GetTokenBalaceByAddress(o.cfg.Contract.ContractAddress, from)
}

/* 코인 전송 */
func (o *ContractServiceImplement) CoinTransferTo(to string, value int64) (receipt string, err error) {
	return utils.TransferWemix(o.cfg.Address.PrivateKey, to, value)
}

/* 개인키 계정으로 부터 코인 전송 */
func (o *ContractServiceImplement) CoinTransferFrom(from string, to string, value int64) (receipt string, err error) {
	return utils.TransferWemix(from, to, value)
}

/* 토큰 전송 */
func (o *ContractServiceImplement) TokenTransferTo(to string, value int64) (receipt string, err error) {
	return utils.TransferCtxBjj(o.cfg.Address.PrivateKey, to, value)
}

/* 개인키 계정으로 부터 토큰 전송 */
func (o *ContractServiceImplement) TokenTransferFrom(from string, to string, value int64) (receipt string, err error) {
	return utils.TransferCtxBjj(from, to, value)
}
