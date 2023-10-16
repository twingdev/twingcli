/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/twingdev/twingcli/server"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "p2p node - server commands",
	Long:  `"p2p node - server commands"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {

	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	serverCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start node server",
		Long:  "Start node server",
		Run: func(cmd *cobra.Command, arg []string) {
			startServer()
		},
	})

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		server.NewNode(ctx)
	}()
	select {}
}
