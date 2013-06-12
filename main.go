package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(log.Lshortfile)

	dirname := "embed"
	embedFilename := "embed.go"
	embedPackage := "main"

	dir, err := os.Open(dirname)
	if err != nil {
		log.Fatalln(err)
	}
	defer dir.Close()

	filenames, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatalln(err)
	}

	embed, err := os.Create(embedFilename)
	if err != nil {
		log.Fatalln(err)
	}
	defer embed.Close()

	fmt.Fprintf(embed, "package %s\n\n", embedPackage)
	fmt.Fprintf(embed, "var embed map[string]string\n\n")

	fmt.Fprintf(embed, "func init() {\n")
	fmt.Fprintf(embed, "\tembed = make(map[string]string)\n\n")
	for _, filename := range filenames {
		path := filepath.Join(dirname, filename)
		log.Println(path)

		contents, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintf(embed, "\tembed[\"%s\"] = %#v\n", filename, string(contents))
	}
	fmt.Fprintf(embed, "}\n")
}
