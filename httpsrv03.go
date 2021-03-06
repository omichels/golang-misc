/* RZFeeser | Alta3 Research
   HTTP Server with r.URL.Query()  */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	http.HandleFunc("/foobar", func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("baz")

		fmt.Fprintf(writer, "your ULR param baz was: %s", key)
	})
	fmt.Printf("Starting server at port 8055\n")
	log.Fatal(http.ListenAndServe(":8055", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	// accept the name parameter
	keys, ok := r.URL.Query()["name"]

	name := "guest"

	if ok {

		name = keys[0]
	}

	fmt.Fprintf(w, "Hello %s!\n", name)
}
