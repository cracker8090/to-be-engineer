package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Param  []byte
	Result chan []byte
}

func main() {
	pool, _ := ants.NewPoolWithFunc(100000, func(payload interface{}) {
		request, ok := payload.(*Request)
		if !ok {
			return
		}
		reverseParam := func(s []byte) []byte {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			fmt.Println("reverseParam:",s)
			return s
		}(request.Param)
		request.Result <- reverseParam
	})
	defer pool.Release()

	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hhh:",r.Body)
		fmt.Println("hhh:",r.URL.Path)
		fmt.Println("hhh:",r.Form)
		param, err := ioutil.ReadAll(r.Body)
		//param:=([]byte)(r.URL.Path)
		if err != nil {
			http.Error(w, "request error", http.StatusInternalServerError)
		}
		defer r.Body.Close()
		request := &Request{Param: param, Result: make(chan []byte)}

		if err := pool.Invoke(request); err != nil {
			http.Error(w, "throttle limit error", http.StatusInternalServerError)
		}
		fmt.Println(<-request.Result)

		w.Write(<-request.Result)
		//w.Write([]byte("10000"))
	})

	http.ListenAndServe(":8090", nil)

}
