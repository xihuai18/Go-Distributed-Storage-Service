// 删除垃圾桶中的文件
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
		// 若文件被再次上传，则要将文件移入文件存放的合理位置
		if !hashInMetadata {
			os.Remove(files[i])
		} else {
			log.Println("reload", hash)
			os.Rename(files[i], os.Getenv("STORAGE_ROOT")+"/objects/"+filepath.Base(files[i]))
		}
	}
}