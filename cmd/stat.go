package cmd

import (
    "github.com/spf13/cobra"
)

var statCmd = &cobra.Command{
    Use: "stat",
    Short: "show statistics about data",
    Run: showStat,
}

func showStat(cmd *cobra.Command, args []string) {
    containerCycle(args, prepStat)
}

func prepStat(args []string) []string {
    cmd := make([]string, 3)
    cmd[0] = "stat"
    cmd[1] = "--uri"
    cmd[2] = datasetURI
    cmd = append(cmd, args...)
    return cmd
}
