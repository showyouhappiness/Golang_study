package http_web_server

import (
	"log"
	"netchan"
)

func exporter() {
	exp, err := netchan.NewExporter("tcp", "netchanserver.mydomain.com:1234")
	if err != nil {
		log.Fatalf("Error making Exporter: %v", err)
	}
	ch := make(chan myType)
	err := exp.Export("sendmyType", ch, netchan.Send)
	if err != nil {
		log.Fatalf("Send Error: %v", err)
	}
}

func importer() {
	imp, err := netchan.NewImporter("tcp", "netchanserver.mydomain.com:1234")
	if err != nil {
		log.Fatalf("Error making Importer: %v", err)
	}
	ch := make(chan myType)
	err = imp.Import("sendmyType", ch, netchan.Receive)
	if err != nil {
		log.Fatalf("Receive Error: %v", err)
	}
}
