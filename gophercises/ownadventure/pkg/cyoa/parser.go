package cyoa

import (
	"encoding/json"
	"flag"
	"os"
)

func JsonToStory(filename string) (story Story, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(f)
	if err = decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

func ParseCmdArgs() (port int, filename string, useHttpServer bool) {
	port = *flag.Int("port", 3000, "a port to start the CYOA web application on")
	filename = *flag.String("file", "config/gopher.json", "the JSON file with the CYOA story")
	useHttpServer = *flag.Bool("server", true, "true for the http server frontend and false for the cmd frontend")
	flag.Parse()
	return port, filename, useHttpServer
}
