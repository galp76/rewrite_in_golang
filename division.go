package main

import (
	"fmt"
	"strings"
	"time"
	"strconv"
)

func obtenerEntradaUsuarioDivision(prompt string, operador string, sesion string, respaldo bool) string {
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

func compararValorDivision(tmpTotal int, prompt string, sesion string, respaldo bool) int {
	for {
		entradaUsuario := obtenerEntradaUsuarioDivision(prompt, "", sesion, respaldo);
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

func division(operandos []string, sesion string, respaldo bool) int {
	var vectorDividendo []string;
	for i := 0; i < len(operandos[0]); i++ {
		vectorDividendo = append(vectorDividendo, operandos[0][i:i+1]);
	}
	// i: para controlar el que se bajaa en cada iteracion
	var i int;
	// numero: el dividendo parcial a usar en cada iteracion
	numero, _ := strconv.Atoi(vectorDividendo[i]);
	divisor, _ := strconv.Atoi(operandos[1]);
	for numero/divisor < 1 {
		i++;
		valorTemporal, _ := strconv.Atoi(vectorDividendo[i]);
		numero = numero * 10 + valorTemporal;
	}
	ejercicio := nuevaDivision(operandos, i + 1);
	sleep();
	prompt := "\nVamos a realizar el siguiente ejercicio:";
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	sleep();
	ejercicio.mostrarDivision(sesion, respaldo);
	sleep();
	prompt = fmt.Sprintf("\nPara comenzar, necesitamos un numero que sea mayor o igual a %d.", divisor);
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	sleep();
	primeraIteracion := true;
	// espacios: para controlar el left padding de los dividendos
	espacios := 0;
	for {
		// cocienteTemporal: el cociente temporal a usar en cada iteracion
		cocienteTemporal := numero / divisor;
		prompt = fmt.Sprintf("\nNecesitamos un número que multiplicado por %d sea igual o lo más cercano posible a %d.\nCual es ese número?", divisor, numero);
		compararValorDivision(cocienteTemporal, prompt, sesion, respaldo);
		ejercicio.cociente.postfix(strconv.Itoa(cocienteTemporal));
		sleep();
		prompt = "\nCorrecto.";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		prompt = "\nY cuanto nos queda de residuo?";
		compararValorDivision(numero % divisor, prompt, sesion, respaldo);
		sleep();
		prompt = "\nCorrecto.";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		i++;
		numero = numero % divisor;
		espacios = i - len(strconv.Itoa(numero));
		if i == len(operandos[0]) {
			prompt = "\nHemos terminado con el ejercicio.";
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
			residuo := nuevaLinea(" ", 5 + espacios, strconv.Itoa(numero), 0, " ");
			ejercicio.operaciones = append(ejercicio.operaciones, residuo);
			sleep();
			ejercicio.mostrarDivision(sesion, respaldo);
			return 0;
		}
		var residuoCero bool;
		if numero == 0 {
			residuoCero = true;
		}
		switch primeraIteracion {
		case true:
			ejercicio.operaciones[0].reemplazar(strconv.Itoa(numero));
			ejercicio.operaciones[0].postfix(vectorDividendo[i]);
			primeraIteracion = false;
			// no se maneja el error porque sabemos que es un numero
			numeroTemporal, _ := strconv.Atoi(vectorDividendo[i]);
			numero = numero * 10 + numeroTemporal;
		case false:
			var lineaTemporal Linea;
			if residuoCero {
				lineaTemporal = nuevaLinea(" ", 5 + espacios, "0", 10, " ");
				lineaTemporal.postfix(vectorDividendo[i]);
				// no se nameja el error porque sabemos que es un numero
				numero, _ = strconv.Atoi(vectorDividendo[i]);
			} else {
				numeroTemporal, _ := strconv.Atoi(vectorDividendo[i]);
				numero = numero * 10 + numeroTemporal;
				lineaTemporal = nuevaLinea(" ", 5 + espacios, strconv.Itoa(numero), 0, " ");
			}
			ejercicio.operaciones = append(ejercicio.operaciones, lineaTemporal);
		}
		prompt = fmt.Sprintf("\nBajamos el %s y continuamos.", vectorDividendo[i]);
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		ejercicio.mostrarDivision(sesion, respaldo);
	}
}

func mainDivision(sesion string, respaldo bool, tarea bool, operacion string) int {
	var operandos []string;
	sleep();
	if !tarea {
		prompt := "\nPor favor escribe la operación, sin espacios, letras ni caracteres especiales.\n\nEn cualquier momento puedes introducir la letra \"s\" si no deseas terminar con el ejercicio.\n\nEjemplo:\n\t12345/67\n";
		for {
			operacion := obtenerEntradaUsuarioDivision(prompt, "/", sesion, respaldo);
			operandos = strings.Split(operacion, "/");

			if len(operandos) != 2 {
				sleep();
				advertencia := "\nLa cantidad de operandos no es la correcta.";
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
		operandos = strings.Split(operacion, "/");
	}

	control := division(operandos, sesion, respaldo);
	return control;
}
