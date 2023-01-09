package app

import (
    "fmt"
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/sirupsen/logrus"
)

const ASCII_ART = `
 ________  ___  _____ ______   ________  ________  ________  ___  ___  ___      ___ ________  ________     
|\   ____\|\  \|\   _ \  _   \|\   __  \|\   __  \|\   ____\|\  \|\  \|\  \    /  /|\   ____\|\  _____\    
\ \  \___|\ \  \ \  \\\__\ \  \ \  \|\  \ \  \|\  \ \  \___|\ \  \\\  \ \  \  /  / | \  \___|\ \  \__/     
 \ \  \    \ \  \ \  \\|__| \  \ \  \\\  \ \   _  _\ \  \  __\ \   __  \ \  \/  / / \ \  \    \ \   __\    
  \ \  \____\ \  \ \  \    \ \  \ \  \\\  \ \  \\  \\ \  \|\  \ \  \ \  \ \    / /   \ \  \____\ \  \_|    
   \ \_______\ \__\ \__\    \ \__\ \_______\ \__\\ _\\ \_______\ \__\ \__\ \__/ /     \ \_______\ \__\     
    \|_______|\|__|\|__|     \|__|\|_______|\|__|\|__|\|_______|\|__|\|__|\|__|/       \|_______|\|__|     
                                                                                                           
`

const (
    IMAGE_TAG = "latest"
    TILEDB_CLI_IMAGE_REF = "tiledb/tiledbvcf-cli" + ":" + IMAGE_TAG
)

type App struct {
    cli *client.Client
    containerID string
    inputDir string
    outputDir string
    datasetURI string
    ImageRef string
}

// defer cli.Close()
// TODO: remember to close client
func NewApp(imageRef ,inputDir, outputDir, datasetURI string) *App {
    fmt.Print(ASCII_ART)
    initLogs()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        logrus.Fatalf("error creating client: %v\n", err)
    }
    app := &App{
        cli: cli,
        ImageRef: imageRef,
        inputDir: inputDir,
        outputDir: outputDir,
        datasetURI: datasetURI,
    }

    return app
}

func (a *App) StartContainer() error {
    ctx := context.Background()
    if err := a.cli.ContainerStart(ctx, a.containerID, types.ContainerStartOptions{}); err != nil {
        logrus.Errorf("error starting container: %v\n", err)
        return err
    }

    statusCh, errCh := a.cli.ContainerWait(ctx, a.containerID, container.WaitConditionNotRunning)
    select {
    case err := <- errCh:
        if err != nil {
            logrus.Errorf("error in running container: %v\n", err)
            return err
        }
    case resp := <- statusCh:
        logrus.Infof("container returned with status code: %d\n", resp.StatusCode)
    }

    return nil
}

func (a *App) RemoveContainer() error {
    ctx := context.Background()
    if err := a.cli.ContainerRemove(ctx, a.containerID, types.ContainerRemoveOptions{}); err != nil {
        logrus.Warnf("couldn't remove container: %v", err)
        return err
    }
    logrus.Infof("removed container with ID: %v",a.containerID)
    return nil
}

func (a *App) ContainerLogsToStdout() error {
    ctx := context.Background()
    out, err := a.cli.ContainerLogs(ctx, a.containerID, types.ContainerLogsOptions{ShowStdout: true})
    if err != nil {
        logrus.Error(err)
        return err
    }
    defer out.Close()

    if _,err := stdcopy.StdCopy(os.Stdout, os.Stderr, out); err != nil {
        logrus.Error(err)
        return err
    }

    return nil
}

func (a *App) CreateContainerWithCommand(name string, commands []string) error {
    ctx := context.Background()
    hostConfig := &container.HostConfig{
        Mounts: []mount.Mount{
            {
                Type: mount.TypeBind,
                Source: a.inputDir,
                Target: a.inputDir,
            },
            {
                Type: mount.TypeBind,
                Source: a.outputDir,
                Target: a.outputDir,
            },
        },
    }
    logrus.Debug(commands)

    containerConfig := &container.Config{
        Image: a.ImageRef,
        Cmd: commands,
        WorkingDir: a.outputDir,
        AttachStdout: true,
        AttachStderr: true,
    }

    resp, err := a.cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, name)
    if err != nil {
        logrus.Errorf("error creating container: %v\n", err)
        return err
    }
    logrus.Infof("Created container: %s\n", resp.ID)
    a.containerID = resp.ID

    return nil
} 

func (a *App) PullImage() error {
    ctx := context.Background()
    if !a.CheckImageExists(a.ImageRef) {
        reader, err := a.cli.ImagePull(ctx, a.ImageRef, types.ImagePullOptions{})
        if err != nil {
            logrus.Errorf("error in pulling images: %v\n", err)
            return err
        }
        defer reader.Close()

        if _, err := io.Copy(os.Stdout, reader); err !=nil {
            logrus.Error(err)
            return err
        }
    }
    return nil
}

func (a *App) CheckImageExists(imageTag string) bool {
    ctx := context.Background()
    imgList, err := a.cli.ImageList(ctx, types.ImageListOptions{})
    if err != nil {
        logrus.Fatalf("error in retrieving list of images: %v\n", err)
    }
    for _, image := range imgList {
        if image.RepoTags[0] == imageTag {
            logrus.Info("Image exists locally")

            return true
        }
    }

    return false
}

func (a *App) ShutdownApp() error {
    return a.cli.Close()
}

func initLogs() {
    logrus.SetOutput(os.Stdout)
    logrus.SetLevel(logrus.DebugLevel)
}
