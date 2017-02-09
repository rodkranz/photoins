// Package instagram
package instagram

import (
	"fmt"
	"net/http"

	"github.com/rodkranz/test/module/config"
	"io/ioutil"
)

func FetchFromTag(tag string) (b []byte, err error) {
	r, err := http.Get(fmt.Sprintf(config.URI_INSTAGRAM_TAG, tag))
	if err != nil {
		return
	}
	defer r.Body.Close()

	b, err = ioutil.ReadAll(r.Body)
	return b, nil
}
