/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jsGolden/go-cli/utils"
	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Util to encode/decode strings.",
	Long: `Util to encode/decode strings.
		Example:
		encode: ./devops base64 --e "string"
		decode: ./devops base64 --d "string"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		encodeStr, _ := cmd.Flags().GetString("e")
		decodeStr, _ := cmd.Flags().GetString("d")

		if encodeStr != "" {
			encode := utils.EncodeString(encodeStr)
			cmd.Println(encode)
		} else if decodeStr != "" {
			decode := utils.DecodeString(decodeStr)
			cmd.Println(decode)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.PersistentFlags().String("e", "", "Encode string")
	base64Cmd.PersistentFlags().String("d", "", "Decode string")
}
