package cmd

import (
	"os"
	"path"
	"strings"

	"github.com/phamvinhdat/mo/env"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   ".env",
	Short: ".env env format",
	Run:   envCmdRun,
}

func init() {
	rootCmd.AddCommand(envCmd)
}

func envCmdRun(cmd *cobra.Command, _ []string) {
	if len(fileExtension) == 0 {
		ext := path.Ext(file)
		if len(ext) == 0 {
			cmd.PrintErrln("config file extension is require if file path don't contain extension")
			return
		}

		fileExtension = ext
	}

	fileExtension = strings.TrimLeft(fileExtension, ".")
	data, err := os.ReadFile(file)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	if err = env.PrintEnv(data, fileExtension, prefix); err != nil {
		cmd.PrintErrln(err)
		return
	}
}
