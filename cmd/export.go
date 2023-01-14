package cmd

import (
    "github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
    Use: "export",
    Short: "export vcf,bcf,vcf.gz,tsv, ... from the dataset",
    Run: export,
}

func export(cmd *cobra.Command, args []string) {
    containerCycle(args, prepExport)
}

func prepExport(args []string) []string {
    cmd := make([]string, 3)
    cmd[0] = "export"
    cmd[1] = "--uri"
    cmd[2] = datasetURI
    cmd = append(cmd, args...)

    return cmd
}
