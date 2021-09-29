package main

import "github.com/spf13/cobra"

type globalFlag struct {
	Verbose bool
}

func initGlobalFlag(cmd *cobra.Command, f *globalFlag) {
	cmd.PersistentFlags().BoolVarP(&f.Verbose, "verbose", "v", false, "display debug output")
}
