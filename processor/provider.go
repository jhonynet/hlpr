package processor

type Provider struct {
	Source  Source
	Mappers []Map
	Sink    Sink
}

func (r *Provider) AddMapper(p Map) *Provider {
	r.Mappers = append(r.Mappers, p)

	return r
}
