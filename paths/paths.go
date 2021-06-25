// install package handles installing tooling for CI/CD work such as linters, apps.
// For example, install golang-lint tools, or other tools as needed.
// This might include docker image pulls as well.
package paths

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goyek/goyek"
)

// BuildRoot reflects the directory above `build` which should be the project directory. This variable provides more predictable path handling once set for all subsequent tasks.
var BuildRoot string

// ArtifactDirectory for all the downloaded and generated artifacts
var ArtifactDirectory string

// ToolsDirectory is for binaries downloaded as part of CI work but not build from source in project
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
	BuildRoot, err = filepath.Abs(parentDirectory)
	if err != nil {
		fmt.Printf("BuildRoot: [%v]", err)
	}
	BuildRoot, err = filepath.Abs(parentDirectory)
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
}

// TaskGetBuildRoot navigates up from `build` directory to ensure the path for the project is globally available for tasks with a simple call.
func TaskGetBuildRoot() goyek.Task {
	return goyek.Task{
		Name: "get-build-root",
		Command: func(tf *goyek.TF) {
			wd, err := os.Getwd()
			if err != nil {
				tf.Errorf("getwd: [%v]", err)
			}
			// WITH HELPER: BuildRoot = resolveParentDirectory(wd)
			projectDirectory := filepath.Join("../", wd)
			BuildRoot, err := filepath.Abs(projectDirectory)
			if err != nil {
				tf.Errorf("filepath.Abs(ProjectDirectory): [%v]", err)
			}
			tf.Logf("BuildRoot: [%s]", BuildRoot)
		},
	}
}

// ResolveParentDirectory returns the directory above the provided directory as a fully qualified absolute path
func resolveParentDirectory(tf *goyek.TF, childDirectory string) (parentDirectory string) {
	projectDirectory := filepath.Join("../", childDirectory)
	parentDirectory, err := filepath.Abs(projectDirectory)
	if err != nil {
		tf.Errorf("filepath.Abs(ProjectDirectory): [%v]", err)
	}
	tf.Logf("childDirectory [%s] --> parentDirectory: [%s]", childDirectory, parentDirectory)
	return parentDirectory
}

// resolveABSPath returns absolute path of any path, and logs error upon failure
func resolveABSPath(tf *goyek.TF, directory string) (ABSPath string) {
	ABSPath, err := filepath.Abs(directory)
	if err != nil {
		tf.Errorf("ABSPath: [%v]", err)
	}
	tf.Logf("directory [%s] --> ABSPath: [%s]", directory, ABSPath)
	return ABSPath
}
