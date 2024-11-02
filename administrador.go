package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"strings"
)

var archivoTareasGlobal string;

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
	archivo, err := os.Create("html/usuarios/crearUsuario/crearUsuarioTemp.html");
	if err != nil {
		log.Fatal(err);
	}
	archivo.Close();
	primeraParte, err2 := fileToSlice("html/usuarios/crearUsuario/crearUsuarioPrimeraMitad.html");
	if err2 != nil {
		log.Fatal(err);
	}
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
	linea := r.URL.Path[len("/procesarNuevoUsuario/"):];
	datos := strings.Split(linea, "/");
	if strings.Contains(datos[0], ";") || strings.Contains(datos[1], ";") || len(datos) != 3 {
		http.Redirect(w, r, "/caracterNoPermitido", 303);
	} else {
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
	linea := r.URL.Path[len("/procesarBorrarUsuario/"):];
	datos := strings.Split(linea, "/");

	// validamos que el usuario a borrar existe
	if validarNuevoUsuario(datos[0], datos[1]) {
		http.Redirect(w, r, "/usuarioNoExiste", 303);
	} else {
		// procedenos a borrarlo de users.txt
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


// ****************** AQUI COMIENZAN LAS FUNCIONES PARA CAMBIAR UN USUARIO DE GRUPO ******************************


func cambiarGrupo(w http.ResponseWriter, r *http.Request) {
	archivo, err := os.Create("html/usuarios/cambiarGrupo/cambiarGrupoTemp.html");
	if err != nil {
		log.Fatal(err);
	}
	archivo.Close();
	primeraParte, err2 := fileToSlice("html/usuarios/cambiarGrupo/cambiarGrupoPrimeraMitad.html");
	if err2 != nil {
		log.Fatal(err);
	}
	for _, item := range primeraParte {
		archivoAgregar("html/usuarios/cambiarGrupo/cambiarGrupoTemp.html", item);
	}
	grupos, err3 := fileToSlice("grupos.txt");
	if err3 != nil {
		log.Fatal(err);
	}
	for _, grupo := range grupos {
		item := fmt.Sprintf("<option value=\"%s\">%s</option>", grupo, grupo);
		archivoAgregar("html/usuarios/cambiarGrupo/cambiarGrupoTemp.html", item);
	}
	item := "</select><br><br>\n<label for=\"unitOfMeasure\">Nuevo grupo:</label><br>\n<select id=\"nuevoGrupo\">\n";
	archivoAgregar("html/usuarios/cambiarGrupo/cambiarGrupoTemp.html", item);
	for _, grupo := range grupos {
		item := fmt.Sprintf("<option value=\"%s\">%s</option>", grupo, grupo);
		archivoAgregar("html/usuarios/cambiarGrupo/cambiarGrupoTemp.html", item);
	}
	segundaParte, err4 := fileToSlice("html/usuarios/cambiarGrupo/cambiarGrupoSegundaMitad.html");
	if err4 != nil {
		log.Fatal(err);
	}
	for _, item := range segundaParte {
		archivoAgregar("html/usuarios/cambiarGrupo/cambiarGrupoTemp.html", item);
	}
	os.Rename("html/usuarios/cambiarGrupo/cambiarGrupoTemp.html", "html/usuarios/cambiarGrupo/cambiarGrupo.html");
	html, _ := cargarHtml("html/usuarios/cambiarGrupo/cambiarGrupo.html");
	fmt.Fprintf(w, string(html));
}

func procesarCambiarGrupo(w http.ResponseWriter, r *http.Request) {
	linea := r.URL.Path[len("/procesarCambiarGrupo/"):];
	datos := strings.Split(linea, "/");

	// validamos que el usuario a reubicar existe
	if validarNuevoUsuario(datos[0], datos[1]) {
		http.Redirect(w, r, "/usuarioNoExiste", 303);
	} else {
		// procedemos a modificar users.txt
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
				item := fmt.Sprintf("%s;%s;%s", partes[0], partes[1], datos[2]);
				archivoAgregar("usuariosTemp.txt", item);
			} else {
				archivoAgregar("usuariosTemp.txt", usuario);
			}
		}
		os.Rename("usuariosTemp.txt", "users.txt");

		http.Redirect(w, r, "/usuarioReasignado", 303);
	}
}

