package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DianaLeee/gocoin/blockchain"
	"github.com/DianaLeee/gocoin/utils"
)


var port string;

type url string;

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u);
	return []byte(url), nil;
}

type urlDescription struct {
	URL url `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription {
		{
			URL: url("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: url("/block"),
			Method: "GET",
			Description: "Get blocks",
		},
		{
			URL: url("/blocks"),
			Method: "POST",
			Description: "Add a Block",
			Payload: "data:string",
		},
		{
			URL: url("/blocks/{id}"),
			Method: "GET",
			Description: "See a Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json") // Inform to header that return content type is json
	json.NewEncoder(rw).Encode(data); // Parsing data as JSON type
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		// decode body of request & put into struct
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func Start (aPort int) {
	handler := http.NewServeMux();

	port = fmt.Sprintf(":%d", aPort)

	handler.HandleFunc("/", documentation);
	handler.HandleFunc("/blocks", blocks);

	fmt.Printf("Listening on http://localhost%s\n", port);
	log.Fatal(http.ListenAndServe(port, handler));
}