package application

type applicationOptions struct {
	// program name
	ProgramName string
	// print help required
	Help bool
	// path to config file
	Config string
	// unrecognized arguments
	UnrecognizedArgs []string
	// enable dry run
	DryRun bool
}

func (opts *applicationOptions) addUnrecognizedArgument(arg string) {
	opts.UnrecognizedArgs = append(opts.UnrecognizedArgs, arg)
}

func (opts *applicationOptions) hasUnrecognizedArgs() bool {
	return len(opts.UnrecognizedArgs) > 0
}

func (opts *applicationOptions) showHelp() {
	opts.Help = true
}

func (opts *applicationOptions) enableDryRun() {
	opts.DryRun = true
}

func parseOptions(args []string) *applicationOptions {
	opts := &applicationOptions{}

	argsLen := len(args)

	if argsLen == 1 {
		opts.showHelp()
	}

	opts.ProgramName = args[0]

	for i := 1; i < argsLen; i++ {
		v := args[i]

		switch v {
		case "-h", "--help":
			opts.showHelp()

		case "-c", "--config":
			if i+1 < argsLen {
				opts.Config = args[i+1]
				i += 1
			} else {
				opts.showHelp()
			}

		case "--dry-run":
			opts.enableDryRun()

		default:
			opts.addUnrecognizedArgument(v)
			opts.showHelp()
		}

	}

	return opts
}
