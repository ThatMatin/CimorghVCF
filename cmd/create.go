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
    App := app.NewApp(app.TILEDB_CLI_IMAGE_REF)
    App.PullImage()
    contCmd := []string{"create", "--help"}
    App.CreateContainerWithCommand("cimorghCont",contCmd)
    App.StartContainer(false)
    App.ContainerLogsToStdout()
    App.RemoveContainer()
}



