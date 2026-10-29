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

func DefaultSTLSteps(templateVariables any, options *STLOptions) []bootstrap.InitStep[STLOptions] {
	return []bootstrap.InitStep[STLOptions]{
		bootstrap.StepInitOutputDirectory[STLOptions](),
		bootstrap.StepRenderTemplates[STLOptions]([]bootstrap.FileToGenerate{
			{Template: "templates/README.tmpl", Output: "README.md"},
		},
			&templatesFS,
			&templateVariables,
		),
		StepOpenNemo(options),
	}
}

func StepOpenNemo(options *STLOptions) bootstrap.InitStep[STLOptions] {
	return bootstrap.InitStep[STLOptions]{
		Name: "Open NEMO",
		Run: func(ctx *bootstrap.StepExecutionContext[STLOptions]) error {
			if !ctx.Options.Nemo {
				return nil
			}

			nemoBinary := "nemo"
			_, err := exec.LookPath(nemoBinary)
			if err != nil {
				return fmt.Errorf("binary %q not found in path", nemoBinary)
			}

			return bootstrap.ExecuteCommand(ctx.OutputDir, nemoBinary, []string{"--existing-window", ctx.OutputDir}...)
		},
	}
}

//go:generate go run github.com/SR-G/argweave/cmd/weavedoc -type=STLOptions -file=main.go -edit=../README.md -marker=weavedoc-bootstrap-stl -fields=long,short,type,required,default,description
//go:generate go run github.com/SR-G/argweave/cmd/weavedoc -type=BaseOptions -file=/home/users/sergio/go/pkg/mod/github.com/!s!r-!g/bootstrap-library@v0.1.1/boostrapper/options.go -edit=../README.md -marker=weavedoc-bootstrap-base-options -fields=long,short,type,required,default,description

type STLOptions struct {
	bootstrap.BaseOptions

	URL  string `arg:"group=STL,long=url,required,env=BOOTSTRAPPER_URL,description=URL of the considered model"` // e.g. printables.com/ ...
	Nemo bool   `arg:"group=STL,long=nemo,env=BOOTSTRAPPER_NEMO,default=true,description=Should the opening of a nemo explorer window be triggered"`
}

func main() {
	bootstrapper := &bootstrap.Bootstrapper[STLOptions]{}
	cr, err := bootstrapper.Init()
	if err != nil {
		bootstrapper.Context.Logger.Error().Err(err).Msg("Fail at start time")
		os.Exit(cr)
	}
	if cr == bootstrap.INIT_HELP_REQUESTED {
		os.Exit(0)
	} else {
		templateVariables := TemplateVariables{
			ProjectName: bootstrapper.Opts.ProjectName,
			URL:         bootstrapper.Opts.URL,
		}
		os.Exit(bootstrapper.Generate(DefaultSTLSteps(templateVariables, bootstrapper.Opts)))
	}
}
