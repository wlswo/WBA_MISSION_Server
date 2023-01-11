package services

import (
	"math/big"
)

type ContractService interface {
	GetTokenSymbolByTokenName(tokenName string) (symbol string, err error)             //토큰 심볼 조회
	GetTokenBalanceByAddress(from string) (balace *big.Int, err error)                 //토큰 잔액 조회
	CoinTransferTo(to string, value int64) (receipt string, err error)                 //코인 전송
	CoinTransferFrom(from string, to string, value int64) (receipt string, err error)  //개인키 계정으로부터 코인 전송
	TokenTransferTo(to string, value int64) (receipt string, err error)                //토큰 전송
	TokenTransferFrom(from string, to string, value int64) (receipt string, err error) //개인키 계정으로부터 토큰 전송
}
