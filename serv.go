package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvissh/basedef"
)

func startServ() {
	fmt.Printf("jarvis shell start...\n")
	fmt.Printf("jarvis shell version is %v \n", basedef.VERSION)

	cfg, err := jarviscore.LoadConfig("cfg/config.yaml")
	if err != nil {
		fmt.Printf("load config.yaml fail!\n")

		return
	}

	jarviscore.InitJarvisCore(cfg, basedef.JARVISNODETYPE, basedef.VERSION)
	defer jarviscore.ReleaseJarvisCore()

	// pprof
	jarviscore.InitPprof(cfg)

	node, err := jarviscore.NewNode(cfg)
	if err != nil {
		fmt.Printf("jarviscore.NewNode fail! %v \n", err)

		return
	}

	node.SetNodeTypeInfo(basedef.JARVISNODETYPE, basedef.VERSION)

	node.Start(context.Background())

	fmt.Printf("jarvis shell end.\n")
}
