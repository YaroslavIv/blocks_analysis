package cmd

import (
	"client/client"
	"fmt"
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

		ginAddr := os.Getenv("GIN_ADDR")
		if ginAddr == "" {
			panic(fmt.Errorf("Not Correct ginAddr: %s\n", ginAddr))
		}

		err := client.Init(asyncAddr, asyncName).Run(ginAddr)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
