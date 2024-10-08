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
					control := mainSuma(sesion, respaldo);
					// este if/else hay que pasarlo al modulo tareas cuando se implemente
					if control == 0 {
						fmt.Println("\nEl ejercicio fue hecho adecuadamente.");
					} else {
						fmt.Println("\nEl usuario decidio no terminar el ejercicio.");
					}
				case 2:
					mainResta(sesion, respaldo);
				default:
					sleep();
					prompt = "\nPor los momentos solamante está implementada la opción 1.";
					fmt.Println(prompt);
					if respaldo {
						archivoAgregar(sesion, prompt);
					}
					sleep();
					prompt = "\nIndique la opción nuevamente.\n";
					fmt.Println(prompt);
					if respaldo {
						archivoAgregar(sesion, prompt);
					}
					sleep();
					mostrarOpciones(opciones, false, sesion, respaldo);
					continue;
			}
		} else {
				sleep();
                prompt = "\nOpción no válida.";
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				sleep();
				prompt = "\nIndique la opción nuevamente.\n";
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				sleep();
				mostrarOpciones(opciones, false, sesion, respaldo);
				continue;
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
