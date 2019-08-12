package plumber

// Plumber the interface is the service which creates Pipelines
type Plumber interface {
	Create(*PipelineOptions) (*PipelineOptions, error)
}