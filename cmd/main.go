package main

import (
	"flag"
	"fmt"
	"os"

	lib "github.com/datagenx/cani/lib"
)

func main() {

	vcshost := flag.String("h", "github.com", "VCS hostname")
	clientid := flag.String("c", "", "Oauth App Client ID")
	secretid := flag.String("s", "", "Oauth App Secret ID")
	redirecturl := flag.String("r", "", "Redirect URL")
	flag.Parse()

	data := lib.NewInput(*vcshost, *clientid, *secretid, *redirecturl)
	err := data.VerifyGit()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("VCS app is valid")
	}

}
