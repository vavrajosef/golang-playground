package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
        Commands: []*cli.Command{
            {
                Name:    "cancel",
                Aliases: []string{"c"},
                Usage:   "cancel context",
                Action: func(cCtx *cli.Context) error {
                    return cancel(cCtx.Args().First())
                },
            },
            {
                Name:    "cancelcause",
                Aliases: []string{"cc"},
                Usage:   "cancel cause context",
                Action: func(cCtx *cli.Context) error {
                    return cancelcause(cCtx.Args().First())
                },
            },
            {
                Name:    "background",
                Aliases: []string{"b"},
                Usage:   "this doesn't really do anything",
                Action: func(cCtx *cli.Context) error {
                    return background(cCtx.Args().First())
                },
            },
            {
                Name:    "deadline",
                Aliases: []string{"d"},
                Usage:   "context with deadline",
                Action: func(cCtx *cli.Context) error {
                    return deadline(cCtx.Args().First())
                },
            },
            {
                Name:    "timeout",
                Aliases: []string{"d"},
                Usage:   "context with timeout",
                Action: func(cCtx *cli.Context) error {
                    return timeout(cCtx.Args().First())
                },
            },
        },
    }

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}

func cancel(arg string) error {
    return errors.New("not implemented")
}

func timeout(arg string) error {
    timeoutMilis, err := convertArgToInt(arg)
    if err != nil {
        return err
    }
    ctx, cancel := context.WithTimeout(createBackgroundContext(), time.Duration(timeoutMilis) * time.Millisecond)
    defer cancel()


    input := receiveInput()
    select {
    case <-ctx.Done():
        // timeout contextu vypršel
        fmt.Println("Operation timedout.")
    default:
        fmt.Println("You were faster than timeout and wrote: " + input)
    }
    return nil
}

func receiveInput() string {
    fmt.Println("Enter whatever text before timeout occurs: ")
    var textBeforeTimeout string
    fmt.Scanln(&textBeforeTimeout)
    return textBeforeTimeout
}

func deadline(arg string) error {
    return errors.New("not implemented")
}

func background(arg string) error {
    return errors.New("not implemented")
}

func cancelcause(arg string) error {
    return errors.New("not implemented")
}

func createBackgroundContext() context.Context {
    return context.Background()
}

func convertArgToInt(arg string) (int, error) {
    if len(arg) == 0 {
        return 0, errors.New("argument cannot be empty")
    }
    return strconv.Atoi(arg)
}