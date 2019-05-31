// 根据请求类型调用函数
package objects

import (
	"net/http"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodDelete {
		fmt.Println("delete objects")
		del(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
