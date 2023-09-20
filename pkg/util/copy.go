package util

import (
	"fmt"
	"github.com/jinzhu/copier"
)

func Copy[T interface{}](fromValue interface{}) (T, error) {
	t := new(T)
	err := copier.CopyWithOption(t, fromValue, copier.Option{IgnoreEmpty: true})
	return *t, fmt.Errorf("util.Copy error:%+v", err)
}
