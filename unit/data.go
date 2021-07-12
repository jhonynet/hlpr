package unit

import "github.com/joeycumines/go-dotnotation/dotnotation"

type Data struct {
	TrackingID int
	Value      interface{}
}

func (u *Data) GetProperty(key string) (interface{}, error) {
	return dotnotation.Get(u.Value, key)
}

func (u *Data) SetProperty(key string, value interface{}) (*Data, error) {
	err := dotnotation.Set(u.Value, key, value)

	return u, err
}

func (u *Data) SetValue(value interface{}) *Data {
	u.Value = value

	return u
}
