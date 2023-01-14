package cmd 

import (
    "github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
    Use: "create",
    Short: "Create a vcf dataset",
    Run: createDataset,
}

func createDataset(cmd *cobra.Command, args []string) {
    containerCycle(args, prepCreate)
}

func prepCreate(args []string) []string {
    cmd := make([]string, 3)
    cmd[0] = "create"
    cmd[1] = "--uri"
    cmd[2] = datasetURI
    cmd = append(cmd, args...)
    return cmd
}
