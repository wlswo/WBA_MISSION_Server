package config

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Server struct {
		Mode string
		Port string
	}
	Log struct {
		Level   string
		Fpath   string
		Msize   int
		Mage    int
		Mbackup int
	}
	KeyStore struct {
		Fpath string
	}
	Contract struct {
		ContractAddress string
	}
	Address struct {
		PrivateKey string
	}
}

func NewConfig(fpath string) *Config {
	c := new(Config)

	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		//toml 파일 디코딩
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			//fmt.Println(c)
			return c
		}
	}
}
