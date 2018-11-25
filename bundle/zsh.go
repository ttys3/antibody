package bundle

import (
	"fmt"
	"github.com/getantibody/antibody/compatible"
	"os"
	"path/filepath"
	"strings"

	"github.com/getantibody/antibody/project"
)

type zshBundle struct {
	Project project.Project
}

func (bundle zshBundle) Get() (result string, err error) {
	if err = bundle.Project.Download(); err != nil {
		return result, err
	}
	info, err := os.Stat(bundle.Project.Path())
	if err != nil {
		return "", err
	}
	// it is a file, not a folder, so just return it
	if info.Mode().IsRegular() {
		// XXX: should we add the parent folder to fpath too?
		return "source " + compatible.PathFix(bundle.Project.Path()), nil
	}
	for _, glob := range []string{"*.plugin.zsh", "*.zsh", "*.sh", "*.zsh-theme"} {
		files, err := filepath.Glob(filepath.Join(bundle.Project.Path(), glob))
		if err != nil {
			return result, err
		}
		if files == nil {
			continue
		}
		var lines []string
		for _, file := range files {
			file  = compatible.PathFix(file)
			lines = append(lines, "source "+file)
		}
		lines = append(lines, fmt.Sprintf("fpath+=( %s )", compatible.PathFix(bundle.Project.Path())))
		//fmt.Printf("debug result: %#v\n", lines)
		return strings.Join(lines, "\n"), err
	}

	//fmt.Printf("debug result: %#v\n", result)
	return result, nil
}
