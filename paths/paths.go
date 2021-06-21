// install package handles installing tooling for CI/CD work such as linters, apps.
// For example, install golang-lint tools, or other tools as needed.
// This might include docker image pulls as well.
package paths

import (
	"os"
	"path/filepath"

	"github.com/goyek/goyek"
)

// BuildRoot reflects the directory above `build` which should be the project directory. This variable provides more predictable path handling once set for all subsequent tasks.
var BuildRoot string

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
