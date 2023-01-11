package main

import (
	"context"
	"flag"
	"fmt"
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

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {

	var configFlag = flag.String("config", "./config/config.toml", "toml file to use for configuration")
	flag.Parse()
	cf := conf.NewConfig(*configFlag)

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
