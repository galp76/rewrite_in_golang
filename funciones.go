package main

import (
	"os"
	"fmt"
	"bufio"
	"time"
	"strconv"
	"strings"
)

func archivoAgregar(sesion, data string) {
	archivoSesion, _ := os.OpenFile(sesion, os.O_APPEND|os.O_WRONLY, 0);
	archivoSesion.WriteString(data);
	archivoSesion.Close();
}

func obtenerEntradaUsuario(sesion string, respaldo bool) int {
	for {
		var entradaUsuario string;
		_, err := fmt.Scanln(&entradaUsuario);
		if err != nil {
			fmt.Println("\nHubo un error leyendo la información.\n");
			if respaldo {
				archivoAgregar(sesion, "\nHubo un error leyendo la información.\n");
			}
			time.Sleep(1 * time.Second);
			fmt.Println("\nIndica la opción nuevamente.\n");
			if respaldo {
				archivoAgregar(sesion, "\nIndica la opción nuevamente.\n");
			}
			continue;
		}

		opcion, err := strconv.Atoi(entradaUsuario);
		if err != nil {
			if respaldo {
				archivoAgregar(sesion, entradaUsuario);
			}
			return 0;
		}
		archivoAgregar(sesion, strconv.Itoa(opcion));
		return opcion;
	}
}

func obtenerOpciones(sesion string, respaldo bool) ([]string, []string) {
	var opciones, binarios []string;
	archivo_opciones, err := os.Open("opciones.txt");
	if err != nil {
		prompt := "\nHubo un problema leyendo el archivo de opciones y el sistema debe cerrarse.\n";
		sleep();
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
	}
	defer archivo_opciones.Close();

	scanner := bufio.NewScanner(archivo_opciones);
	scanner.Split(bufio.ScanLines);
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";");
		opciones = append(opciones, parts[0]);
		binarios = append(binarios, parts[1]);
	}

	return opciones, binarios;
}
