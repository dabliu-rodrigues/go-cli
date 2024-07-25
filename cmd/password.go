/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jsGolden/go-cli/utils"
	"github.com/spf13/cobra"
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("s")
		cmd.Println(utils.GeneratePassword(size))
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
	passwordCmd.PersistentFlags().Int("s", 16, "Password size to be generated. Default: 16")
}
