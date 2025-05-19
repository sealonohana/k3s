package version

import "strings"

var (
	Program      = "k3s"
	ProgramUpper = strings.ToUpper(Program)
	Version      = "1.32.3-sp1"
	GitCommit    = "HEAD"

	UpstreamGolang = ""
)
