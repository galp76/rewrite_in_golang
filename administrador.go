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
	html, _ := cargarHtml("html/usuarios/crearUsuario/crearUsuario.html");
	fmt.Fprintf(w, string(html));
}

// retorna true si el usuario es nuevo, toma en consideracion el usuario y el grupo
func validarNuevoUsuario(usuario string, grupo string) bool {
	usuarios, _ := fileToSlice("users.txt");
	for _, linea := range usuarios {
		partes := strings.Split(linea, ";");
		if partes[0] == usuario && partes[2] == grupo {
			return false;
		}
	}

	return true;
}

func procesarNuevoUsuario(w http.ResponseWriter, r *http.Request) {
	linea := r.URL.Path[len("/procesarNuevoUsuario/"):];
	datos := strings.Split(linea, "/");
	if validarNuevoUsuario(datos[0], datos[2]) {
		nuevoUsuario := fmt.Sprintf("%s;%s;%s", datos[0], datos[1], datos[2]);
		archivoAgregar("users.txt", nuevoUsuario);	// agrega el usuario a users.txt
		err := os.MkdirAll(fmt.Sprintf("usuarios/%s/sesiones", datos[0]), 0750);	// crea el directorio 'sesiones' para el usuario
		if err != nil {
			log.Fatal(err);
		}
		err = os.MkdirAll(fmt.Sprintf("usuarios/%s/tareas", datos[0]), 0750);	// crea el directorio 'tareas' para el usuario
		if err != nil {
			log.Fatal(err);
		}
		http.Redirect(w, r, "/usuarioCreado", 303);
	} else {
		http.Redirect(w, r, "/usuarioYaExiste", 303);
	}
}

func usuarioCreado(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/crearUsuario/usuarioCreado.html");
	fmt.Fprintln(w, string(html));
}

func usuarioYaExiste(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/crearUsuario/usuarioYaExiste.html");
	fmt.Fprintln(w, string(html));
}

// ********** AQUI EMPIEZAN LAS FUNCIONES DE BORRAR USUARIO **************************
func borrarUsuario(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/borrarUsuario/borrarUsuario.html");
	fmt.Fprintln(w, string(html));
}

func procesarBorrarUsuario(w http.ResponseWriter, r *http.Request) {
	linea := r.URL.Path[len("/procesarBorrarUsuario/"):];
	datos := strings.Split(linea, "/");

	// validamos que el usuario a borrar existe
	if validarNuevoUsuario(datos[0], datos[1]) {
		http.Redirect(w, r, "/usuarioNoExiste", 303);
	} else {
		// procedenos a borrarlo de users.txt
		archivo, err := os.Create(fmt.Sprintf("usuariosTemp.txt"));
		if err != nil {
			log.Fatal(err);
		}
		archivo.Close();
		usuarios, err := fileToSlice("users.txt");
		if err != nil {
			log.Fatal(err);
		}
		for _, usuario := range usuarios {
			partes := strings.Split(usuario, ";");
			if partes[0] == datos[0] && partes[2] == datos[1] {
				continue;
			} else {
				archivoAgregar("usuariosTemp.txt", usuario);
			}
		}
		os.Rename("usuariosTemp.txt", "users.txt");

		// procedemos a eliminar el directorio del usuario
		os.RemoveAll(fmt.Sprintf("usuarios/%s", datos[0]));
		http.Redirect(w, r, "/usuarioBorrado", 303);
	}
}

func usuarioNoExiste(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/borrarUsuario/usuarioNoExiste.html");
	fmt.Fprintf(w, string(html));
}

func usuarioBorrado(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/borrarUsuario/usuarioBorrado.html");
	fmt.Fprintf(w, string(html));
}

// ********** AQUI EMPIEZAN LAS FUNCIONES DE MOSTRAR USUARIO **************************
func mostrarUsuario(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/mostrarUsuario/mostrarUsuario.html");
	fmt.Fprintf(w, string(html));
}

func mainAdministrador() {
//	http.HandleFunc("/view/", viewHandler);
	http.HandleFunc("/", index);		// de las mias
	http.HandleFunc("/menuPrincipal", index);		// de las mias
	
	http.HandleFunc("/crearUsuario", crearUsuario);		// de las mias
	http.HandleFunc("/procesarNuevoUsuario/", procesarNuevoUsuario);		// de las mias
	http.HandleFunc("/usuarioCreado", usuarioCreado);		// de las mias
	http.HandleFunc("/usuarioYaExiste", usuarioYaExiste);		// de las mias

	http.HandleFunc("/borrarUsuario", borrarUsuario);		// de las mias
	http.HandleFunc("/procesarBorrarUsuario/", procesarBorrarUsuario);	// de las mias
	http.HandleFunc("/usuarioNoExiste", usuarioNoExiste);	// de las mias
	http.HandleFunc("/usuarioBorrado", usuarioBorrado);	// de las mias

	http.HandleFunc("/mostrarUsuario", mostrarUsuario);	// de las mias

	fmt.Println("Iniciando servidor...");
	log.Fatal(http.ListenAndServe(":8080", nil));
}
