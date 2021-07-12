package http

const (
	StageIdentifier = "http"
)

// Processor stage params.
type Stage struct {
	Method  string
	Content string
	URL     string
	Body    string
	Threads int
	Headers map[string]string
}
