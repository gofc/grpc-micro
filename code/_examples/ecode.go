package _examples

import (
	"github.com/gofc/grpc-micro/code"
)

//Stringメソッドを自動生成
//go:generate stringer -type=errorCode

type errorCode code.Code

const (
	Internal errorCode = 100000
)
