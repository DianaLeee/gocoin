package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000";

// omitempty -> hide the field
type URLDescription struct {
	URL string `json:"url"` // JSON일 때 (소문자)url로 변경된다는 뜻
	Method string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"` // Field가 비어있으면 field를 숨겨줌.
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: "/",
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: "/blocks",
			Method: "POST",
			Description: "Add a Block",
			Payload: "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json") // Inform to header that return content type is json
	json.NewEncoder(rw).Encode(data); // Parsing data as JSON type
}

func main() {
	http.HandleFunc("/", documentation);
	fmt.Printf("Listening on http://localhost%s\n", port);
	log.Fatal(http.ListenAndServe(port, nil));

}