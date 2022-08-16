/**
* Program starts a server and apply url mapping
*
* You can find go mod, Yaml, Json, Enum, BoltDB SQL, http and tests examples
*
* Note! All the handlers use MapHandler under the hood
 */

package main

import (
	"fmt"
	"net/http"
	"urlshortener/parsers"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	return mux
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func main() {
	file, err := parsers.ParseCmdFlags()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	urlsPaths, err := parsers.ParseFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mux := defaultMux()
	mapHandler := MapHandler(urlsPaths, mux)
	fmt.Println("Map data: " , urlsPaths)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}
