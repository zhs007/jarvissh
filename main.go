package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/zhs007/jarviscore"
	pb "github.com/zhs007/jarviscore/proto"
)

func main() {
	fmt.Print("jarvis shell start...")

	var rundir string

	flag.StringVar(&rundir, "run", "./", "run path")
	flag.Parse()

	fmt.Print("jarvis shell runpath is " + rundir)

	mycfg, err := loadConfig(path.Join(rundir, "./config.yaml"))
	if err != nil {
		fmt.Print("load config.yaml fail!")

		return
	}

	cfg := jarviscore.Config{
		RunPath:     rundir,
		DefPeerAddr: mycfg.RootServAddr,
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
	defer node.Stop()

	node.Start()

	fmt.Print("jarvis shell end.")
}
