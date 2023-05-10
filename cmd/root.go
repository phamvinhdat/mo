package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mo",
}

var (
	file           string
	prefix         string
	fileExtension  string
	arraySeparator string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "config file. ex: ./config.yaml")
	rootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "env's prefix")
	rootCmd.PersistentFlags().StringVarP(&fileExtension, "extension", "e", "", "config file extension, require if file path don't contain extension")
	rootCmd.PersistentFlags().StringVarP(&arraySeparator, "array_separator", "s", " ", "array separated")

	err := rootCmd.MarkPersistentFlagRequired("file")
	if err != nil {
		rootCmd.PrintErr(err)
	}
}
