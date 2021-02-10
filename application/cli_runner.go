package application

import (
	"github.com/shults/mongomove/application/config"
)

type CliRunner struct {
	opts   *applicationOptions
	sys    systemInteractions
	reader config.Reader
}

func NewCliRunner(args []string) *CliRunner {
	return newCliApp(args, &realSystemInteractions{}, config.NewReader())
}

func newCliApp(args []string, sys systemInteractions, reader config.Reader) *CliRunner {
	return &CliRunner{
		opts:   parseOptions(args),
		sys:    sys,
		reader: reader,
	}
}

func (r *CliRunner) Run() {
	if r.opts.hasUnrecognizedArgs() {
		r.sys.Printf("Got unrecognized args: %s\n", r.opts.UnrecognizedArgs)
		r.printHelp()
		return
	}

	if r.opts.Help {
		r.printHelp()
		return
	}

	conf, err := r.reader.Read(r.opts.Config)

	if err != nil {
		r.printErrorAndExit(err, 1)
		return
	}

	app := NewApp(conf, r.opts)
	err = app.Run()

	if err != nil {
		r.printErrorAndExit(err, 2)
		return
	}
}

func (r *CliRunner) printErrorAndExit(err error, exitCode int) {
	r.sys.Printf("Got error: %s\n\n", err.Error())
	r.printHelp()
	r.sys.Exit(exitCode)
}

func (r *CliRunner) printHelp() {
	var usage = `Usage: %s [--help|-h] [-c|--config <path/to/config.yml>] [--dry-run]  
Options:
	--dry-run		runs all process without remove stage
	-c|--config 	path to config.yml
	-h|--help 		prints help message
`
	r.sys.Printf(usage, r.opts.ProgramName)
}
