package route

import (
	"fmt"

	ctl "lecture/go-contracts/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	cc *ctl.ContractController
}

/* 주문자, 피주문자 컨트롤러 할당 */
func NewRouter(_cc *ctl.ContractController) (*Router, error) {
	r := &Router{cc: _cc}
	return r, nil
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 임의 인증을 위한 함수
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*----------- 인증 프로세스 -----------*/
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		auth := c.GetHeader("Authorization")

		if auth != "codz" {
			//로직 추가 가능 현재는 Print 로만 처리
			fmt.Println("Authorization failed")
		}
		/*--------------- END -------------*/
		c.Next()
	}
}

func (p *Router) Idx() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	e := gin.Default()
	e.GET("/health")

	//[GET] 관련 라우팅
	view_contract := e.Group("api/v01", liteAuth())
	{
		view_contract.GET("/symbol/:tokenName", p.cc.GetTokenSymbol)
		view_contract.GET("/token-balance/:address", p.cc.GetTokenBalanceByAddress)

	}
	//[POST] 관련 라우팅
	post_contract := e.Group("api/v02", liteAuth())
	{
		post_contract.POST("/coin/:address", p.cc.CoinTransferTo)                      //코인 전송
		post_contract.POST("/coin-approve-transfer/:address", p.cc.CoinTransferFrom)   //개인키를 특정한 주소에 지정한 양의 받아서 코인 전송
		post_contract.POST("/token/:address", p.cc.TokenTransferFrom)                  //토큰 전송
		post_contract.POST("/token-approve-transfer/:address", p.cc.TokenTransferFrom) //다른 개인 키로 특정한 주소에 지정한 양의 토큰 전송

	}

	return e
}
