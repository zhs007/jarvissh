package main

import (
	"context"

	"go.uber.org/zap"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarviscore/base"
)

func main() {
	jarvisbase.Info("jarvis shell start...")
	jarvisbase.Info("jarvis shell version is 0.1.3")

	cfg, err := jarviscore.LoadConfig("cfg/config.yaml")
	if err != nil {
		jarvisbase.Warn("load config.yaml fail!", zap.Error(err))
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
	// defer node.Stop()

	node.Start(context.Background())

	jarvisbase.Info("jarvis shell end.")
}
