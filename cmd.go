package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/zhs007/jarvissh/basedef"
)

func startCmd() error {
	rootCmd := &cobra.Command{
		Use: "jarvissh",
	}

	//--------------------------------------------------------------------------------------------------------------------
	// start

	var daemon bool

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start jarvis shell",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				command := exec.Command("./jarvissh", "start")
				err := command.Start()
				if err != nil {
					fmt.Printf("start jarvissh error. %v \n", err)

					os.Exit(-1)
				}

				fmt.Printf("jarvissh start, [PID] %d running...\n", command.Process.Pid)
				ioutil.WriteFile("jarvissh.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)

				daemon = false

				os.Exit(0)
			} else {
				fmt.Printf("jarvissh start.\n")
			}

			startServ()
		},
	}

	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")

	rootCmd.AddCommand(startCmd)

	//--------------------------------------------------------------------------------------------------------------------
	// stop

	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop jarvis shell",
		Run: func(cmd *cobra.Command, args []string) {
			strb, err := ioutil.ReadFile("jarvissh.pid")
			if err != nil {
				fmt.Printf("read jarvissh.pid error %v\n", err)

				os.Exit(-1)
			}

			command := exec.Command("kill", string(strb))
			command.Start()
			fmt.Printf("jarvissh stop.\n")
		},
	}

	rootCmd.AddCommand(stopCmd)

	//--------------------------------------------------------------------------------------------------------------------
	// version

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "get jarvis shell version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("jarvis shell version is %v \n", basedef.VERSION)
		},
	}

	rootCmd.AddCommand(versionCmd)

	return rootCmd.Execute()
}
