package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	cli "github.com/urfave/cli/v2"
)

//nolint:gochecknoglobals
var (
	appName        = "ghin"
	appVersion     = "unknown"
	appDescription = "A GitHub releases installer"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErrArgs
)

func (code exitCode) Int() int {
	return int(code)
}

func main() {
	app := initApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func initApp() *cli.App {
	//nolint:exhaustivestruct,exhaustruct
	return &cli.App{
		Name:        appName,
		Version:     appVersion,
		Usage:       appDescription,
		Description: appDescription,
		Suggest:     true,
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() == 0 {
				return cli.Exit("must require arguments", exitCodeErrArgs.Int())
			}

			return cli.Exit("", exitCodeOK.Int())
		},
		Commands: []*cli.Command{
			{
				Name:      "install",
				Usage:     "Install releases",
				Aliases:   []string{"i"},
				ArgsUsage: "QUERY...",
				Action:    runInstallCommand,
			},
			{
				Name:      "remove",
				Usage:     "Uninstall binaries",
				Aliases:   []string{"r"},
				ArgsUsage: "QUERY...",
				Action:    runRemoveCommand,
			},
			{
				Name:      "list",
				Usage:     "List installed releases",
				Aliases:   []string{"l"},
				ArgsUsage: "[QUERY...]",
				Action:    runListCommand,
			},
		},
	}
}

func runInstallCommand(ctx *cli.Context) error {
	return nil
}

func runRemoveCommand(ctx *cli.Context) error {
	return nil
}

func runListCommand(ctx *cli.Context) error {
	return nil
}

//nolint:deadcode,unused
func download(url, destDir string) error {
	//nolint:gosec,noctx
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch the file: %w", err)
	}

	defer res.Body.Close()

	fname := path.Base(url)
	fpath := path.Join(destDir, fname)

	file, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("failed to create the file: %w", err)
	}

	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		return fmt.Errorf("failed to write content into file: %w", err)
	}

	return nil
}
