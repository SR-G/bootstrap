package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"

	bootstrap "github.com/SR-G/bootstrap-library/boostrapper"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

type TemplateVariables struct {
	ProjectName string
	URL         string
}

func DefaultSTLSteps(templateVariables any, options *STLOptions) []bootstrap.InitStep {
	return []bootstrap.InitStep{
		bootstrap.StepInitOutputDirectory(),
		bootstrap.StepRenderTemplates([]bootstrap.FileToGenerate{
			{Template: "templates/README.tmpl", Output: "README.md"},
		},
			&templatesFS,
			&templateVariables,
		),
		StepOpenNemo(),
	}
}

func StepOpenNemo() bootstrap.InitStep {
	return bootstrap.InitStep{
		Name: "Open NEMO",
		Run: func(ctx *bootstrap.StepExecutionContext) error {
			nemoBinary := "nemo"
			_, err := exec.LookPath(nemoBinary)
			if err != nil {
				return fmt.Errorf("binary %q not found in path", nemoBinary)
			}

			return bootstrap.ExecuteCommand(ctx.OutputDir, nemoBinary, []string{"--existing-window", ctx.OutputDir}...)
		},
	}
}

//go:generate go run github.com/SR-G/argweave/cmd/weavedoc -type=STLOptions -file=main.go -edit=../README.md -marker=weavedoc-bootstrap-stl -fields=long,short,type,env,required,default,description
//go:generate go run github.com/SR-G/argweave/cmd/weavedoc -type=BaseOptions -file=/home/users/sergio/go/pkg/mod/github.com/!s!r-!g/bootstrap-library@v0.1.1/boostrapper/options.go -edit=../README.md -marker=weavedoc-bootstrap-base-options -fields=long,short,type,env,required,default,description

type STLOptions struct {
	*bootstrap.BaseOptions

	URL string `arg:"long=url,required,env=BOOTSTRAPPER_URL"` // e.g. printables.com/ ...
}

func main() {
	bootstrapper := &bootstrap.Bootstrapper[STLOptions]{}
	if cr, err := bootstrapper.Init(); err != nil {
		bootstrapper.Context.Logger.Error().Err(err).Msg("Fail at start time")
		os.Exit(cr)
	}
	templateVariables := TemplateVariables{
		ProjectName: bootstrapper.Opts.ProjectName,
		URL:         bootstrapper.Opts.URL,
	}
	os.Exit(bootstrapper.Generate(DefaultSTLSteps(templateVariables, bootstrapper.Opts)))
}
