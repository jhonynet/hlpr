package unit

type Error struct {
	Err error
}

func ErrorFrom(err error) Error {
	return Error{
		Err: err,
	}
}
