package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/DianaLeee/gocoin/blockchain"
)

const (
	templateDir string = "explorer/templates/"
)
var templates *template.Template


type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

// Render http template
func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			templates.ExecuteTemplate(rw, "add", nil)
		case "POST":
			r.ParseForm()
			data := r.Form.Get("blockData");
			blockchain.GetBlockchain().AddBlock(data);
			http.Redirect(rw, r, "/", http.StatusPermanentRedirect)

	}
}

func Start(aPort int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml")); 
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home);
	handler.HandleFunc("/add", add);

	fmt.Printf("Listening on http://localhost:%d\n", aPort);
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", aPort), handler))
}

