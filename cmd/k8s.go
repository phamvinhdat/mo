package cmd

import (
	"os"
	"path"
	"strings"

	"github.com/phamvinhdat/mo/k8s"

	"github.com/spf13/cobra"
)

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "k8s env format",
	Run:   k8sCmdRun,
}

func init() {
	rootCmd.AddCommand(k8sCmd)
}

func k8sCmdRun(cmd *cobra.Command, _ []string) {
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

	if err = k8s.PrintEnv(data, fileExtension, prefix, arraySeparator); err != nil {
		cmd.PrintErrln(err)
		return
	}
}
