package cmd

import (
    "os"

    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

var (
    outputDir string
    inputDir string
    datasetURI string
)

var rootCmd = &cobra.Command{
    Use: "cimorghvcf",
    Short: "CimorghVCF is a tool for storing and working on vcf files",
}

func validateFlags() {
    if inputDir == outputDir {
        logrus.Fatalf("input and output paths must differ")
    }
    if datasetURI == "" {
        logrus.Fatalf("no database uri specified")
    }
}

func init() {
    rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "~", "The path for output files")
    rootCmd.PersistentFlags().StringVarP(&datasetURI, "uri", "u", "", "specifies name for the dataset")
    rootCmd.PersistentFlags().StringVarP(&inputDir, "input", "i", "~", "The root path of input files")
    rootCmd.MarkPersistentFlagRequired("input")
    rootCmd.MarkPersistentFlagRequired("output")
    rootCmd.MarkPersistentFlagRequired("uri")
    rootCmd.AddCommand(createCmd)
    rootCmd.AddCommand(ingestCmd)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
