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
	Unknown
)

type Fileinfo struct {
	Name string
	Type Filetype
}

type parserHelperStruct struct {
	Path string
	Url  string
}

func ParseCmdFlags() (file Fileinfo) {
	flag.StringVar(&file.Name, "file", "", "json or yaml file with the url maps")
	flag.Parse()

	if strings.HasSuffix(file.Name, "json") {
		file.Type = Json
		return file
	} else if strings.HasSuffix(file.Name, "yaml") {
		file.Type = Yaml
		return file
	}

	file.Type = Unknown
	return file
}

func ParseFile(file Fileinfo) (map[string]string, error) {
	data, err := os.ReadFile(file.Name)
	if err != nil {
		return nil, err
	}

	var unmarshaler func(in []byte, out interface{}) (err error)
	if file.Type == Yaml {
		unmarshaler = yaml.Unmarshal
	} else if file.Type == Json {
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
