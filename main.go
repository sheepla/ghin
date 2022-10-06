package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/sheepla/ghin/gh"
	"github.com/sheepla/ghin/ui"
	cli "github.com/urfave/cli/v2"
)

//nolint:gochecknoglobals
var (
	appName        = "ghin"
	appVersion     = "unknown"
	appUsage       = "A GitHub releases installer"
	appDescription = "A GitHub releases installer with APT-like sub commands and fzf-like interactive UI"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErrArgs
	exitCodeErrSelectRepo
	exitCodeErrSelectAsset
	exitCodeErrDownload
	exitCodeErrURL
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
		Usage:       appUsage,
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
	query := strings.Join(ctx.Args().Slice(), " ")
	param := gh.NewSearchParam(query)

	repo, err := selectRepo(param)
	if err != nil {
		return cli.Exit(err, exitCodeErrSelectRepo.Int())
	}

	ast, err := selectAsset(repo)
	if err != nil {
		return cli.Exit(err, exitCodeErrSelectAsset.Int())
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dlurl, err := url.Parse(ast.DownloadURL)
	if err != nil {
		return cli.Exit(
			fmt.Errorf("invalid download URL: %s, %w", ast.DownloadURL, err),
			exitCodeErrURL.Int(),
		)
	}

	if err := download(dlurl, cwd); err != nil {
		return cli.Exit(err, exitCodeErrDownload.Int())
	}

	return nil
}

func runRemoveCommand(ctx *cli.Context) error {
	return nil
}

func runListCommand(ctx *cli.Context) error {
	return nil
}

func selectRepo(param *gh.SearchParam) (*gh.Repo, error) {
	repos, err := gh.Search(param)
	if err != nil {
		return nil, fmt.Errorf("failed to search repositories: %w", err)
	}

	repo, err := ui.SelectRepo(*repos)
	if err != nil {
		return nil, fmt.Errorf("failed to select a repository: %w", err)
	}

	return repo, nil
}

func selectAsset(repo *gh.Repo) (*gh.Asset, error) {
	rels, err := gh.GetReleases(repo.Owner, repo.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get releases: %w", err)
	}

	rel, err := ui.SelectTag(*rels)
	if err != nil {
		return nil, fmt.Errorf("failed to select a tag: %w", err)
	}

	ast, err := ui.SelectAsset(rel.Assets)
	if err != nil {
		return nil, fmt.Errorf("falied to select an asset: %w", err)
	}

	return ast, nil
}

func download(u *url.URL, destDir string) error {
	//nolint:noctx
	res, err := http.Get(u.String())
	if err != nil {
		return fmt.Errorf("failed to fetch the file: %w", err)
	}

	defer res.Body.Close()

	fname := path.Base(u.String())
	fpath := path.Join(destDir, fname)

	file, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("failed to create the file: %w", err)
	}

	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		return fmt.Errorf("failed to write content into file %s: %w", fpath, err)
	}

	return nil
}
