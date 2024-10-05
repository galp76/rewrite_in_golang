package main

import (
	"fmt"
	"time"
	"os"
	"strings"
	"strconv"
)

func obtenerEntradaUsuarioSuma(prompt string, operador string, sesion string, respaldo bool) string {
	var entrada string;
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	for {
		_, err := fmt.Scanln(&entrada);
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

		if entrada == "s" || entrada == "S" {
			if respaldo {
				archivoAgregar(sesion, entrada);
			}
			sleep();
			prompt2 := "\nSaliendo del ejercicio...";
			fmt.Println(prompt2);
			if respaldo {
				archivoAgregar(sesion, prompt2);
			}
			sleep();
			os.Exit(0);
		}
		var repetir = false;
		for i := 0; i < len(entrada); i++ {
			ch := entrada[i:i+1];
			caracteres := fmt.Sprintf("0123456789%s", operador);
			if !strings.Contains(caracteres, ch) {
				sleep();
				if respaldo {
					archivoAgregar(sesion, entrada);
				}
				prompt2 := fmt.Sprintf("\nCaracter no válido encontrado: %s", ch);
				fmt.Println(prompt2);
				if respaldo {
					archivoAgregar(sesion, prompt2);
				}
				sleep();
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				repetir = true;
				break;
			}
		}
		if repetir {
			continue;
		}
		break;
	}
	if respaldo {
		archivoAgregar(sesion, entrada);
	}

	return entrada;
}

func suma(operandos []string, sesion string, respaldo bool) {
	var total int;
	for _, item := range operandos {
		numero, _ := strconv.Atoi(item);
		total += numero;
	}
	ejercicio := nuevaSuma(operandos);
	prompt := "\nVamos a realizar el siguiente ejercicio:";
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	sleep();
	ejercicio.mostrarSuma(sesion, respaldo);
// LLEGAMOS HASTA LA LINEA 127 DE RUST/sum_rust/lib.rs
}

func mainSuma(sesion string, respaldo bool) {
	var operandos []string;
	sleep();
	prompt := "\nEn cualquier momento puedes introducir la letra \"s\" si no deseas terminar el ejercicio.\n\nPor favor escribe la operación sin espacios, letras ni caracteres especiales.\n\nEjemplo:\n\t12345+6789+12345+78965\n";
	for {
		operacion := obtenerEntradaUsuarioSuma(prompt, "+", sesion, respaldo);
		operandos = strings.Split(operacion, "+");

		if len(operandos) == 1 {
			sleep();
			advertencia := "\nLa cantidad de sumandos no es la correcta.";
			fmt.Println(advertencia);
			if respaldo {
				archivoAgregar(sesion, advertencia);
			}
			sleep();
			continue;
		}

		var repetir bool;
		for _, item := range operandos {
			if len(item) == 0 {
				sleep();
				advertencia := "\nLa cantidad de operadores no es la correcta.";
				fmt.Println(advertencia);
				if respaldo {
					archivoAgregar(sesion, advertencia);
				}
				sleep();
				repetir = true;
				break;
			}
		}
		if repetir {
			continue;
		}
		break;
	}

	sleep();
	suma(operandos, sesion, respaldo);
}
