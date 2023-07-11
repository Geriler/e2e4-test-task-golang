package adapters

import (
	xj "github.com/basgys/goxml2json"
	"strings"
)

func Xml2json(xml string) (string, error) {
	json, err := xj.Convert(strings.NewReader(xml))
	if err != nil {
		return "", err
	}

	return json.String(), nil
}
