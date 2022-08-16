package parsers

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Filetype int64

const (
	Yaml Filetype = iota
	Json
)

type Fileinfo struct {
	filename string
	filetype Filetype
}

type parserHelperStruct struct {
	Path string
	Url  string
}

func ParseCmdFlags() (file Fileinfo, err error) {
	flag.StringVar(&file.filename, "file", "urls.yaml", "json or yaml file with the url maps")
	flag.Parse()

	if strings.HasSuffix(file.filename, "json") {
		file.filetype = Json
		return file, nil
	} else if strings.HasSuffix(file.filename, "yaml") {
		file.filetype = Yaml
		return file, nil
	} else {
		return file, fmt.Errorf("invalid file is provided")
	}
}

func ParseFile(file Fileinfo) (map[string]string, error) {
	data, err := os.ReadFile(file.filename)
	if err != nil {
		return nil, err
	}

	var unmarshaler func(in []byte, out interface{}) (err error)
	if file.filetype == Yaml {
		unmarshaler = yaml.Unmarshal
	} else if file.filetype == Json {
		unmarshaler = json.Unmarshal
	} else {
		return nil, fmt.Errorf("can not parse unknown file format")
	}

	var parserHelperStructs []parserHelperStruct
	err = unmarshaler(data, &parserHelperStructs)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, s := range parserHelperStructs {
		if s.Path == "" || s.Url == "" {
			return nil, fmt.Errorf("failed while parsing file to the map. empty field")
		}

		pathsToUrls[s.Path] = s.Url
	}

	return pathsToUrls, nil
}
