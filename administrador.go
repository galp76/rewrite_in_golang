package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"strings"
	"sync"
)

// de las mias
func cargarHtml(archivo string) ([]byte, error) {
	var mu sync.Mutex;
	mu.Lock();
	html, err := os.ReadFile(archivo);
	mu.Unlock();
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
	var mu sync.Mutex;
	archivo, err := os.Create("html/usuarios/crearUsuario/crearUsuarioTemp.html");
	if err != nil {
		log.Fatal(err);
	}
	archivo.Close();
	primeraParte, err2 := fileToSlice("html/usuarios/crearUsuario/crearUsuarioPrimeraMitad.html");
	if err2 != nil {
		log.Fatal(err);
	}
	mu.Lock();
	for _, item := range primeraParte {
		archivoAgregar("html/usuarios/crearUsuario/crearUsuarioTemp.html", item);
	}
	grupos, err3 := fileToSlice("grupos.txt");
	if err3 != nil {
		log.Fatal(err);
	}
	for _, grupo := range grupos {
		item := fmt.Sprintf("<option value=\"%s\">%s</option>", grupo, grupo);
		archivoAgregar("html/usuarios/crearUsuario/crearUsuarioTemp.html", item);
	}
	segundaParte, err4 := fileToSlice("html/usuarios/crearUsuario/crearUsuarioSegundaMitad.html");
	if err4 != nil {
		log.Fatal(err);
	}
	for _, item := range segundaParte {
		archivoAgregar("html/usuarios/crearUsuario/crearUsuarioTemp.html", item);
	}
	mu.Unlock();
	os.Rename("html/usuarios/crearUsuario/crearUsuarioTemp.html", "html/usuarios/crearUsuario/crearUsuario.html");
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
	var mu sync.Mutex;
	linea := r.URL.Path[len("/procesarNuevoUsuario/"):];
	datos := strings.Split(linea, "/");
	if strings.Contains(datos[0], ";") || strings.Contains(datos[1], ";") || len(datos) != 3 {
		http.Redirect(w, r, "/caracterNoPermitido", 303);
	} else {
		if validarNuevoUsuario(datos[0], datos[2]) {
			nuevoUsuario := fmt.Sprintf("%s;%s;%s", datos[0], datos[1], datos[2]);
			mu.Lock();
			archivoAgregar("users.txt", nuevoUsuario);	// agrega el usuario a users.txt
			mu.Unlock();
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
}

func caracterNoPermitido(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/crearUsuario/caracterNoPermitido.html");
	fmt.Fprintln(w, string(html));
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
	var mu sync.Mutex;
	mu.Lock();
	defer mu.Unlock()
	archivo, err := os.Create("html/usuarios/borrarUsuario/borrarUsuarioTemp.html");
	if err != nil {
		log.Fatal(err);
	}
	archivo.Close();
	primeraParte, err2 := fileToSlice("html/usuarios/borrarUsuario/borrarUsuarioPrimeraMitad.html");
	if err2 != nil {
		log.Fatal(err);
	}
	for _, item := range primeraParte {
		archivoAgregar("html/usuarios/borrarUsuario/borrarUsuarioTemp.html", item);
	}
	grupos, err3 := fileToSlice("grupos.txt");
	if err3 != nil {
		log.Fatal(err);
	}
	for _, grupo := range grupos {
		item := fmt.Sprintf("<option value=\"%s\">%s</option>", grupo, grupo);
		archivoAgregar("html/usuarios/borrarUsuario/borrarUsuarioTemp.html", item);
	}
	segundaParte, err4 := fileToSlice("html/usuarios/borrarUsuario/borrarUsuarioSegundaMitad.html");
	if err4 != nil {
		log.Fatal(err);
	}
	for _, item := range segundaParte {
		archivoAgregar("html/usuarios/borrarUsuario/borrarUsuarioTemp.html", item);
	}
	os.Rename("html/usuarios/borrarUsuario/borrarUsuarioTemp.html", "html/usuarios/borrarUsuario/borrarUsuario.html");
	html, _ := cargarHtml("html/usuarios/borrarUsuario/borrarUsuario.html");
	fmt.Fprintln(w, string(html));
}

func procesarBorrarUsuario(w http.ResponseWriter, r *http.Request) {
	var mu sync.Mutex;
	linea := r.URL.Path[len("/procesarBorrarUsuario/"):];
	datos := strings.Split(linea, "/");

	// validamos que el usuario a borrar existe
	if validarNuevoUsuario(datos[0], datos[1]) {
		http.Redirect(w, r, "/usuarioNoExiste", 303);
	} else {
		// procedenos a borrarlo de users.txt
		mu.Lock();
		defer mu.Unlock();
		archivo, err := os.Create("usuariosTemp.txt");
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
	var mu sync.Mutex;
	mu.Lock();
	defer mu.Unlock()
	archivo, err := os.Create("html/usuarios/mostrarUsuario/mostrarUsuarioTemp.html");
	if err != nil {
		log.Fatal(err);
	}
	archivo.Close();
	primeraParte, err2 := fileToSlice("html/usuarios/mostrarUsuario/mostrarUsuarioPrimeraMitad.html");
	if err2 != nil {
		log.Fatal(err);
	}
	for _, item := range primeraParte {
		archivoAgregar("html/usuarios/mostrarUsuario/mostrarUsuarioTemp.html", item);
	}
	grupos, err3 := fileToSlice("grupos.txt");
	if err3 != nil {
		log.Fatal(err);
	}
	for _, grupo := range grupos {
		item := fmt.Sprintf("<option value=\"%s\">%s</option>", grupo, grupo);
		archivoAgregar("html/usuarios/mostrarUsuario/mostrarUsuarioTemp.html", item);
	}
	segundaParte, err4 := fileToSlice("html/usuarios/mostrarUsuario/mostrarUsuarioSegundaMitad.html");
	if err4 != nil {
		log.Fatal(err);
	}
	for _, item := range segundaParte {
		archivoAgregar("html/usuarios/mostrarUsuario/mostrarUsuarioTemp.html", item);
	}
	os.Rename("html/usuarios/mostrarUsuario/mostrarUsuarioTemp.html", "html/usuarios/mostrarUsuario/mostrarUsuario.html");
	html, _ := cargarHtml("html/usuarios/mostrarUsuario/mostrarUsuario.html");
	fmt.Fprintf(w, string(html));
}

func procesarMostrarUsuario(w http.ResponseWriter, r *http.Request) {
	var mu sync.Mutex;
	grupo := r.URL.Path[len("/procesarMostrarUsuario/"):];

	mu.Lock();
	defer mu.Unlock()
	archivo, err := os.Create("html/usuarios/mostrarUsuario/listaDeUsuariosTemp.html");
	if err != nil {
		log.Fatal(err);
	}
	archivo.Close();
	primeraParte, err2 := fileToSlice("html/usuarios/mostrarUsuario/listaDeUsuariosPrimeraMitad.html");
	if err2 != nil {
		log.Fatal(err);
	}
	for _, item := range primeraParte {
		archivoAgregar("html/usuarios/mostrarUsuario/listaDeUsuariosTemp.html", item);
	}
	item := fmt.Sprintf("<h2>Grupo: %s</h2>\n<table style=\"width:300px\">\n<tr>\n<th>Usuario</th>\n<th>Clave</th>\n</tr>", grupo);
	archivoAgregar("html/usuarios/mostrarUsuario/listaDeUsuariosTemp.html", item);
	usuarios, err3 := fileToSlice("users.txt");
	if err3 != nil {
		log.Fatal(err);
	}
	for _, linea := range usuarios {
		partes := strings.Split(linea, ";");
		if partes[2] == grupo {
			item := fmt.Sprintf("<tr>\n<td>%s</td>\n<td>%s</td>\n</tr>\n", partes[0], partes[1]);
			archivoAgregar("html/usuarios/mostrarUsuario/listaDeUsuariosTemp.html", item);
		}
	}
	segundaParte, err4 := fileToSlice("html/usuarios/mostrarUsuario/listaDeUsuariosSegundaMitad.html");
	if err4 != nil {
		log.Fatal(err);
	}
	for _, item := range segundaParte {
		archivoAgregar("html/usuarios/mostrarUsuario/listaDeUsuariosTemp.html", item);
	}
	os.Rename("html/usuarios/mostrarUsuario/listaDeUsuariosTemp.html", "html/usuarios/mostrarUsuario/listaDeUsuarios.html");
	http.Redirect(w, r, "/listaDeUsuarios", 303);
}

func listaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/mostrarUsuario/listaDeUsuarios.html");
	fmt.Fprintf(w, string(html));
}


// ****************** AQUI COMIENZAN LAS FUNCIONES DE MOSTRAR GRUPOS ******************************


func mostrarGrupos(w http.ResponseWriter, r *http.Request) {
	var mu sync.Mutex;
	mu.Lock();
	defer mu.Unlock()
	archivo, err := os.Create("html/grupos/mostrarGrupos/listaDeGruposTemp.html");
	if err != nil {
		log.Fatal(err);
	}
	archivo.Close();
	primeraParte, err2 := fileToSlice("html/grupos/mostrarGrupos/listaDeGruposPrimeraMitad.html");
	if err2 != nil {
		log.Fatal(err);
	}
	for _, item := range primeraParte {
		archivoAgregar("html/grupos/mostrarGrupos/listaDeGruposTemp.html", item);
	}
	item := "<table style=\"width:300px\">\n<tr>\n<th>Grupo</th>\n</tr>";
	archivoAgregar("html/grupos/mostrarGrupos/listaDeGruposTemp.html", item);
	grupos, err3 := fileToSlice("grupos.txt");
	if err3 != nil {
		log.Fatal(err);
	}
	for _, grupo := range grupos {
		item := fmt.Sprintf("<tr>\n<td>%s</td>\n</tr>\n", grupo);
		archivoAgregar("html/grupos/mostrarGrupos/listaDeGruposTemp.html", item);
	}
	segundaParte, err4 := fileToSlice("html/grupos/mostrarGrupos/listaDeGruposSegundaMitad.html");
	if err4 != nil {
		log.Fatal(err);
	}
	for _, item := range segundaParte {
		archivoAgregar("html/grupos/mostrarGrupos/listaDeGruposTemp.html", item);
	}
	os.Rename("html/grupos/mostrarGrupos/listaDeGruposTemp.html", "html/grupos/mostrarGrupos/listaDeGrupos.html");

	html, _ := cargarHtml("html/grupos/mostrarGrupos/listaDeGrupos.html");
	fmt.Fprintf(w, string(html));
}


// ****************** AQUI COMIENZAN LAS FUNCIONES DE CREAR GRUPO ********************************


func crearGrupo(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/grupos/crearGrupo/crearGrupo.html");
	fmt.Fprintf(w, string(html));
}

// retorna true si el grupo es nuevo
func validarNuevoGrupo(grupo string) bool {
	var mu sync.Mutex;
	mu.Lock();
	defer mu.Unlock()
	grupos, _ := fileToSlice("grupos.txt");
	for _, linea := range grupos {
		partes := strings.Split(linea, ";");
		if partes[0] == grupo {
			return false;
		}
	}

	return true;
}

func procesarNuevoGrupo(w http.ResponseWriter, r *http.Request) {
	var mu sync.Mutex;
	linea := r.URL.Path[len("/procesarNuevoGrupo/"):];
	if strings.Contains(linea, ";") || strings.Contains(linea, "/") {
		http.Redirect(w, r, "/caracterNoPermitido", 303);
	} else {
		if validarNuevoGrupo(linea) {
			mu.Lock();
			defer mu.Unlock()
			archivoAgregar("grupos.txt", linea);
			http.Redirect(w, r, "/grupoCreado", 303);
		} else {
			http.Redirect(w, r, "/grupoYaExiste", 303);
		}
	}
}

func grupoCreado(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/grupos/crearGrupo/grupoCreado.html");
	fmt.Fprintln(w, string(html));
}

func grupoYaExiste(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/grupos/crearGrupo/grupoYaExiste.html");
	fmt.Fprintln(w, string(html));
}

func mainAdministrador() {
//	http.HandleFunc("/view/", viewHandler);
	http.HandleFunc("/", index);		
	http.HandleFunc("/menuPrincipal", index);
	
	// USUARIOS
	http.HandleFunc("/crearUsuario", crearUsuario);
	http.HandleFunc("/procesarNuevoUsuario/", procesarNuevoUsuario);
	http.HandleFunc("/caracterNoPermitido", caracterNoPermitido);
	http.HandleFunc("/usuarioCreado", usuarioCreado);
	http.HandleFunc("/usuarioYaExiste", usuarioYaExiste);

	http.HandleFunc("/borrarUsuario", borrarUsuario);	
	http.HandleFunc("/procesarBorrarUsuario/", procesarBorrarUsuario);
	http.HandleFunc("/usuarioNoExiste", usuarioNoExiste);
	http.HandleFunc("/usuarioBorrado", usuarioBorrado);

	http.HandleFunc("/mostrarUsuario", mostrarUsuario);
	http.HandleFunc("/procesarMostrarUsuario/", procesarMostrarUsuario);
	http.HandleFunc("/listaDeUsuarios", listaDeUsuarios);

	// GRUPOS
	http.HandleFunc("/mostrarGrupos", mostrarGrupos);
	http.HandleFunc("/crearGrupo", crearGrupo);
	http.HandleFunc("/procesarNuevoGrupo/", procesarNuevoGrupo);
	http.HandleFunc("/grupoCreado", grupoCreado);
	http.HandleFunc("/grupoYaExiste", grupoYaExiste);

	fmt.Println("Iniciando servidor...");
	log.Fatal(http.ListenAndServe(":8080", nil));
}
