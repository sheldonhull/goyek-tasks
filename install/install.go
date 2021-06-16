// install package handles installing tooling for CI/CD work such as linters, apps.
// For example, install golang-lint tools, or other tools as needed.
// This might include docker image pulls as well.
package install

import (
	"time"

	"github.com/goyek/goyek"
	"github.com/pterm/pterm"
)

func TaskInstallLintingTools() goyek.Task {
	return goyek.Task{
		Name:  "install-linters",
		Usage: "Install various linting tools that pre-commit or other tooling will need",
		Command: func(tf *goyek.TF) {
			goToolsRepos := []string{
				"github.com/securego/gosec/v2/cmd/gosec@master",
				"golang.org/x/tools/cmd/goimports@master",
				"github.com/sqs/goreturns@master",
				"github.com/golangci/golangci-lint/cmd/golangci-lint@master",
				"github.com/go-critic/go-critic/cmd/gocritic@master",
			}
			totalToInstall := len(goToolsRepos)
			pterm.DefaultSection.Printfln("Installing linters for development: [%d to install]",totalToInstall)
			p, _ := pterm.DefaultProgressbar.WithTotal(totalToInstall).WithTitle("Installing stuff").Start()
			for _, i := range goToolsRepos {
				p.Title = "Installing " + i
				GetCommand := tf.Cmd("go", "install", i)
				if err := GetCommand.Run(); err != nil {
					pterm.Warning.Printf("Could not install [%s] per [%v]\n", i, err)
					// tf.Logf("Could not install [%s] per [%v]\n", i, err)
				}
				p.Increment()
				time.Sleep(time.Millisecond * 350)


			}
			p.Title = "linters installed successfully"
			_, _ = p.Stop()
		},
	}
}
