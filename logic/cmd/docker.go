package cmd

import (
	"fmt"
	"logic/logic"
	"os"

	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use: "docker",
	Run: func(cmd *cobra.Command, args []string) {
		asyncAddr := os.Getenv("ASYNC_ADDR")
		if asyncAddr == "" {
			panic(fmt.Errorf("Not Correct asyncAddr: %s\n", asyncAddr))
		}
		asyncName := os.Getenv("ASYNC_NAME")
		if asyncAddr == "" {
			panic(fmt.Errorf("Not Correct asyncName: %s\n", asyncName))
		}

		ramAddr := os.Getenv("RAM_ADDR")
		if ramAddr == "" {
			panic(fmt.Errorf("Not Correct ramAddr: %s\n", ramAddr))
		}

		logic := logic.Init(
			asyncAddr,
			asyncName,
			ramAddr,
		)

		logic.Run()
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
