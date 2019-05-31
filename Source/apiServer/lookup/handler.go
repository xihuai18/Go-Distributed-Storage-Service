// 查询，读取全部文件的名字，返回
package lookup

import (
	"net/http"
	"strings"
	"lib/es"
	"log"
	"io"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	objs, e := es.SearchAllObjects()
	if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
	}
	for i := range(objs) {
		reader := strings.NewReader(objs[i]+"\n")
		_, e = io.Copy(w, reader)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}