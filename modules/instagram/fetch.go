// Package instagram
package instagram

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"errors"
	"github.com/rodkranz/photoins/modules/setting"
)

func FetchFromTag(tag string) (b []byte, err error) {
	if len(tag) == 0 {
		return b, errors.New("Tag cannot be empty")
	}

	r, err := http.Get(fmt.Sprintf(setting.UrlInstagram, tag))
	if err != nil {
		return
	}
	defer r.Body.Close()

	b, err = ioutil.ReadAll(r.Body)
	return b, nil
}
