package main

import (
	"fmt"
//	"os"
	"strings"
)

func practica(sesion string, respaldo bool) {
	opciones, _ := obtenerOpciones(sesion, respaldo);
	mostrarOpciones(opciones, true, sesion, respaldo);
	var prompt string;
	var entradaUsuario int;
	for {
		entradaUsuario = obtenerEntradaUsuario(sesion, respaldo);
		if entradaUsuario != 0 && entradaUsuario <= len(opciones) {
			switch entradaUsuario {
			case 1:
				fmt.Println("Llamando mainSuma()");
//				os.Exit(0);
				mainSuma(sesion, respaldo);
			default:
				sleep();
                prompt = "\nOpción no válida.\n\nIndique la opción nuevamente:\n";
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				sleep();
				mostrarOpciones(opciones, false, sesion, respaldo);
				continue;
			}
		}
		sleep();
		prompt = fmt.Sprintf("\n%s", strings.Repeat("*", 105));
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
        prompt = "\nSelecciona una opción para continuar con otro ejercicio, o introduce la letra \"s\" para salir del sistema.\n";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		mostrarOpciones(opciones, false, sesion, respaldo);
	}
}
