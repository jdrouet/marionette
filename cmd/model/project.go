package model

import (
	git "github.com/jdrouet/marionette/cmd/git"
)

type Project struct {
	Name         string   `json:"name"`
	Path         string   `json:"path"`
	Dependencies []string `json:"dependencies,omitempty"`
	//
	repository *Repository
}

func (p *Project) getChangedFiles(context string, reference string) ([]string, error) {
	return git.Diff(context, reference, p.Path)
}

// Get changed files of a project
func (p *Project) Diff(context string, reference string) ([]string, error) {
	changed, err := p.getChangedFiles(context, reference)
	if err != nil {
		return nil, err
	}
	changed = Omit(changed, "")
	// TODO filter changed files
	return changed, nil
}

func (p *Project) HasChanged(context string, reference string) (bool, error) {
	changes, err := p.Diff(context, reference)
	if err != nil {
		return false, err
	}
	return len(changes) > 0, nil
}

func (p *Project) BuildQueue(originalQueue []string, context string, reference string) ([]string, error) {
	queue := append(originalQueue)
	for _, dep := range p.Dependencies {
		child := p.repository.GetProject(dep)
		newQueue, err := child.BuildQueue(queue, context, reference)
		if err != nil {
			return nil, err
		}
		// because I cannot do queue, err = buildQueue
		queue = newQueue
	}
	hasChanged, err := p.HasChanged(context, reference)
	if err != nil {
		return nil, err
	}
	queueChanged := len(queue) != len(originalQueue)
	if hasChanged || queueChanged {
		queue = append(queue, p.Name)
	}
	return queue, nil
}
