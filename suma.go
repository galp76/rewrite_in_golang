package main

import (
	"fmt"
	"time"
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
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
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
			
			// este return es para indicar qque no se terminó el ejercicio en el caso del módulo Tareas
			// y que no se marque como ejercicio hecho
			return "000";
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

func compararValorSuma(tmpTotal int, prompt string, sesion string, respaldo bool) int {
	for {
		entradaUsuario := obtenerEntradaUsuarioSuma(prompt, "", sesion, respaldo);
		// si entradaUsuario == "000" el usuario eligio no terminar el ejercicio, y en el caso
		// de que sea una tarea, no se debe marcar como ejercicio hecho
		if entradaUsuario == "000" {
			return 1;
		}
		// no se maneja el error porque se sabe que entradaUsuario es un numero
		numero, _ := strconv.Atoi(entradaUsuario);
		if numero == tmpTotal {
			break;
		} else {
			sleep();
			prompt := "\nNo es el número que estamos buscando, por favor intenta de nuevo.";
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
			sleep();
		}
	}

	return 0;
}

// retorna 0 si el ejercicio se terminó correctamente, y retorna 1
// si el usuario decidió no terminar el ejercicio
func suma(operandos []string, sesion string, respaldo bool) int {
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
	var numeros []int;
	for _, item := range operandos {
		numero, _ := strconv.Atoi(item);
		numeros = append(numeros, numero);
	}
	var llevamos int = 0;
	for total != 0 {
		var tmpTotal int;
		var tmpString string;
		for i := 0; i < len(numeros); i++ {
			numero := numeros[i];
			tmpTotal += numero % 10;
			tmpNumero := fmt.Sprintf("%d + ", numero % 10);
			tmpString += tmpNumero;
			numero /= 10;
			numeros[i] = numero;
		}
		sleep();
		prompt2 := fmt.Sprintf("\nCuanto es %s?", tmpString[:len(tmpString) - 3]);
		control := compararValorSuma(tmpTotal, prompt2, sesion, respaldo);
		if control == 1 {
			return 1;
		}
		if ejercicio.mostrarLlevamos {
			sleep();
			fmt.Println("\nCorrecto.");
			if respaldo {
				archivoAgregar(sesion, "\nCorrecto.");
			}
			sleep();
			prompt2 = fmt.Sprintf("Y con %d que llevamos cuanto es?", llevamos);
			tmpTotal += llevamos;
			control := compararValorSuma(tmpTotal, prompt2, sesion, respaldo);
			if control == 1 {
				return 1;
			}
		}
		llevamos = tmpTotal / 10;
		if total >= 10 {
			sleep();
			fmt.Println("\nCorrecto.");
			if respaldo {
				archivoAgregar(sesion, "\nCorrecto.");
			}
			var tmpSuma int;
			for _, numero := range numeros {
				tmpSuma += numero;
			}
			if tmpSuma == 0 {
				sleep();
				prompt := fmt.Sprintf("Colocamos el %d, y terminamos con el ejercicio.", tmpTotal);
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				numeroString := fmt.Sprintf("%d", tmpTotal);
				ejercicio.lineaResultado.prefix(numeroString);
				sleep();
				ejercicio.mostrarSuma(sesion, respaldo);
				return 0;
			}
			ejercicio.lineaLlevamos.prefix(strconv.Itoa(llevamos));
			sleep();
            prompt = fmt.Sprintf("Colocamos el %d y llevamos %d.", tmpTotal % 10, tmpTotal / 10);
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
			sleep();
			prompt = "Continuamos con el ejercicio.";
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
		} else {
			sleep();
			prompt = "\nCorrecto, hemos terminado con el ejercicio.";
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
		}
		ejercicio.lineaResultado.prefix(strconv.Itoa(tmpTotal % 10));
		sleep();
		if !ejercicio.mostrarLlevamos {
			ejercicio.mostrarLlevamos = true;
		}
		ejercicio.mostrarSuma(sesion, respaldo);
		total /= 10;
	}

	return 0;
}

func mainSuma(sesion string, respaldo bool, tarea bool, ejercicio string) int {
	var operandos []string;
	if !tarea {
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
	} else {
		operandos = strings.Split(ejercicio, "+");
	}

	sleep();
	control := suma(operandos, sesion, respaldo);

	return control;
}
