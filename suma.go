package main

import (
	"fmt"
)

func obtenerEntradaUsuarioSuma(prompt string, operador string, sesion string, respaldo bool) {
	var entrada string;
	fmt.Println(prompt);
	if respaldo {
		archivoRespaldo(sesion, prompt);
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
// LLEGAMOS HASTA LA LINEA 55 DE RUST/sum_rust/lib.rs
	}
}

func main(sesion, respaldo) {
	var operandos []string;
	sleep();
	prompt := "\nEn cualquier momento puedes introducir la letra \"s\" si no deseas terminar el ejercicio.\n\nPor favor escribe la operación sin espacios, letras ni caracteres especiales.\n\nEjemplo:\n\t12345+6789+12345+78965\n";
	for {
		operacion := obtenerEntradaUsuarioSuma(prompt, "+", sesion, respaldo);
		break;
	}
}
