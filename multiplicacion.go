package main

import (
	"fmt"
	"strings"
	"time"
	"strconv"
)

func obtenerEntradaUsuarioMultiplicacion(prompt string, operador string, sesion string, respaldo bool) string {
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

func compararValorMultiplicacion(tmpTotal int, prompt string, sesion string, respaldo bool) int {
	for {
		entradaUsuario := obtenerEntradaUsuarioMultiplicacion(prompt, "", sesion, respaldo);
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

// ****** La función sumaMultiplicacion hace falta para totalizar al final

// retorna 0 si el ejercicio se terminó correctamente, y retorna 1
// si el usuario decidió no terminar el ejercicio
func sumaMultiplicacion(operandos []string, sesion string, respaldo bool) int {
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

func multiplicacion(operandos []string, sesion string, respaldo bool) int {
	ejercicio := nuevaMultiplicacion(operandos);
	lineaResultadoTemporal := nuevaLinea(" ", 15, "", 5, " ");
	var numeros []int;
	for _, item := range operandos {
		// no se maneja el error porque se sabe que 'item' es un numero
		numeroTemporal, _ := strconv.Atoi(item);
		numeros = append(numeros, numeroTemporal);
	}
	longitudSegundoOperando := len(operandos[1]);
	contador := 0;
	llevamos := ((numeros[0]%10) * (numeros[1]%10)) / 10;
	ejercicio.lineaLlevamos.prefix(" ");
	sleep();
	prompt := "\nVamos a realizar el siguiente ejercicio:";
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	sleep();
	ejercicio.mostrarMultiplicacion(sesion, respaldo);
	for longitudSegundoOperando - contador != 0 {
		valorFijado := numeros[1] % 10;
		primerOperando := numeros[0];
		for primerOperando != 0 {
			totalTemporal := (primerOperando % 10) * valorFijado;
			stringTemporal := fmt.Sprintf("%d * %d", primerOperando % 10, valorFijado);
			primerOperando /= 10;
			prompt = fmt.Sprintf("\nCuanto es %s?", stringTemporal);
			sleep();
			control := compararValorMultiplicacion(totalTemporal, prompt, sesion, respaldo);
			if control == 1 {
				return 1;
			}
			if ejercicio.mostrarLlevamos {
				sleep();
				prompt = "\nCorrecto.";
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				sleep();
				prompt = fmt.Sprintf("\nY con %d que llevamos cuanto es?", llevamos);
				totalTemporal += llevamos;
				control := compararValorMultiplicacion(totalTemporal, prompt, sesion, respaldo);
				if control == 1 {
					return 1;
				}
			}
			llevamos = totalTemporal / 10;
			if !((longitudSegundoOperando - contador == 1) && primerOperando == 0) {
				ejercicio.lineaLlevamos.prefix(strconv.Itoa(llevamos));
				sleep();
				prompt = "\nCorrecto.";
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
				sleep();
				if primerOperando != 0 {
					prompt = fmt.Sprintf("Colocamos el %d y llevamos %d.", totalTemporal % 10, totalTemporal / 10);
					fmt.Println(prompt);
					if respaldo {
						archivoAgregar(sesion, prompt);
					}
				} else {
					prompt = fmt.Sprintf("Colocamos el %d, y continuamos con las multiplicaciones correspondientes al siguiente número: %d", totalTemporal, (numeros[1] / 10) % 10);
					fmt.Println(prompt);
					if respaldo {
						archivoAgregar(sesion, prompt);
					}
				}
				sleep();
				prompt = "Continuamos con el ejercicio.";
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
			} else {
				sleep();
				prompt = fmt.Sprintf("\nCorrecto, colocamos el %d y hemos terminado con las multiplicaciones.", totalTemporal);
				fmt.Println(prompt);
				if respaldo {
					archivoAgregar(sesion, prompt);
				}
			}
			sleep();
			ejercicio.mostrarLlevamos = true;
			if primerOperando != 0 {
				lineaResultadoTemporal.prefix(strconv.Itoa(totalTemporal % 10));
			} else {
				// para que no imprima la linea Llevamos cuando primerOperando se hace cero
				ejercicio.mostrarLlevamos = false;
				lineaResultadoTemporal.prefix(strconv.Itoa(totalTemporal));
			}
			ejercicio.mostrarMultiplicacion(sesion, respaldo);
			prompt = lineaResultadoTemporal.construir();
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
		}

		ejercicio.resultados = append(ejercicio.resultados, lineaResultadoTemporal);
		lineaResultadoTemporal = nuevaLinea(" ", 15, "", 5, " ");
		ejercicio.lineaLlevamos = nuevaLinea(" ", 15, "", 5, " ");
		ejercicio.lineaLlevamos.prefix(" ");
		contador++;
		for i := 0; i < contador; i++ {
			lineaResultadoTemporal.prefix(" ");
		}
		numeros[1] /= 10;
	}
	
	sleep();
	if len(ejercicio.resultados) == 1 {
		prompt = "\nHemos terminado con el ejercicio.";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		return 0;
	}

	var operandosParaLaSuma []string;
	for _, item := range ejercicio.resultados {
		operandosParaLaSuma = append(operandosParaLaSuma, strings.ReplaceAll(item.contenido, " ", "0"));
	}

	prompt = "\nAhora tenemos que sumar los resultados parciales.";
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}

	control := sumaMultiplicacion(operandosParaLaSuma, sesion, respaldo);
	return control;
}

func mainMultiplicacion(sesion string, respaldo bool, tarea bool, operacion string) int {
	var operandos []string;
	sleep();
	if !tarea {
		prompt := "\nPor favor escribe la operación, sin espacios, letras ni caracteres especiales.\n\nEn cualquier momento puedes introducir la letra \"s\" si no deseas terminar con el ejercicio.\n\nEjemplo:\n\t12345*678\n";
		for {
			operacion := obtenerEntradaUsuarioMultiplicacion(prompt, "*", sesion, respaldo);
			operandos = strings.Split(operacion, "*");

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
		operandos = strings.Split(operacion, "*");
	}

	control := multiplicacion(operandos, sesion, respaldo);
	return control;
}
