package main

import (
	"errors"
	_ "net/http/pprof"

	"github.com/zhs007/jarviscore/base"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			err, isok := p.(error)
			if isok {
				jarvisbase.DumpPanic(err)
			} else {
				jarvisbase.DumpPanic(errors.New("invalid error"))
			}
		}
	}()

	startCmd()
}
