// install package handles installing tooling for CI/CD work such as linters, apps.
// For example, install golang-lint tools, or other tools as needed.
// This might include docker image pulls as well.
package install

import (
	// "context"
	// "reflect"
	"context"
	"testing"

	"github.com/goyek/goyek"
	"github.com/matryer/is"
)

func TestTaskInstallLintingTools(t *testing.T) {
    var ctx context.Context = nil

	is := is.New(t)
	flow := &goyek.Taskflow{}
	_ = flow.Register(TaskInstallLintingTools())
	// ctx := context.Context
	response := flow.Run(ctx ,"install-linters")
	is.Equal(response, 0) // install-linters exits without error
}
