package cmd

import (
	"fmt"
	"os"
	"strconv"
	"sync_logic/node"

	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use: "docker",
	Run: func(cmd *cobra.Command, args []string) {

		wss := os.Getenv("WSS")
		if wss == "" {
			panic(fmt.Errorf("Not Correct wss: %s\n", wss))
		}

		nameTable := os.Getenv("NAME_TABLE")
		if nameTable == "" {
			panic(fmt.Errorf("Not Correct nameTable: %s\n", nameTable))
		}

		ramAddr := os.Getenv("RAM_ADDR")
		if ramAddr == "" {
			panic(fmt.Errorf("Not Correct ramAddr: %s\n", ramAddr))
		}

		dbAddr := os.Getenv("DB_ADDR")
		if dbAddr == "" {
			panic(fmt.Errorf("Not Correct dbAddr: %s\n", dbAddr))
		}
		dbDatabase := os.Getenv("DB_DATABASE")
		if dbDatabase == "" {
			panic(fmt.Errorf("Not Correct dbDatabase: %s\n", dbDatabase))
		}
		dbUsername := os.Getenv("DB_USERNAME")
		if dbUsername == "" {
			panic(fmt.Errorf("Not Correct dbUsername: %s\n", dbUsername))
		}
		dbPassword := os.Getenv("DB_PASSWORD")

		maxCountBlock, err := strconv.ParseUint(os.Getenv("MAX_COUNT_BLOCK"), 10, 64)
		if err != nil {
			panic(err)
		}
		if maxCountBlock < 1 {
			panic(fmt.Errorf("Not Correct maxCountBlock: %d\n", maxCountBlock))
		}

		eth := node.Init(
			node.ETH,
			wss,
			maxCountBlock,
			nameTable,
			ramAddr,
			dbAddr,
			dbDatabase,
			dbUsername,
			dbPassword,
		)

		if err := eth.Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
