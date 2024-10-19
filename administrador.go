package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"strings"
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

// ************** AQUI COMIENZAN LAS FUNCIONES PARA CREAR UN USUARIO ******************
func crearUsuario(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/crearUsuario.html");
	fmt.Fprintf(w, string(html));
}

// retorna true si el usuario es nuevo
func validarNuevoUsuario(usuario string) bool {
	usuarios, _ := fileToSlice("users.txt");
	for _, linea := range usuarios {
		partes := strings.Split(linea, ";");
		if partes[0] == usuario {
			return false;
		}
	}

	return true;
}

func procesarNuevoUsuario(w http.ResponseWriter, r *http.Request) {
	linea := r.URL.Path[len("/procesarNuevoUsuario/"):];
	datos := strings.Split(linea, "/");
	if validarNuevoUsuario(datos[0]) {
		nuevoUsuario := fmt.Sprintf("%s;%s;%s", datos[0], datos[1], datos[2]);
		archivoAgregar("users.txt", nuevoUsuario);
//		html, _ := cargarHtml("html/usuarios/crearUsuarioExito.html");
//		fmt.Fprintln(w, string(html));
		http.Redirect(w, r, "html/usuarios/crearUsuarioExito.html", 303);
	}
}

func mainAdministrador() {
//	http.HandleFunc("/view/", viewHandler);
	http.HandleFunc("/", index);		// de las mias
	http.HandleFunc("/crearUsuario", crearUsuario);		// de las mias
	http.HandleFunc("/procesarNuevoUsuario/", procesarNuevoUsuario);		// de las mias
	fmt.Println("Iniciando servidor...");
	log.Fatal(http.ListenAndServe(":8080", nil));
}
