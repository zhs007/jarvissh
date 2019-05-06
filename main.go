package main

import (
	_ "net/http/pprof"
)

func main() {
	// defer func() {
	// 	if p := recover(); p != nil {
	// 		err, isok := p.(error)
	// 		if isok {
	// 			jarvisbase.DumpPanic(err)
	// 		} else {
	// 			jarvisbase.DumpPanic(errors.New("invalid error"))
	// 		}
	// 	}
	// }()

	startCmd()
}
