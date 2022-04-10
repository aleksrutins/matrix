package config

type NamedValue[V any] struct {
	Name  string
	Value V
}

type Cell struct {
	Command       NamedValue[[]string]
	Configuration NamedValue[string]
}

func (c MatrixConfig) IterateAll() <-chan Cell {
	ch := make(chan Cell)

	go func(ch chan Cell) {
		defer close(ch)
		for cname, configuration := range c.Configurations {
			for tname, target := range c.Targets {
				ch <- Cell{
					Command: NamedValue[[]string]{
						Name:  tname,
						Value: target,
					},
					Configuration: NamedValue[string]{
						Name:  cname,
						Value: configuration,
					},
				}
			}
		}
	}(ch)

	return ch
}
