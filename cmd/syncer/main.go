package main

import (
	"github.com/juxuny/log-server/log"
	"github.com/spf13/cobra"
	"os"
)

var logger = log.NewLogger("syncer")

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "root",
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
}
