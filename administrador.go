package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

// de las mias
func cargarHtml(archivo string) ([]byte, error) {
	html, err := os.ReadFile(archivo);
	if err != nil {
		return nil, err;
	}

	return html, nil;
}
/*
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):];
	p, _ := loadPage(title);
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body);
}
*/
// de las mias
func index(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/index.html");
	fmt.Fprintf(w, string(html));
}

func crearUsuario(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/crearUsuario.html");
	fmt.Fprintf(w, string(html));
}

func procesarNuevoUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World</h1>");
}

func main() {
//	http.HandleFunc("/view/", viewHandler);
	http.HandleFunc("/", index);		// de las mias
	http.HandleFunc("/crearUsuario", crearUsuario);		// de las mias
	http.HandleFunc("/procesarNuevoUsuario/", procesarNuevoUsuario);		// de las mias
	fmt.Println("Iniciando servidor...");
	log.Fatal(http.ListenAndServe(":8080", nil));
}
