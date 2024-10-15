package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func tarea(sesion string, respaldo bool, usuario string) int {
	directorio := fmt.Sprintf("./usuarios/%s/tareas", usuario);
	archivos, err := os.ReadDir(directorio);
	if err != nil {
		prompt := fmt.Sprintf("\nHubo un error leyendo el directorio %s: %s.\n", directorio, err);
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		prompt = "\nEl programa debe cerrarse.";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		os.Exit(1);
	}
	var prompt string;
	if len(archivos) == 0 {
		prompt = "\nEl directorio de tareas está vacío.";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		return 1;
	}
	prompt = "\nPor favor selecciona una de las siguientes opciones:\n";
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	sleep();
	for i, archivo := range archivos {
		if i < 9 {
			prompt = fmt.Sprintf("  %d. %s", i + 1, archivo.Name())
		} else {
			prompt = fmt.Sprintf(" %d, %s", i + 1, archivo.Name());
		}
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
	}
	prompt = "\nOpción:";
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	var entradaUsuario int;
	for {
		entradaUsuario = obtenerEntradaUsuario(sesion, respaldo);
		if entradaUsuario >= 1 && entradaUsuario < len(archivos) {
			prompt = "\nSe muestran los ejercicios de la lista seleccionada:\n";
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
			sleep();
			break;
		} else {
			prompt = "\nOpción no válida.\n\nIndique la opción nuevamente:\n";
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
			continue;
		}
	}
	
	directorio = fmt.Sprintf("./usuarios/%s/tareas/%s", usuario, archivos[entradaUsuario - 1].Name());
	operaciones := map[string]string{
		"1": "Suma",
		"2": "Resta",
		"3": "Multiplicación",
		"4": "División",
		"5": "Factores primos",
		"6": "Mínimo Común Múltiplo",
		"7": "Máximo Común Divisor",
		"8": "Operaciones combinadas",
		"9": "Operaciones con fracciones",
		"10": "Suma con decimales",
		"11": "Resta con decimales",
		"12": "Multiplicación con decimales",
		"13": "División con decimales",
	};

	readFile, err := os.Open(directorio);
    defer readFile.Close();
    if err != nil {
		prompt = fmt.Sprintf("\nHubo un error leyendo el archivo %s: %s.", directorio, err);
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		prompt = "\nEl programa debe cerrarse.";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		os.Exit(1);
    }
    fileScanner := bufio.NewScanner(readFile);
    fileScanner.Split(bufio.ScanLines);
	i := 0;
    for fileScanner.Scan() {
        var line = fileScanner.Text();
		parts := strings.Split(line, " ");
		if i < 9 {
			prompt = fmt.Sprintf("  %d. %s: %s - %s", i + 1, operaciones[parts[0]], parts[1], parts[2]);
		} else {
			prompt = fmt.Sprintf(" %d. %s: %s - %s", i + 1, operaciones[parts[0]], parts[1], parts[2]);
		}
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		i++;
    }

	return 0;
}
