package init

import "C"
import (
	"github.com/spf13/cobra"
	"os"
)

var Cli = &cobra.Command{
	Use:     "go-scheduler",
	Short:   "go-scheduler",
	Version: "0.0.1",
}

func Execute() {
if err := Cli.Execute(); err != nil {
		os.Exit(1)
	}
}
