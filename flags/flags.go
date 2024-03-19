package flags

import (
	"github.com/spf13/pflag"
)

var Repo *string = pflag.StringP("repo", "r", "~/.dotfiles", "Path to the git repository")
var Keep *int = pflag.IntP("keep", "k", 10, "Keep last generations")
var Update *bool = pflag.BoolP("update", "u", false, "Update the system")
var Diff *bool = pflag.BoolP("diff", "d", false, "Show the diff of the last generation")
