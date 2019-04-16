package main

import (
	"fmt"
	"github.com/cloudfoundry/go-cnb/golang"
	"os"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/libcfbuildpack/detect"
)

func main() {
	context, err := detect.DefaultDetect()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create a default detection context: %s", err)
		os.Exit(101)
	}

	code, err := runDetect(context)
	if err != nil {
		context.Logger.Info(err.Error())
	}

	os.Exit(code)
}

func runDetect(context detect.Detect) (int, error) {
	return context.Pass(buildplan.BuildPlan{
		golang.Layer: buildplan.Dependency{
		Metadata: buildplan.Metadata{"build": true, "launch": true},
	}})
}