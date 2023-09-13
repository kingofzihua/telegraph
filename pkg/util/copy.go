package util

import "github.com/jinzhu/copier"

func Copy[T interface{}](fromValue interface{}) (T, error) {
	t := new(T)
	err := copier.CopyWithOption(t, fromValue, copier.Option{IgnoreEmpty: true})
	return *t, err
}
