package application

type CliApp struct {
	opts *applicationOptions
	sys  systemInteractions
}

func NewCliApp(args []string) *CliApp {
	return newCliApp(args, &realSystemInteractions{})
}

func newCliApp(args []string, sys systemInteractions) *CliApp {
	return &CliApp{
		opts: parseOptions(args),
		sys:  sys,
	}
}

func (app *CliApp) Run() {
	if app.opts.hasUnrecognizedArgs() {
		app.sys.Printf("Got unrecognized args: %s\n", app.opts.UnrecognizedArgs)
		app.printHelp()
		return
	}

	if app.opts.Help {
		app.printHelp()
		return
	}

	// todo: config an
	app.sys.Printf("Config file: %s\n", app.opts.Config)
}

func (app *CliApp) printHelp() {
	app.sys.Printf("Program: %s\n", app.opts.ProgramName)
	app.sys.Printf("Usage: \n")
}
