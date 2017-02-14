// Package instagram
package instagram

import (
	"fmt"
	"regexp"
)

func ParseJsonData(b []byte) ([]byte, error) {
	var bs []byte

	reg, err := regexp.Compile("window._sharedData =(.*);</script>")
	if err != nil {
		return bs, fmt.Errorf("Regex Error: %s", err)
	}

	found := reg.FindAllStringSubmatch(string(b), -1)
	if len(found) == 0 {
		return bs, fmt.Errorf("Found Html Error: %s", err)
	}

	if len(found) == 0 {
		return bs, fmt.Errorf("Not Found: %s", err)
	}

	if len(found[0]) == 0 {
		return bs, fmt.Errorf("Not Found: %s", err)
	}

	resultFound := string(found[0][0])
	return []byte(resultFound[len("window._sharedData =") : len(resultFound)-len(";</script>")]), nil
}
