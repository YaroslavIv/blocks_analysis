package cmd

import (
	"client/client"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		asyncAddr, err := cmd.Flags().GetString("async-addr")
		if err != nil {
			panic(err)
		}
		if asyncAddr == "" {
			panic("asyncAddr empty")
		}

		asyncName, err := cmd.Flags().GetString("async-name")
		if err != nil {
			panic(err)
		}

		ginAddr, err := cmd.Flags().GetString("gin-addr")
		if err != nil {
			panic(err)
		}

		err = client.Init(asyncAddr, asyncName).Run(ginAddr)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("async-addr", "a", "amqp://guest:guest@localhost:5672/", "url RabbitMQ")
	startCmd.Flags().StringP("async-name", "n", "top", "name RabbitMQ")
	startCmd.Flags().StringP("gin-addr", "g", "localhost:8080", "addr gin server")
}
