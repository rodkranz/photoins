// Package models
package models

import "fmt"

type ErrImageNotExist struct {
	UID         int64
	InstagramID string
}

func IsErrImageNotExist(err error) bool {
	_, ok := err.(ErrImageNotExist)
	return ok
}

func (err ErrImageNotExist) Error() string {
	return fmt.Sprintf("Image does not exist [uid: %d, instagram_id: %s]", err.UID, err.InstagramID)
}
