// +build tools

package tools

// Manage tool dependencies via go.mod.
//
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// https://github.com/golang/go/issues/25922
import (
	_ "github.com/aevea/commitsar"
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/evilmartians/lefthook"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/quasilyte/go-ruleguard/dsl"
	_ "github.com/securego/gosec/v2/cmd/gosec"
	_ "github.com/zricethezav/gitleaks/v7"
	_ "mvdan.cc/gofumpt"
)
