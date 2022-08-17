/**
* Program starts a server and apply url mapping
*
* You can find go mod, Yaml, Json, Enum, BoltDB SQL (https://github.com/boltdb/bolt), http and tests examples
 */

package main

import (
	"fmt"
	"net/http"
	"urlshortener/db"
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
	file := parsers.ParseCmdFlags()
	urlsPaths := make(map[string]string)

	if file.Type == parsers.Unknown {
		fmt.Println("no file - using db")

		database, err := db.Open()

		if err != nil {
			fmt.Println(err)
			return
		}
		defer database.Close()

		database.FillDB()

		urlsPaths = database.GetMap()
	} else {
		fmt.Println("file provided")
		
		var err error
		urlsPaths, err = parsers.ParseFile(file)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	
	mapHandler := MapHandler(urlsPaths, defaultMux())

	fmt.Println("Map data: ", urlsPaths)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}
