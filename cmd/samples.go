package cmd

import (
    "github.com/spf13/cobra"
)

var samplesCmd = &cobra.Command{
    Use: "samples",
    Short: "lists samples in the specified dataset",
    Run: listSamples,
}

func listSamples(cmd *cobra.Command, args []string) {
    containerCycle(args, prepList)
}

func prepList(args []string) []string {
    cmd := make([]string, 3)
    cmd[0] = "list"
    cmd[1] = "--uri"
    cmd[2] = datasetURI
    cmd = append(cmd, args...)
    return cmd
}