func usuarioReasignado(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/usuarios/cambiarGrupo/usuarioReasignado.html");
	fmt.Fprintf(w, string(html));
}


// ****************** AQUI COMIENZAN LAS FUNCIONES DE MOSTRAR GRUPOS ******************************


func mostrarGrupos(w http.ResponseWriter, r *http.Request) {
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
	grupos, _ := fileToSlice("grupos.txt");
	for _, linea := range grupos {
		if linea == grupo {
			return false;
		}
	}

	return true;
}

func procesarNuevoGrupo(w http.ResponseWriter, r *http.Request) {
	linea := r.URL.Path[len("/procesarNuevoGrupo/"):];
	if strings.Contains(linea, ";") || strings.Contains(linea, "/") {
		http.Redirect(w, r, "/caracterNoPermitido", 303);
	} else {
		if validarNuevoGrupo(linea) {
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


// ****************** AQUI COMIENZAN LAS FUNCIONES DE BORRAR GRUPO ***********************


func borrarGrupo(w http.ResponseWriter, r *http.Request) {
	archivo, err := os.Create("html/grupos/borrarGrupo/borrarGrupoTemp.html");
	if err != nil {
		log.Fatal(err);
	}
	archivo.Close();
	primeraParte, err2 := fileToSlice("html/grupos/borrarGrupo/borrarGrupoPrimeraMitad.html");
	if err2 != nil {
		log.Fatal(err);
	}
	for _, item := range primeraParte {
		archivoAgregar("html/grupos/borrarGrupo/borrarGrupoTemp.html", item);
	}
	grupos, err3 := fileToSlice("grupos.txt");
	if err3 != nil {
		log.Fatal(err);
	}
	for _, grupo := range grupos {
		item := fmt.Sprintf("<option value=\"%s\">%s</option>", grupo, grupo);
		archivoAgregar("html/grupos/borrarGrupo/borrarGrupoTemp.html", item);
	}
	segundaParte, err4 := fileToSlice("html/grupos/borrarGrupo/borrarGrupoSegundaMitad.html");
	if err4 != nil {
		log.Fatal(err);
	}
	for _, item := range segundaParte {
		archivoAgregar("html/grupos/borrarGrupo/borrarGrupoTemp.html", item);
	}
	os.Rename("html/grupos/borrarGrupo/borrarGrupoTemp.html", "html/grupos/borrarGrupo/borrarGrupo.html");
	html, _ := cargarHtml("html/grupos/borrarGrupo/borrarGrupo.html");
	fmt.Fprintln(w, string(html));
}

func procesarBorrarGrupo(w http.ResponseWriter, r *http.Request) {
	linea := r.URL.Path[len("/procesarBorrarGrupo/"):];

	// procedenos a borrarlo de grupos.txt
	archivo, err := os.Create("gruposTemp.txt");
	if err != nil {
		log.Fatal(err);
	}
	archivo.Close();
	grupos, err := fileToSlice("grupos.txt");
	if err != nil {
		log.Fatal(err);
	}
	for _, grupo := range grupos {
		if linea == grupo {
			continue;
		} else {
			archivoAgregar("gruposTemp.txt", grupo);
		}
	}
	os.Rename("gruposTemp.txt", "grupos.txt");

	// procedemos a eliminar de users.txt todos los usuarios asignados al grupo y sus directorios
	archivo, err = os.Create("usuariosTemp.txt");
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
		if partes[2] == linea {
			// borramos el directorio del usuario
			os.RemoveAll(fmt.Sprintf("usuarios/%s", partes[0]));
			continue;
		} else {
			archivoAgregar("usuariosTemp.txt", usuario);
		}
	}
	os.Rename("usuariosTemp.txt", "users.txt");

	http.Redirect(w, r, "/grupoBorrado", 303);
}

func grupoBorrado(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/grupos/borrarGrupo/grupoBorrado.html");
	fmt.Fprintf(w, string(html));
}


// ********** AQUI EMPIEZAN LAS FUNCIONES DE LISTAR USUARIOS **************************


func mostrarUsuario(w http.ResponseWriter, r *http.Request) {
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
	grupo := r.URL.Path[len("/procesarMostrarUsuario/"):];

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


// ******************* AQUI COMIENZAN LAS FUNCIONES DE CREAR TAREA ***********************


func crearTarea(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/tareas/crearTarea/crearTarea.html");
	fmt.Fprintf(w, string(html));
}

func procesarCrearArchivo(w http.ResponseWriter, r *http.Request) {
	// creamos el archivo txt con el nommbre suministrado
	nombre := r.URL.Path[len("/procesarCrearArchivo/"):];
	nombreArchivo := fmt.Sprintf("tareas/%s.txt", nombre);
	archivoTareasGlobal = nombreArchivo;
	archivoTarea, err := os.Create(nombreArchivo);
	if err != nil {
		log.Fatal(err);
	}
	archivoTarea.Close();

	http.Redirect(w, r, "/crearOperaciones", 303);
}

func crearOperaciones(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/tareas/crearTarea/crearOperaciones.html");
	fmt.Fprintf(w, string(html));
}

func procesarNuevaOperacion(w http.ResponseWriter, r *http.Request) {
	linea := r.URL.Path[len("/procesarNuevaOperacion/"):];
	datosOperacion := strings.Split(linea, "/");
	switch datosOperacion[0] {
	case "Suma":
		operacion := strings.Replace(datosOperacion[1], "%2B", "+", -1);
		// chequeamos si tiene caracteres distintos a los permitidos (numeros del 0 al 9 y el simbolo "+")
		var caracterNoPermitido bool;
		for _, ch := range operacion {
			if !strings.Contains("0123456789+", string(ch)) {
				caracterNoPermitido = true;
				break;
			}
		}
		// chequeamos que tenga al menos 2 operandos validos
		operandos := strings.Split(operacion, "+");
		var operandoVacio bool;
		for _, item := range operandos {
			if len(item) == 0 {
				operandoVacio = true;
			}
		}
		// si len(datosOperacion) > 2 el usuario escribio un "/" en la operacion
		if caracterNoPermitido || len(datosOperacion) > 2 || len(operandos) < 2 || operandoVacio {
			http.Redirect(w, r, "/tareasCaracterNoPermitido", 303);
		} else {
			operacion = fmt.Sprintf("1 %s Pendiente", operacion);
			archivoAgregar(archivoTareasGlobal, operacion);		
			http.Redirect(w, r, "/crearOperaciones", 303);
		}
	case "Resta":
		operacion := datosOperacion[1];
		// buscamos caracteres distintos a numeros (0-9) y el simbolo "-"
		var caracterNoPermitido bool;
		for _, ch := range operacion {
			if !strings.Contains("0123456789-", string(ch)) {
				caracterNoPermitido = true;
				break;
			}
		}
		// chequeamos que tenga al menos 2 operandos validos
		operandos := strings.Split(operacion, "-");
		var operandoVacio bool;
		for _, item := range operandos {
			if len(item) == 0 {
				operandoVacio = true;
			}
		}
		// si len(datosOperacion) > 2 el usuario escribio un "/" en la operacion
		if caracterNoPermitido || len(datosOperacion) > 2 || len(operandos) < 2 || operandoVacio {
			http.Redirect(w, r, "/tareasCaracterNoPermitido", 303);
		} else {
			operacion = fmt.Sprintf("2 %s Pendiente", operacion);
			archivoAgregar(archivoTareasGlobal, operacion);		
			http.Redirect(w, r, "/crearOperaciones", 303);
		}
	case "Multiplicación":
		operacion := datosOperacion[1];
		// buscamos caracteres distintos a numeros (0-9) y el simbolo "*"
		var caracterNoPermitido bool;
		for _, ch := range operacion {
			if !strings.Contains("0123456789*", string(ch)) {
				caracterNoPermitido = true;
				break;
			}
		}
		// chequeamos que la operacion tenga exactamente 2 operandos validos
		operandos := strings.Split(operacion, "*");
		var operandoVacio bool;
		for _, item := range operandos {
			if len(item) == 0 {
				operandoVacio = true;
			}
		}
		// si len(datosOperacion) > 2 el usuario escribio un "/" en la operacion
		if caracterNoPermitido || len(datosOperacion) > 2 || len(operandos) != 2 || operandoVacio {
			http.Redirect(w, r, "/tareasCaracterNoPermitido", 303);
		} else {
			operacion = fmt.Sprintf("3 %s Pendiente", operacion);
			archivoAgregar(archivoTareasGlobal, operacion);		
			http.Redirect(w, r, "/crearOperaciones", 303);
		}
	case "División":
		// este "if" evalua el caso de operaciones como 78956/
		if len(datosOperacion) > 2 && len(datosOperacion[2]) == 0 {
			http.Redirect(w, r, "/tareasCaracterNoPermitido", 303);
		} else {
			var operacion string;
			// construye operacion solo si se tiene el numero correcto de operandos
			if len(datosOperacion) > 2 {
				operacion = fmt.Sprintf("%s/%s", datosOperacion[1], datosOperacion[2]);
			}
			// buscamos caracteres distintos a numeros (0-9) y el simbolo "/"
			var caracterNoPermitido bool;
			for _, ch := range operacion {
				if !strings.Contains("0123456789/", string(ch)) {
					caracterNoPermitido = true;
					break;
				}
			}
			// si len(datosOperacion) > 3 el usuario escribio uno o mas "/" adicionales en la operacion
			// si len(datosOperacion) < 3 el usuario escribio una operacion como /5896
			// por lo tanto len(datosOperacion) debe ser exactamente 3
			if caracterNoPermitido || len(datosOperacion) != 3 {
				http.Redirect(w, r, "/tareasCaracterNoPermitido", 303);
			} else {
				operacion = fmt.Sprintf("4 %s Pendiente", operacion);
				archivoAgregar(archivoTareasGlobal, operacion);		
				http.Redirect(w, r, "/crearOperaciones", 303);
			}
		}
	}
}

func tareasCaracterNoPermitido(w http.ResponseWriter, r *http.Request) {
	html, _ := cargarHtml("html/tareas/crearTarea/caracterNoPermitido.html");
	fmt.Fprintf(w, string(html));
}

func descartarTarea(w http.ResponseWriter, r *http.Request) {
	err := os.Remove(archivoTareasGlobal);
	if err != nil {
		log.Fatal(err);
	}

	http.Redirect(w, r, "/", 303);
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

	http.HandleFunc("/cambiarGrupo", cambiarGrupo);
	http.HandleFunc("/procesarCambiarGrupo/", procesarCambiarGrupo);
	http.HandleFunc("/usuarioReasignado", usuarioReasignado);

	// GRUPOS
	http.HandleFunc("/mostrarGrupos", mostrarGrupos);

	http.HandleFunc("/crearGrupo", crearGrupo);
	http.HandleFunc("/procesarNuevoGrupo/", procesarNuevoGrupo);
	http.HandleFunc("/grupoCreado", grupoCreado);
	http.HandleFunc("/grupoYaExiste", grupoYaExiste);

	http.HandleFunc("/borrarGrupo", borrarGrupo);
	http.HandleFunc("/procesarBorrarGrupo/", procesarBorrarGrupo);
	http.HandleFunc("/grupoBorrado", grupoBorrado);

	http.HandleFunc("/mostrarUsuario", mostrarUsuario);
	http.HandleFunc("/procesarMostrarUsuario/", procesarMostrarUsuario);
	http.HandleFunc("/listaDeUsuarios", listaDeUsuarios);

	// TAREAS
	http.HandleFunc("/crearTarea", crearTarea);
	http.HandleFunc("/procesarCrearArchivo/", procesarCrearArchivo);
	http.HandleFunc("/crearOperaciones", crearOperaciones);
	http.HandleFunc("/procesarNuevaOperacion/", procesarNuevaOperacion);
	http.HandleFunc("/tareasCaracterNoPermitido", tareasCaracterNoPermitido);
	http.HandleFunc("/descartarTarea", descartarTarea);

	fmt.Println("Iniciando servidor...");
	log.Fatal(http.ListenAndServe(":8080", nil));
}
