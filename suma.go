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
// LLEGAMOS HASTA LA LINEA 8 DE RUST/sum_rust/lib.rs
	}
}

func mainSuma(sesion, respaldo) {
	var operandos []string;
	sleep();
	prompt := "\nEn cualquier momento puedes introducir la letra \"s\" si no deseas terminar el ejercicio.\n\nPor favor escribe la operación sin espacios, letras ni caracteres especiales.\n\nEjemplo:\n\t12345+6789+12345+78965\n";
	for {
		operacion := obtenerEntradaUsuarioSuma(prompt, "+", sesion, respaldo);
		break;
	}
}
