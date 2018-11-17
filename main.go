package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
)

func main() {
	fmt.Print("jarvis shell start...")

	cfg, err := jarviscore.LoadConfig("cfg/config.yaml")
	if err != nil {
		fmt.Print("load config.yaml fail!")

		return
	}

	// cfg := jarviscore.Config{
	// 	DBPath:         mycfg.DBPath,
	// 	LogPath:        mycfg.LogPath,
	// 	DefPeerAddr:    mycfg.RootServAddr,
	// 	AnkaDBHttpServ: mycfg.AnkaDB.HTTPServ,
	// 	AnkaDBEngine:   mycfg.AnkaDB.Engine,
	// 	LogConsole:     mycfg.LogConsole,
	// 	LogLevel:       mycfg.LogLevel,
	// }

	// bi := &jarviscore.BaseInfo{
	// 	Name:     cfg.BaseNodeInfo.NodeName,
	// 	BindAddr: cfg.BaseNodeInfo.BindAddr,
	// 	ServAddr: cfg.BaseNodeInfo.ServAddr,
	// 	// NodeType: pb.NODETYPE_SH,
	// }

	jarviscore.InitJarvisCore(cfg)
	defer jarviscore.ReleaseJarvisCore()

	node := jarviscore.NewNode(cfg)
	// defer node.Stop()

	node.Start(context.Background())

	fmt.Print("jarvis shell end.")
}
