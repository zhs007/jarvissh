package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
	pb "github.com/zhs007/jarviscore/proto"
)

func main() {
	fmt.Print("jarvis shell start...")

	mycfg, err := loadConfig("cfg/config.yaml")
	if err != nil {
		fmt.Print("load config.yaml fail!")

		return
	}

	cfg := jarviscore.Config{
		DBPath:         mycfg.DBPath,
		LogPath:        mycfg.LogPath,
		DefPeerAddr:    mycfg.RootServAddr,
		AnkaDBHttpServ: mycfg.AnkaDB.HTTPServ,
		AnkaDBEngine:   mycfg.AnkaDB.Engine,
		LogConsole:     mycfg.LogConsole,
		LogLevel:       mycfg.LogLevel,
	}

	myinfo := jarviscore.BaseInfo{
		Name:     mycfg.NodeName,
		BindAddr: mycfg.BindAddr,
		ServAddr: mycfg.ServAddr,
		NodeType: pb.NODETYPE_SH,
	}

	jarviscore.InitJarvisCore(cfg)
	defer jarviscore.ReleaseJarvisCore()

	node := jarviscore.NewNode(myinfo)
	// defer node.Stop()

	node.Start(context.Background())

	fmt.Print("jarvis shell end.")
}
