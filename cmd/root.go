package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ThatMatin/CimorghVCF/app"
)

var (
    outputDir string
    inputDir string
    datasetURI string
    debug bool
    App *app.App
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
    rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", true, "show debug logs")
    rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "~", "The path for output files")
    rootCmd.PersistentFlags().StringVarP(&datasetURI, "uri", "u", "", "specifies name for the dataset")
    rootCmd.PersistentFlags().StringVarP(&inputDir, "input", "i", "~", "The root path of input files")
    rootCmd.MarkPersistentFlagRequired("input")
    rootCmd.MarkPersistentFlagRequired("output")
    rootCmd.MarkPersistentFlagRequired("uri")
    rootCmd.AddCommand(createCmd)
    rootCmd.AddCommand(ingestCmd)
    rootCmd.AddCommand(samplesCmd)
    rootCmd.AddCommand(statCmd)
    rootCmd.AddCommand(exportCmd)
}

func Execute() {
    if debug {
        initLogs(logrus.DebugLevel)
    } else {
        initLogs(logrus.ErrorLevel)
    }
    go handleInterrups()
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func initLogs(level logrus.Level) {
    logrus.SetOutput(os.Stderr)
    logrus.SetLevel(level)
}

func containerCycle(args []string, fn func([]string) []string) {
    validateFlags()
    command := fn(args)
    App = app.NewApp(
        app.TILEDB_CLI_IMAGE_REF,
        inputDir,
        outputDir,
        datasetURI,
    )
    App.PullImage()
    App.CreateContainerWithCommand("cimorghCont",command)
    App.StartContainer()
    App.ContainerLogsToStdout()
    App.RemoveContainer()
    App.ShutdownApp()
}

func handleInterrups() {
    c := make(chan os.Signal, 1)
    signal.Notify(
        c, os.Interrupt,
        syscall.SIGTERM,
        syscall.SIGINT,
    )
    <-c
    logrus.Info("Forced to stop. Cleaning up...")
    App.StopContainer()
    App.RemoveContainer()
    App.ShutdownApp()
}
