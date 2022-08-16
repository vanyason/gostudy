package parsers

import (
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
	f, err := os.ReadFile(file.filename)
	if err != nil {
		return nil, err
	}

	if file.filetype == Yaml {
		return ParseYaml(f)
	} else if file.filetype == Json {
		return ParseJson(f)
	}

	return nil, fmt.Errorf("can not parse unknown file format")
}

func ParseYaml(data []byte) (map[string]string, error) {
	/* create a yaml struct and parse data into it */
	type yamlStruct struct {
		Path string `yaml:"path"`
		Url  string `yaml: "url"`
	}

	var yamlStructs []yamlStruct
	err := yaml.Unmarshal(data, &yamlStructs)
	if err != nil {
		return nil, err
	}

	/* convert yaml struct to a map */
	pathsToUrls := make(map[string]string)
	for _, s := range yamlStructs {
		if s.Path == "" || s.Url == "" {
			return nil, fmt.Errorf("failed while parsing file to the map. empty field")
		}

		pathsToUrls[s.Path] = s.Url
	}

	return pathsToUrls, nil
}

/* To implement */
func ParseJson(data []byte) (map[string]string, error) {
	return nil, nil
}
