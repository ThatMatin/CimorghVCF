package cmd

import (
    "github.com/spf13/cobra"
)

var samples string

var ingestCmd = &cobra.Command{
    Use: "ingest",
    Short: "ingest samples into existing database",
    Long: `specify and input path containing everything related to input (*.bcf/vcf files, samples.txt 
    containing name of the samples). Also specify output path to point to the directory containing a database`,
    Run: ingestSamples,
}

func init() {
    ingestCmd.Flags().StringVarP(&samples,"samples", "s", "", "a text file with one sample name per line")
}

func ingestSamples(cmd *cobra.Command, args []string) {
    containerCycle(args, prepIngest)
}

func prepIngest(args []string) []string {
    cmd := make([]string, 3)
    cmd[0] = "store"
    cmd[1] = "--uri"
    cmd[2] = datasetURI
    if samples != "" {
        cmd = append(cmd, "--samples-file", samples)
    }
    cmd = append(cmd, args...)
    return cmd
}
