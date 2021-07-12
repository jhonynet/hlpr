package file

const (
	StageIdentifier = "file"
)

// File stage params.
type Stage struct {
	Path  string
	Split bool
	Mode  string
	Flags string
}
