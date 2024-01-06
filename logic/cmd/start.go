package cmd

import (
	"logic/logic"

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
		name, err := cmd.Flags().GetString("async-name")
		if err != nil {
			panic(err)
		}

		ramAddr, err := cmd.Flags().GetString("arm-addr")
		if err != nil {
			panic(err)
		}
		if ramAddr == "" {
			panic("ramAddr empty")
		}

		logic := logic.Init(asyncAddr, name, ramAddr)

		logic.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("arm-addr", "u", "localhost:6379", "url RabbitMQ")
	startCmd.Flags().StringP("async-addr", "a", "amqp://guest:guest@localhost:5672/", "addr RabbitMQ")
	startCmd.Flags().StringP("async-name", "n", "top", "name RabbitMQ")
}
