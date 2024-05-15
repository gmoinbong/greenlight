package vcs

import (
	"fmt"
	"runtime/debug"
)

func Version() string {
	var revision string
	var modifed bool

	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range bi.Settings {
			switch s.Key {
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
		return fmt.Sprintf("%s-dirty", revision)
	}
	return revision
}
