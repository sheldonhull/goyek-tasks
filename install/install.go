// install package handles installing tooling for CI/CD work such as linters, apps.
// For example, install golang-lint tools, or other tools as needed.
// This might include docker image pulls as well.
package install

import (
	"time"

	"github.com/goyek/goyek"
	// "github.com/pterm/pterm"
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
				"github.com/mgechev/revive@master",
			}
			totalToInstall := len(goToolsRepos) // + 1

			// pterm.DefaultSection.Printfln("Installing linters for development: [%d to install]", totalToInstall)
			tf.Logf("🔨 Installing linters for development: [%d to install]", totalToInstall)
			// p, err := pterm.DefaultProgressbar.
			// 	WithTotal(totalToInstall).
			// 	WithTitle("Installing stuff").
			// 	WithShowElapsedTime(true).
			// 	Start()
			// if err != nil {
			// 	pterm.Warning.Printfln("DefaultProgressbar [%v]", err)
			// }
			for _, i := range goToolsRepos {
				// p.Title = "Installing " + i
				tf.Logf("🔨 installing [%s]\n", i)
				GetCommand := tf.Cmd("go", "install", i)
				if err := GetCommand.Run(); err != nil {
					// pterm.Warning.Printf("Could not install [%s] per [%v]\n", i, err)
					tf.Errorf("❗ Could not install [%s] per [%v]\n", i, err)
				}
				time.Sleep(time.Millisecond * 350)
				// p.Increment()
			}
			tf.Log("✅ linters installed successfully\n")
			// p.Title = "linters installed successfully"
			// _, _ = p.Stop()
		},
	}
}
