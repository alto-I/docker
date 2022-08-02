package main

import(
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/",
	func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "hello go")
	})
	http.ListenAndServe(":80", nil)
}
