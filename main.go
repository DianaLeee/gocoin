package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000";

type URL string;

/*
	interface
	- 붕어빵 틀 같은거지...
	- interface 사용하면 어떻게 출력될지 정할 수도 있음.
	-  Implement하고 싶으면 그냥 사용하면 되고, JAVA처럼 implement 선언해주지 않아도 된다.
*/
// Implement the MarshalText method
// intercepting original MarshalText function
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u);
	return []byte(url), nil;
}

type URLDescription struct {
	URL URL `json:"url"` // JSON일 때 (소문자)url로 변경된다는 뜻
	Method string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"` // Field가 비어있으면 field를 숨겨줌.
}



func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: URL("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: URL("/blocks"),
			Method: "POST",
			Description: "Add a Block",
			Payload: "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json") // Inform to header that return content type is json
	json.NewEncoder(rw).Encode(data); // Parsing data as JSON type
}

func main() {
	fmt.Println(URLDescription{
		URL: "/",
		Method: "GET",
		Description: "See Documentation",
	})
	http.HandleFunc("/", documentation);
	fmt.Printf("Listening on http://localhost%s\n", port);
	log.Fatal(http.ListenAndServe(port, nil));

}