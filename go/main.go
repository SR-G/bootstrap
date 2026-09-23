package main

import (
	"embed"
	"os"
	"path/filepath"

	bootstrap "github.com/SR-G/bootstrap-library/boostrapper"
	"github.com/SR-G/sul"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

// templateData is exposed to every project template.
type TemplateVariables struct {
	ProjectName string
	ModulePath  string
}

// DefaultSteps returns the standard list of steps used to bootstrap a new project.
func DefaultGoSteps(templateVariables any, options *GoOptions) []bootstrap.InitStep {
	return []bootstrap.InitStep{
		bootstrap.StepInitOutputDirectory(),
		bootstrap.StepRenderTemplates(
			[]bootstrap.FileToGenerate{
				{Template: "templates/gomod.tmpl", Output: "go.mod"},
				{Template: "templates/gitignore.tmpl", Output: ".gitignore"},
				{Template: "templates/main.go.tmpl", Output: "main.go"},
				{Template: "templates/options.go.tmpl", Output: "options.go"},
				{Template: "templates/Makefile.tmpl", Output: "Makefile"},
				{Template: "templates/README.tmpl", Output: "README.md"},
			},
			&templatesFS,
			&templateVariables,
		),
		StepGoModInit(options),
	}
}

// GoModInitStep runs "go mod init <ModulePath>" in WorkDir. Does nothing if a go.mod already exists.
func StepGoModInit(options *GoOptions) bootstrap.InitStep {
	return bootstrap.InitStep{
		Name: "Initialize go.mod",
		Run: func(ctx *bootstrap.StepExecutionContext) error {
			if sul.FileExists(filepath.Join(ctx.OutputDir, "go.mod")) {
				return nil
			} else {
				return bootstrap.ExecuteCommand(ctx.OutputDir, "go", []string{"mod", "init", options.ModulePath}...)
			}
		},
	}
}

//go:generate go run github.com/SR-G/argweave/cmd/weavedoc -type=GoOptions -file=main.go -edit=../README.md -marker=weavedoc-bootstrap-go -fields=long,short,type,required,default,description
type GoOptions struct {
	*bootstrap.BaseOptions

	ModulePath string `arg:"long=module,required,env=BOOTSTRAPPER_MODULE_PATH"` // e.g. github.com/SR-G/my-new-app
}

func main() {
	bootstrapper := &bootstrap.Bootstrapper[GoOptions]{}
	if cr, err := bootstrapper.Init(); err != nil {
		bootstrapper.Context.Logger.Error().Err(err).Msg("Fail at start time")
		os.Exit(cr)
	}
	t := &TemplateVariables{
		ProjectName: bootstrapper.Opts.ProjectName,
		ModulePath:  bootstrapper.Opts.ModulePath,
	}
	os.Exit(bootstrapper.Generate(DefaultGoSteps(t, bootstrapper.Opts)))
}
