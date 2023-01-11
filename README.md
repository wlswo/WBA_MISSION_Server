# Go X BlockChain NewWork

Blockchain network communication using Go
## Installation



```bash
$ git clone https://github.com/wlswo/WBA_MISSION_Server.git
```

## Usage

### Make config.toml
```bash
$ cd config/
$ touch config.toml
```

```
[log]
level = "debug" # debug or info
fpath = "./logs/go-loger" 
msize = 2000    # 2g : megabytes
mage = 7        # 7days
mbackup = 5     # number of log files

[server] ##nomal type
mode = "dev"
port = ":8080"

[keystore]
fpath = "config/yourKeyStoreName.json"

[contract]
contractaddress = "yourContractAddress"

```

### Save Your Private KeyStore in config Folder
[Create keystore
](https://github.com/miguelmota/go-ethereum-hdwallet)

### start server
```bash
$ go run main.go
> password : {input your password}

# success 
> 2023/01/12 03:20:24 Server Start..

# failurl
> panic: could not decrypt key with given password
```



