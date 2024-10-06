package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
	"strings"
)

func main() {
	verificado, usuario := autenticacion();
	if verificado {
		fmt.Printf("\nUsuario %s validado.\n", usuario);
	} else {
		fmt.Println("\nHay un error con el usuario y/o la clave suministrados.");
		os.Exit(0);
	}

	now := time.Now();
	year, _, _ := now.Date();
	if year > 2024 /*|| month.String() > "December"*/ {
		fmt.Println("\nEl perìodo de prueba ha terminado.\n");
		os.Exit(0);
	}

	if usuario == "admin" {
		time.Sleep(1 * time.Second);
		fmt.Println("\nContinuamos al módulo de administración de usuarios.\n");
		time.Sleep(1 * time.Second);
		os.Exit(0);
	}

	var respaldo = false;
	archivo_configuracion, err := os.Open("configuracion.txt");
	if err != nil {
		fmt.Printf("\nHubo un error leyendo el archivo de configuración: %s.\n", err);
		time.Sleep(1 * time.Second);
		fmt.Println("\nSaliendo del sistema.\n");
		os.Exit(1);
	}
	defer archivo_configuracion.Close();

	scanner := bufio.NewScanner(archivo_configuracion);
	scanner.Split(bufio.ScanLines);
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=");
		if parts[0] == "respaldo" && parts[1] == "true" {
			respaldo = true;
			break;
		}
	}

	var sesion string;
	if respaldo {
		sesion = fmt.Sprintf("usuarios/%s/sesiones/%s.txt", usuario, now.Format("2006-01-02_03:04"));
		archivoSesion, err := os.Create(sesion);
		if err != nil {
			fmt.Println("\nEl archivo de respaldo no pudo ser creado.\n");
			time.Sleep(1 * time.Second);
			fmt.Println("\nSaliendo del sistema.\n");
			os.Exit(1);
		}
		archivoSesion.Close();
	}

	prompt := "\nPor favor indica una de las siguientes opciones:\n\n  1. Práctica.\n  2. Tarea.\n  3. Salir del sistema.\n\nOpción:";
	sleep();
	var seleccion int;
	for {
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		seleccion = obtenerEntradaUsuario(sesion, respaldo);
		switch seleccion {
			case 1:
				prompt = "\nContinuamos al módulo de práctica.";
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				time.Sleep(1 * time.Second);
				practica(sesion, respaldo);
			case 2:
				prompt = "\nContinuamos al módulo de tareas.\n";
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				sleep();
				os.Exit(0);
			case 3:
				prompt = "\nSaliendo del sistema.\n";
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				sleep();
				os.Exit(0);
			default:
				prompt2 := "\nOpción no válida. Indique la selección nuevamente.\n";
				fmt.Println(prompt2);
				if respaldo {
					archivoAgregar(sesion, prompt2);
				}
				sleep();
				continue;
		}

	}
}
