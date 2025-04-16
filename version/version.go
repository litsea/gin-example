package version

import (
	"fmt"
	"runtime"
)

// Populated during build, don't touch!
var (
	Version   = "v0.0.0"
	GitRev    = "undefined"
	GitBranch = "undefined"
	BuildDate = "Mon, 02 Jan 2006 15:04:05 -0700"
	// GoVersion system go version.
	GoVersion = runtime.Version()
	// Platform info.
	Platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)
