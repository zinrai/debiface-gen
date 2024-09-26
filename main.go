package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/zinrai/debiface-gen/api"
	"github.com/zinrai/debiface-gen/cli"
)

func main() {
	serverMode := flag.Bool("server", false, "Run in server mode")
	flag.Usage = func() {
		fmt.Println("debiface-gen: Debian Network Interface Configuration Generator")
		fmt.Println("\nUsage:")
		fmt.Println("  debiface-gen [options] <command> [command options]")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		fmt.Println("\nCommands:")
		fmt.Println("  bonding   Generate bonding configuration")
		fmt.Println("  dsr       Generate DSR configuration")
		fmt.Println("  standard  Generate standard interface configuration")
		fmt.Println("  bridge    Generate bridge configuration")
		fmt.Println("\nUse 'debiface-gen <command> --help' for more information about a command")
	}
	flag.Parse()

	if *serverMode {
		startServer()
	} else {
		cli.Run()
	}
}

func startServer() {
	http.HandleFunc("/api/bonding", api.HandleBonding)
	http.HandleFunc("/api/dsr", api.HandleDSR)
	http.HandleFunc("/api/standard", api.HandleStandard)
	http.HandleFunc("/api/bridge", api.HandleBridge)

	fmt.Println("debiface-gen server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
