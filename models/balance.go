package models

type RequestValue struct {
	Value int64 `json:"value"`
}

type RequestTransferFrom struct {
	Value      int64  `json:"value"`
	PrivateKey string `json:"privatekey"`
}
