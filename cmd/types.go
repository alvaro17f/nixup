package cmd

import "github.com/alvaro17f/nixup/internal/utils"

type flagDetails[valueType any] struct {
	long        string
	short       string
	value       valueType
	description string
}

func (f *flagDetails[valueType]) getFlagDetails() (string, string, valueType, string) {
	return f.long, f.short, f.value, f.description
}

const defaultDiffValue = false

var diff = flagDetails[bool]{
	long:        "diff",
	short:       "d",
	value:       defaultDiffValue,
	description: "Show the diff of the last generation",
}

var defaultHostnameValue = utils.GetHostname()

var hostname = flagDetails[string]{
	long:        "hostname",
	short:       "n",
	value:       defaultHostnameValue,
	description: "Set the hostname",
}

const defaultKeepValue = 10

var keep = flagDetails[int]{
	long:        "keep",
	short:       "k",
	value:       defaultKeepValue,
	description: "Keep last generations",
}

const defaultRepoValue = "~./dotfiles"

var repo = flagDetails[string]{
	long:        "repo",
	short:       "r",
	value:       defaultRepoValue,
	description: "Path to the git repository",
}

const defaultUpdateValue = false

var update = flagDetails[bool]{
	long:        "update",
	short:       "u",
	value:       defaultUpdateValue,
	description: "Update the system",
}
