package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvissh/basedef"
)

func main() {
	fmt.Print("jarvis shell start...")
	fmt.Print("jarvis shell version is " + basedef.VERSION)

	cfg, err := jarviscore.LoadConfig("cfg/config.yaml")
	if err != nil {
		fmt.Print("load config.yaml fail!")
		// fmt.Print("load config.yaml fail!")

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
	node.SetNodeTypeInfo(basedef.JARVISNODETYPE, basedef.VERSION)
	// defer node.Stop()

	node.Start(context.Background())

	fmt.Print("jarvis shell end.")
}
