// install package handles installing tooling for CI/CD work such as linters, apps.
// For example, install golang-lint tools, or other tools as needed.
// This might include docker image pulls as well.
package paths

import (
	"fmt"
	"os"
	"path/filepath"
)

// BuildRoot reflects the directory above `build` which should be the project directory. This variable provides more predictable path handling once set for all subsequent tasks.
var BuildRoot string

// ArtifactDirectory is the build output directory for all binaries and other files. This simplifies project management instead of having possible binaries in each directory go files are built from.
var ArtifactDirectory = "artifacts"

// BuildDir is the current build directory for the goyek files
var BuildDir = "build"

// ToolsDirectory contains local CI binaries for tooling that shouldn't get committed in git
var ToolsDirectory string

// InitBuildPathVariables sets the global variables for build tooling and artifacts
func InitBuildPathVariables() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("getwd: [%v]", err)
	}
	projectDirectory := filepath.Dir(wd)
	parentDirectory, err := filepath.Abs(projectDirectory)
	if err != nil {
		fmt.Printf("filepath.Abs(ProjectDirectory): [%v]\n", err)
	}
	BuildRootRelative, err := os.Getwd()
	if err != nil {
		fmt.Printf("getwd(): [%v]\n", err)
	}
	BuildRoot, err = filepath.Abs(BuildRootRelative)
	if err != nil {
		fmt.Printf("BuildRoot: [%v]", err)
	}
	ArtifactDirectory, err = filepath.Abs(filepath.Join(parentDirectory, "artifacts"))
	if err != nil {
		fmt.Printf("ArtifactDirectory: [%v]", err)
	}
	ToolsDirectory, err = filepath.Abs(filepath.Join(parentDirectory, "tools"))
	if err != nil {
		fmt.Printf("ToolsDirectory: [%v]", err)
	}
	fmt.Printf(`=== VARIABLES ===
variables

BuildRoot         : [%s]
ArtifactDirectory : [%s]
ToolsDirectory    : [%s]
`, BuildRoot, ArtifactDirectory, ToolsDirectory)
}
