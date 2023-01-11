package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	conf "lecture/go-contracts/config"
	ctl "lecture/go-contracts/controllers"
	"lecture/go-contracts/logger"
	rt "lecture/go-contracts/router"
	"lecture/go-contracts/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"golang.org/x/sync/errgroup"
	"golang.org/x/term"
)

var (
	g errgroup.Group
)

func main() {

	var configFlag = flag.String("config", "./config/config.toml", "toml file to use for configuration")
	flag.Parse()
	cf := conf.NewConfig(*configFlag)

	/* 개인키 검증 Password 입력 */
	keyStore, _ := ioutil.ReadFile(cf.KeyStore.Fpath)
	fmt.Print("password : ")

	if password, err := term.ReadPassword(0); err != nil {
		panic(err)
	} else if key, err := keystore.DecryptKey(keyStore, string(password)); err != nil {
		panic(err)
	} else {
		cf.Address.PrivateKey = hex.EncodeToString(key.PrivateKey.D.Bytes())
	}

	/* 로그 초기화 */
	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	logger.Debug("ready server....")

	if ContractService, err := services.NewContractService(context.TODO(), cf); err != nil {
		panic(err)
	} else if oc, err := ctl.NewContractController(ContractService); err != nil {
		panic(err)
	} else if rt, err := rt.NewRouter(&oc); err != nil {
		panic(fmt.Errorf("router.NewRouter > %v", err))
	} else {
		/* Server 설정 */
		mapi := &http.Server{
			Addr:           cf.Server.Port,
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Warn("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			logger.Error("Server Shutdown:", err)
		}

		select {
		case <-ctx.Done():
			logger.Info("timeout of 1 seconds.")
		}

		logger.Info("Server exiting")
	}
	if err := g.Wait(); err != nil {
		logger.Error(err)
	}

}
