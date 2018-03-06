package flaggy

// ArgumentParser represents the set of vars and subcommands we are expecting
// from our input args.
type ArgumentParser struct {
	SubCommand
}

// NewArgumentParser creates a new ArgumentParser ready to parse inputs
func NewArgumentParser() *ArgumentParser {
	return &ArgumentParser{}
}

// Parse calculates all flags and subcommands
func (ap *ArgumentParser) Parse() {
	for _, command := range ap.SubCommands {
		command.parse(1) // initial depth of parsing is one command deep
	}
}
