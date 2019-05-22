package main

import (
	"lib/es"
	"log"
	// "fmt"
)

const MIN_VERSION_COUNT = 5

func main() {
	buckets, e := es.SearchVersionStatus(MIN_VERSION_COUNT + 1)
	if e != nil {
		log.Println(e)
		return
	}
	// fmt.Println("%d size", len(buckets))
	for i := range buckets {
		bucket := buckets[i]
		// fmt.Println("%d versions", bucket.Doc_count)
		for v := 0; v < bucket.Doc_count-MIN_VERSION_COUNT; v++ {
			es.DelMetadata(bucket.Key, v+int(bucket.Min_version.Value))
		}
	}
}
