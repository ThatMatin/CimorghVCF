package cmd

import (
    "os"

    "github.com/spf13/cobra"
)

var (
    inputDir string
    outputDir string
    datasetURI string
)

var rootCmd = &cobra.Command{
    Use: "cimorghvcf",
    Short: "CimorghVCF is a tool for storing and working on vcf files",
}

func init() {
    rootCmd.PersistentFlags().StringVarP(&inputDir, "input", "i", ".", "The root path of input files")
    rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", ".", "The path for output files")
    rootCmd.PersistentFlags().StringVarP(&datasetURI, "uri", "u", "", "specifies name for the dataset")
    rootCmd.AddCommand(createCmd)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
