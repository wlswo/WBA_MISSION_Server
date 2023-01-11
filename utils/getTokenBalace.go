package utils

import (
	"fmt"
	"lecture/go-contracts/contracts"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTokenBalaceByAddress(_contractAddress string, _address string) (balance *big.Int, err error) {
	// 블록체인 네트워크와 연결할 클라이언트를 생성하기 위한 rpc url 연결
	client, err := ethclient.Dial("https://api.test.wemix.com")
	if err != nil {
		fmt.Println("client error")
	}

	// 본인이 배포한 토큰 컨트랙트 어드레스
	tokenAddress := common.HexToAddress(_contractAddress)
	instance, err := contracts.NewContracts(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	// 타겟 어드레스
	address := common.HexToAddress(_address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		fmt.Println(err)
	}

	return bal, nil
}
