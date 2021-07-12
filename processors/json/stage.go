package json

const (
	StageIdentifier = "json"

	Encode Mode = "encode"
	Decode Mode = "decode"
)

type Mode string

type Stage struct {
	Mode Mode
}
