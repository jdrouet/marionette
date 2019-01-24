package model

type Repository struct {
	Projects []Project `json:"projects"`
	context  string
	cache    map[string]Project
}

func (r *Repository) Init() error {
	r.cache = make(map[string]Project)
	for _, project := range r.Projects {
		project.repository = r
		r.cache[project.Name] = project
	}
	return nil
}

func (r *Repository) GetProject(name string) Project {
	return r.cache[name]
}

func (r *Repository) Digest(context string, reference string) ([]Project, error) {
	heads := r.getHeads()
	queue := []string{}
	for _, head := range heads {
		project := r.GetProject(head)
		newQueue, err := project.BuildQueue(queue, context, reference)
		if err != nil {
			return nil, err
		}
		queue = newQueue
	}
	queue = RemoveDuplicates(queue)
	result := []Project{}
	for _, item := range queue {
		result = append(result, r.GetProject(item))
	}
	return result, nil
}

func (r *Repository) getHeads() []string {
	heads := []string{}
	for _, project := range r.Projects {
		heads = append(heads, project.Name)
	}
	for _, project := range r.Projects {
		for _, dep := range project.Dependencies {
			heads = Omit(heads, dep)
		}
	}
	return heads
}
