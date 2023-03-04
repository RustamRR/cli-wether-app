/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/RustamRR/cli-wether-app/configs"
	"github.com/RustamRR/cli-wether-app/internal/app"
	"github.com/spf13/cobra"
)

// migrationsCmd represents the migrations command
var migrationsCmd = &cobra.Command{
	Use:   "migrations",
	Short: "Выполнить миграции",
	Run: func(cmd *cobra.Command, args []string) {
		server := app.New(configs.GetConfig())
		server.Migrate()
		fmt.Println("Миграции выполнены")
	},
}

func init() {
	rootCmd.AddCommand(migrationsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
