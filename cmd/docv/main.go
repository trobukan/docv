package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	port := flag.String("port", "8080", "Puerta para el servidor HTTP")
	noBrowser := flag.Bool("no-browser", false, "No abrir el navegador automatico")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Uso: docv [opciones] <nombre-del-doc>")
		fmt.Println("Ejemplo: docv -port=9090 flatpak")
		return
	}

	docPath := filepath.Join("/usr", "share", "doc")

	targetDir := docPath
	if len(flag.Args()) > 0 {
		targetDir = filepath.Join(docPath, flag.Arg(0))
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		log.Fatalf("Error: La ruta '%s' no existe en el sistema", targetDir)
	}

	url := fmt.Sprintf("http://localhost:%s", *port)

	if !*noBrowser {
		go func() {
			cmd := exec.Command("xdg-open", url)
			_ = cmd.Run()
		}()
	}

	fmt.Printf("Docv sirviendo %s\n", targetDir)
	fmt.Printf("Accede en %s\n", url)

	if err := http.ListenAndServe(":"+*port, http.FileServer(http.Dir(targetDir))); err != nil {
		log.Fatal()
	}
}
