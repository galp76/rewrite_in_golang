package main

import (
	"fmt"
	"strings"
	"time"
	"strconv"
	"os"
)

func obtenerEntradaUsuarioResta(prompt string, operador string, sesion string, respaldo bool) string {
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

func compararValorResta(tmpTotal int, prompt string, sesion string, respaldo bool) int {
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

// ****** La función sumaResta hace falta para el caso de más de 1 sustraendo

// retorna 0 si el ejercicio se terminó correctamente, y retorna 1
// si el usuario decidió no terminar el ejercicio
func sumaResta(operandos []string, sesion string, respaldo bool) int {
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

func resta(operandos []string, longitudOriginal int, sesion string, respaldo bool) {
	ejercicio := nuevaResta(operandos);
	// No se maneja el error porque se sabe que ambos operandos son numeros
	var numeros [2]int;
	numeros[0], _ = strconv.Atoi(operandos[0]);
	numeros[1], _ = strconv.Atoi(operandos[1]);
	// longitudmaxima: para controlar el "while" principal
	longitudMaxima := 0;
	for _, item := range operandos {
		if len(item) > longitudMaxima {
			longitudMaxima = len(item);
		}
	}
	total := numeros[0] - numeros[1];
	var prompt string;
	if total < 0 {
		prompt = fmt.Sprintf("\nEl resultado es negativo: %s.", total);
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}	
		sleep();
		prompt = "Verifica el ejercicio e intenta nuevamente.";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}	
		os.Exit(0);
	}
	contadorMinuendoModificado := 0;
	sleep();
	if longitudOriginal == 2 {
		prompt = "\nVamos a realizar el siguiente ejercicio:";
	} else {
		prompt = "\nAhora continuamos con la resta:";
	}
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	sleep();
	ejercicio.mostrarResta(sesion, respaldo);
	for longitudMaxima > 0 {
		totalTemporal := numeros[0]%10 - numeros[1]%10;
		stringTemporal := fmt.Sprintf("%s - %s", numeros[0]510, numeros[1]%10);
		if totalTemporal < 0 {
			if !ejercicio.mostrarMinuendoMod {
				ejercicio.mostrarMinuendoMod = true;
				ejercicio.minuendoModificado.prefix(" ");
			}
			sleep();
            prompt = fmt.Sprintf("\nComo %s es menor que %s, pedimos prestado a la izquierda, y continuamos:", numeros[0]%10, numeros[1]%10);
			fmt.Printf(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
// HAY QUE ESCRIBIR MODIFY_MINUEND PARA DESPUES CONTINUAR EN LA LINEA 146 DE RUST/subtraction/lib.rs
		}
	}
}

func mainResta(sesion string, respaldo bool) {
	var operandos []string;
	sleep();
	prompt := "\nPor favor escribe la operación, sin espacios, letras ni caracteres especiales.\n\nEn cualquier momento puedes introducir la letra \"s\" si no deseas terminar con el ejercicio.\n\nEjemplo:\n\t12345-6789-789-145\n";
	for {
		operacion := obtenerEntradaUsuarioResta(prompt, "-", sesion, respaldo);
		operandos = strings.Split(operacion, "-");

		if len(operandos) == 1 {
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
	ejercicio := nuevaResta(operandos);
	longitudOriginal := len(operandos);
	if len(operandos) > 2 {
		sleep();
		prompt = "\nVamos a realizar el siguiente ejercicio:";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		ejercicio.mostrarResta(sesion, respaldo);
		sleep();
		prompt = "\nPero primero vamos a totalizar los sustraendos.";
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
		sleep();
		sumaResta(operandos[1:], sesion, respaldo);
		var minuendo int;
		for i := 1; i < len(operandos); i++ {
			numero, _ := strconv.Atoi(operandos[i]);
			minuendo += numero;
		}
		operandos[1] = strconv.Itoa(minuendo);
	}

	resta(operandos[:2], longitudOriginal, sesion, respaldo);
}
