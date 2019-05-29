package main

import (
	"lib/es"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main()  {
	files, _ = filepath.Glob(os.Getenv("STORAGE_ROOT")+"/garbage/*")
	
	for i:range files {
		hash = strings.Split(filepath.Base(files[i]), ".")[0]
		hashInMetadata, e = es.HasHash(hash)
		if e != nil {
			log.Println(e)
			return
		}
		if !hashInMetadata {
			os.Remove(files[i])
		} else {
			log.Println("reload", hash)
			os.Rename(files[i], os.Getenv("STORAGE_ROOT")+"/objects/"+filepath.Base(files[i]))
		}
	}
}