package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goyek/goyek"
)

// BuildRoot reflects the directory above `build` which should be the project directory. This variable provides more predictable path handling once set for all subsequent tasks.
var BuildRoot string

// ArtifactDirectory is the build output directory for all binaries and other files. This simplifies project management instead of having possible binaries in each directory go files are built from.
var ArtifactDirectory = "artifacts"

// BuildDir is the current build directory for the goyek files
var BuildDir = "build"

// ToolsDirectory contains local CI binaries for tooling that shouldn't get committed in git
var ToolsDirectory string

func main() {
	flow().Main()
}

func flow() *goyek.Taskflow {
	InitBuildPathVariables()
	flow := &goyek.Taskflow{}

	// parameters
	ci := flow.RegisterBoolParam(goyek.BoolParam{
		Name:  "ci",
		Usage: "Whether CI is calling the build script",
	})
	fmt.Printf("ci: [%v]", ci)
	// // tasks
	// clean := flow.Register(taskClean())
	// _ = flow.Register(taskPrecommitInit())
	// _ = flow.Register(taskPrecommitRun())
	// lefthookinit := flow.Register(taskLeftHookInit())
	// lefthookrun := flow.Register(taskLeftHookRun())

	// build := flow.Register(taskBuild())
	// fmt := flow.Register(taskFmt())

	// markdownlint := flow.Register(taskMarkdownLint())
	// misspell := flow.Register(taskMisspell())
	// golangciLint := flow.Register(taskGolangciLint())
	// test := flow.Register(taskTest())
	// modTidy := flow.Register(taskModTidy())
	// diff := flow.Register(taskDiff(ci))
	// init := flow.Register(TaskInit())
	initDir := flow.Register(TaskInitDir())
	// _ = flow.Register(taskComposeUp())
	// _ = flow.Register(taskComposeDestroy())

	// docker tasks
	// _ = flow.Register(taskDockerBuild())

	// setup and initialization tasks for a brand new environment
	_ = flow.Register(taskInit(goyek.Deps{
		// clean,
		// modTidy,
		initDir,
		// precommitinit,
		// lefthookinit,
	}))
	// pipelines
	// lint := flow.Register(taskLint(goyek.Deps{
	// 	// misspell,
	// 	// markdownlint,
	// 	// golangciLint,
	// 	// precommitrun,
	// 	lefthookrun,
	// }))
	// _ = flow.Register(taskAll(goyek.Deps{
	// 	initDir,
	// 	// init,
	// 	clean,
	// 	build,
	// 	fmt,
	// 	lint,
	// 	test,
	// 	modTidy,
	// 	diff,
	// }))
	return flow
}

func taskInit(deps goyek.Deps) goyek.Task {
	return goyek.Task{
		Name:  "init",
		Usage: "initialize all developer tooling and project tools",
		Deps:  deps,
	}
}

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
		fmt.Printf("BuildRoot: [%v]\n", err)
	}
	BuildDir, err = filepath.Abs(filepath.Join(BuildRoot, "build"))
	if err != nil {
		fmt.Printf("BuildDir: [%v]\n", err)
	}
	ArtifactDirectory, err = filepath.Abs(filepath.Join(parentDirectory, "artifacts"))
	if err != nil {
		fmt.Printf("ArtifactDirectory: [%v]\n", err)
	}
	ToolsDirectory, err = filepath.Abs(filepath.Join(parentDirectory, "tools"))
	if err != nil {
		fmt.Printf("ToolsDirectory: [%v]\n", err)
	}
	fmt.Printf(`=== VARIABLES ===
variables

BuildRoot         : [%s]
ArtifactDirectory : [%s]
ToolsDirectory    : [%s]
`, BuildRoot, ArtifactDirectory, ToolsDirectory)
}

// TaskInitDir creates the project directories and sets permissions to use them
func TaskInitDir() goyek.Task {
	return goyek.Task{
		Name:  "initdir",
		Usage: "create project directories",
		Action: func(tf *goyek.TF) {
			if err := os.Chdir(BuildRoot); err != nil {
				tf.Error("Unable to chdir to BuildRoot")
			}
			for _, i := range []string{ArtifactDirectory, ToolsDirectory} {
				if err := os.Mkdir(i, 0o700); err != nil {
					if strings.Contains(err.Error(), "file exists") {
						tf.Logf("🔃 [%s] dir already exists\n", i)
					} else {
						tf.Errorf("failed to mkdir: [%s] with error: %v", i, err)
					}
					continue
				}
				tf.Logf("✅ [%s] dir created\n", i)
			}
		},
	}
}
