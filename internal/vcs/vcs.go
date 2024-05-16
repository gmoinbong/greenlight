package vcs

import (
	"fmt"
	"runtime/debug"
)

func Version() string {
	var (
		time     string
		revision string
		modifed  bool
	)
	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.time":
				time = s.Value
			case "vcs.revision":
				revision = s.Value
			case "vsc.modified":
				if s.Value == "true" {
					modifed = true
				}
			}
		}
	}

	if modifed {
		return fmt.Sprintf("%s-%s-dirty", time, revision)
	}
	return fmt.Sprintf("%s-%s", time, revision)
}
