package cmd 

import (
    "github.com/ThatMatin/CimorghVCF/app"

    "github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
    Use: "create",
    Short: "Create a vcf dataset",
    Run: createDataset,
}

func createDataset(cmd *cobra.Command, args []string) {
    validateFlags()
    commands := prepCreate(args)
    App := app.NewApp(
        app.TILEDB_CLI_IMAGE_REF,
        inputDir,
        outputDir,
        datasetURI,
    )

    App.PullImage()
    App.CreateContainerWithCommand("cimorghCont",commands)
    App.StartContainer()
    App.ContainerLogsToStdout()
    App.RemoveContainer()
}

func prepCreate(args []string) []string {
    cmd := make([]string, 3)
    cmd[0] = "create"
    cmd[1] = "--uri"
    cmd[2] = datasetURI
    cmd = append(cmd, args...)
    return cmd
}
