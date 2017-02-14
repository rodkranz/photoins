// Package models
package models

import "fmt"

type ErrTagNotExist struct {
	UID  int64
	Name string
}

func IsErrTagNotExist(err error) bool {
	_, ok := err.(ErrTagNotExist)
	return ok
}

func (err ErrTagNotExist) Error() string {
	return fmt.Sprintf("Tag does not exist [uid: %d, name: %s]", err.UID, err.Name)
}
