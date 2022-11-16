package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/DianaLeee/gocoin/explorer"
	"github.com/DianaLeee/gocoin/rest"
)

func usage () {
	fmt.Printf("Welcome\n");
	fmt.Printf("Please use the follwing flags:\n");
	fmt.Printf("-port:		Set the PRT of the server\n");
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n");
	os.Exit(0);
}


func Start() {
	if len(os.Args) == 1 {
		usage();
	}
	
	port := flag.Int("port", 4000, "Sets the port of the server (default 4000)"); // 4000 is default value
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'"); // rest is default value

	flag.Parse()

	switch *mode {
		case "rest":
			rest.Start(*port)
		case "html":
			explorer.Start(*port)
		default: 
			usage();
	}


	/* 
		#7.2 parsing flag
			rest := flag.NewFlagSet("rest", flag.ExitOnError)
			portFlag := rest.Int("port", 4000, "Sets the port of the server (default 4000)");

			switch os.Args[1] {
			case "explorer":
				fmt.Printf("Start Explorer\n");
			case "rest":
				rest.Parse(os.Args[2:]);
			default: 
				usage();
			}

			if rest.Parsed() {
				fmt.Println(*portFlag)

				fmt.Println("Start the server");
			}
	*/
}