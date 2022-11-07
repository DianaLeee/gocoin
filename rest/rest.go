package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DianaLeee/gocoin/blockchain"
	"github.com/DianaLeee/gocoin/utils"
	"github.com/gorilla/mux"
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

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
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
			URL: url("/blocks/{height}"),
			Method: "GET",
			Description: "See a Block",
		},
	}
	// rw.Header().Add("Content-Type", "application/json") // Inform to header that return content type is json
	json.NewEncoder(rw).Encode(data); // Parsing data as JSON type
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		// decode body of request & put into struct
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	id, err := strconv.Atoi(vars["height"]);
	utils.HandleErr(err);
	block, err := blockchain.GetBlockchain().GetBlock(id);
	encoder := json.NewEncoder(rw);

	// #6.7 Error Handling
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
	
}

// #6.8 TODO: 모든 request에 application.json을 추가해주는 미들웨어 만들기
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	// http.HandlerFunc() -> 함수 call 하는게 아니라 adapter pattern임. 
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}


func Start (aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")

	fmt.Printf("Listening on http://localhost%s\n", port);
	log.Fatal(http.ListenAndServe(port, router));
}