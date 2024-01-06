package cmd

import (
	"fmt"
	"sync_logic/node"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		wss, err := cmd.Flags().GetString("wss")
		if err != nil || wss == "" {
			panic(fmt.Errorf("Not Correct wss: %s\n", wss))
		}

		maxCountBlock, err := cmd.Flags().GetUint64("max-count-block")
		if err != nil {
			panic(err)
		}

		nameTable, err := cmd.Flags().GetString("name-table")
		if err != nil {
			panic(err)
		}

		ramAddr, err := cmd.Flags().GetString("ram-addr")
		if err != nil {
			panic(err)
		}

		dbAddr, err := cmd.Flags().GetString("db-addr")
		if err != nil {
			panic(err)
		}
		dbDatabase, err := cmd.Flags().GetString("db-database")
		if err != nil {
			panic(err)
		}
		dbUsername, err := cmd.Flags().GetString("db-username")
		if err != nil {
			panic(err)
		}
		dbPassword, err := cmd.Flags().GetString("db-password")
		if err != nil {
			panic(err)
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

// TODO Подредактировать описание флагов
func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("wss", "w", "", "wss to BlockChain")
	startCmd.Flags().Uint64("max-count-block", 100, "max Count Block")
	startCmd.Flags().StringP("name-table", "n", "eth", "name table in SQL")
	startCmd.Flags().StringP("ram-addr", "r", "localhost:6379", "ram addr")

	startCmd.Flags().String("db-addr", "localhost:9000", "database addr")
	startCmd.Flags().String("db-database", "default", "database database")
	startCmd.Flags().String("db-username", "default", "database username")
	startCmd.Flags().String("db-password", "", "database password")
}
