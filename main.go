package main

import (
    "context"
    "fmt"
    "io"
    "os"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
    "github.com/docker/docker/pkg/stdcopy"
)


const (
    IMAGE_TAG = "latest"
    TILEDB_CLI_IMAGE_REF = "tiledb/tiledbvcf-cli" + ":" + IMAGE_TAG
)

func main() {
    ctx := context.Background()

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    reader, err := cli.ImagePull(ctx, TILEDB_CLI_IMAGE_REF, types.ImagePullOptions{})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, reader)

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: TILEDB_CLI_IMAGE_REF,
        Cmd: []string{"version"},
    }, nil,nil,nil, "")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Created container ID %s\n", resp.ID)
    
    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }

    statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
    select {
    case err := <- errCh:
        if err != nil {
            panic(err)
        }
    case <- statusCh:
    }

    out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
    if err != nil {
        panic(err)
    }

    stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}
