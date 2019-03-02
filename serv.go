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

	jarviscore.InitJarvisCore(cfg)
	defer jarviscore.ReleaseJarvisCore()

	node := jarviscore.NewNode(cfg)
	node.SetNodeTypeInfo(basedef.JARVISNODETYPE, basedef.VERSION)

	node.Start(context.Background())

	fmt.Printf("jarvis shell end.\n")
}
