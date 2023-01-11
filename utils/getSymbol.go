package utils

import (
	"errors"
	"lecture/go-contracts/contracts"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTokenSymbol(_contractAddress string, _tokenName string) (string, error) {
	// 블록체인 네트워크와 연결할 클라이언트를 생성하기 위한 rpc url 연결
	client, err := ethclient.Dial("https://api.test.wemix.com")
	if err != nil {
		return "", errors.New("client error")
	}

	// 본인이 배포한 토큰 컨트랙트 어드레스
	tokenAddress := common.HexToAddress(_contractAddress)
	instance, err := contracts.NewContracts(tokenAddress, client)
	if err != nil {
		return "", err
	}
	//토큰 이름 비교
	if name, err := instance.Name(&bind.CallOpts{}); err != nil || strings.Split(name, " ")[0] != _tokenName {
		return "", errors.New("토큰 심볼을 가져올수 없습니다")
	}

	// symbol 출력
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", err
	}

	return symbol, nil

}
