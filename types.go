package main

import (
	"fmt"
	"strings"
)

// AQUI COMIENZAN LA DEFINICION Y LAS FUNCIONES DEL struct LINEA
type Linea struct {
	sepIzquierdo string;
	padIzquierdo int;
	contenido string;
	padDerecho int;
	sepDerecho string;
}

func nuevaLinea(sepIzq string, padIzq int, contenido string, padDer int, sepDer string) Linea {
	var resultado Linea;
	resultado.sepIzquierdo = sepIzq;
	resultado.padIzquierdo = padIzq;
	resultado.contenido = contenido;
	resultado.padDerecho = padDer;
	resultado.sepDerecho = sepDer;

	return resultado;
}

func (linea Linea) construir() string {
/*	prompt := fmt.Sprintf("\npadIzquierdo: %d\npadDerecho: %d\n", linea.padIzquierdo, linea.padDerecho);
	fmt.Println(prompt);*/
	resultado := fmt.Sprintf("%s%s%s%s%s", 
		linea.sepIzquierdo,
		strings.Repeat(" ", linea.padIzquierdo),
		linea.contenido,
		strings.Repeat(" ", linea.padDerecho),
		linea.sepDerecho);

	return resultado;
}

func (linea *Linea) prefix(nuevo string) {
	linea.contenido = fmt.Sprintf("%s%s", nuevo, linea.contenido);
	linea.padIzquierdo -= len(nuevo);
/*	prompt := fmt.Sprintf("\npadIzquierdo: %d\npadDerecho: %d\n", linea.padIzquierdo, linea.padDerecho);
	fmt.Println(prompt);*/
}

func (linea *Linea) postfix(nuevo string) {
	linea.contenido = fmt.Sprintf("%s%s", linea.contenido, nuevo);
	linea.padDerecho -= len(nuevo);
/*	prompt := fmt.Sprintf("\npadIzquierdo: %d\npadDerecho: %d\n", linea.padIzquierdo, linea.padDerecho);
	fmt.Println(prompt);*/
}

func (linea *Linea) reemplazar(nuevo string) {
	longitud := len(linea.contenido);
	linea.contenido = fmt.Sprintf("%s%s", strings.Repeat(" ", longitud - len(nuevo)), nuevo);
/*	prompt := fmt.Sprintf("\npadIzquierdo: %d\npadDerecho: %d\n", linea.padIzquierdo, linea.padDerecho);
	fmt.Println(prompt);*/
}

// AQUI COMIENZA LA DEFINICION Y LAS FUNCIONES DEL struct SUMA
type Suma struct {
	mostrarLlevamos bool;
	lineaLlevamos Linea;
	sumandos []Linea;
	lineaResultado Linea;
}

func nuevaSuma(operandos []string) Suma {
	var resultado Suma;
	resultado.lineaLlevamos = nuevaLinea(" ", 14, "", 5, " ");
	longitud := len(operandos[0]);
	tmp := nuevaLinea(" ", 15 - len(operandos[0]), operandos[0], 3, "+");
	resultado.sumandos = append(resultado.sumandos, tmp);
	for i := 1; i < len(operandos); i++ {
		if len(operandos[i]) > longitud {
			longitud = len(operandos[i]);
		}
		tmp = nuevaLinea(" ", 15 - len(operandos[i]), operandos[i], 5, " ");
		resultado.sumandos = append(resultado.sumandos, tmp);
	}
	tmp = nuevaLinea(" ", 15 - longitud, strings.Repeat("-", longitud), 5, " ");
	resultado.sumandos = append(resultado.sumandos, tmp);
	resultado.lineaResultado = nuevaLinea(" ", 15, "", 5, " ");

	return resultado;
}

func (ejercicio Suma) mostrarSuma(sesion string, respaldo bool) {
	var prompt string;
	if ejercicio.mostrarLlevamos {
		prompt = fmt.Sprintf("\n%s%s", ejercicio.lineaLlevamos.construir(), "<--- Llevamos");
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
	}
	fmt.Println("");
	for i, _ := range ejercicio.sumandos {
		prompt = fmt.Sprintf("%s", ejercicio.sumandos[i].construir());
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
	}
	prompt = ejercicio.lineaResultado.construir();
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
}

// AQUI COMIENZA LA DEFINICION Y LAS FUNCIONES DEL struct RESTA
type Resta struct {
	mostrarMinuendoMod bool;
	minuendoModificado Linea;
	operandos []Linea;
	lineaResultado Linea;
}

func nuevaResta(operandos []string) Resta {
	var resultado Resta;
	resultado.minuendoModificado = nuevaLinea(" ", 15, "", 5, " ");
	longitud := len(operandos[0]);
	resultado.operandos = append(resultado.operandos, nuevaLinea(" ", 15 - len(operandos[0]), operandos[0], 3, "-"));
	for i := 1; i < len(operandos); i++ {
		if len(operandos[i]) > longitud {
			longitud = len(operandos[i]);
		}
		resultado.operandos = append(resultado.operandos, nuevaLinea(" ", 15 - len(operandos[i]), operandos[i], 5, " "));
	}
	resultado.operandos = append(resultado.operandos, nuevaLinea(" ", 15 - longitud, strings.Repeat("-", longitud), 5, " "));
	resultado.lineaResultado = nuevaLinea(" ", 15, "", 5, " ");

	return resultado;
}

func (ejercicio Resta) mostrarResta(sesion string, respaldo bool) {
	var prompt string;
	if ejercicio.mostrarMinuendoMod {
		prompt = fmt.Sprintf("\n%s%s\n", ejercicio.minuendoModificado.construir(), "<--- Minuendo modificado");
		fmt.Printf("%s", prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
	}
	fmt.Println("");
	for i, _ := range ejercicio.operandos {
		prompt = fmt.Sprintf("%s", ejercicio.operandos[i].construir());
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
	}
	prompt = ejercicio.lineaResultado.construir();
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
}

// AQUI COMIENZAN LA DEFINICION Y LAS FUNCIONES DEL struct Multiplicacion
type Multiplicacion struct {
	mostrarLlevamos bool;
	lineaLlevamos Linea;
	operandos []Linea;
	resultados []Linea;
}

func nuevaMultiplicacion(operandos []string) Multiplicacion {
	var resultado Multiplicacion;
	resultado.lineaLlevamos = nuevaLinea(" ", 15, "", 5, " ");
	longitud := len(operandos[0]);
	resultado.operandos = append(resultado.operandos, nuevaLinea(" ", 15 - len(operandos[0]), operandos[0], 3, "*"));
	if longitud < len(operandos[1]) {
		longitud = len(operandos[1]);
	}
	resultado.operandos = append(resultado.operandos, nuevaLinea(" ", 15 - len(operandos[1]), operandos[1], 3, " "));
	resultado.operandos = append(resultado.operandos, nuevaLinea(" ", 15 - longitud, strings.Repeat("-", longitud), 5, " "));

	return resultado;
}

func (ejercicio Multiplicacion) mostrarMultiplicacion(sesion string, respaldo bool) {
	var prompt string;
	if ejercicio.mostrarLlevamos {
		prompt = fmt.Sprintf("\n%s%s\n", ejercicio.lineaLlevamos.construir(), "<--- Llevamos");
		fmt.Printf("%s", prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
	}
	fmt.Println("");
	for i, _ := range ejercicio.operandos {
		prompt = ejercicio.operandos[i].construir();
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
	}
	for i, _ := range ejercicio.resultados {
		prompt = ejercicio.resultados[i].construir();
		fmt.Println(prompt);
		if respaldo {
			archivoAgregar(sesion, prompt);
		}
	}
}

// AQUI COMIENZAN LA DEFINICION Y LAS FUNCIONES DEL struct Division
type Division struct {
	dividendo Linea;
	divisor Linea;
	operaciones []Linea;
	cociente Linea;
}

func nuevaDivision(operandos []string, numero int) Division {
	var resultado Division;
	resultado.dividendo = nuevaLinea(" ", 5, operandos[0], 1, "|");
	resultado.divisor = nuevaLinea(" ", 0, operandos[1], 0, " ");
	resultado.operaciones = append(resultado.operaciones, nuevaLinea(" ", 5, strings.Repeat(" ", numero), len(operandos[0]) - numero, " "));
	resultado.cociente = nuevaLinea(" ", 0, "", 10, " ");

	return resultado;
}

func (division Division) mostrarDivision(sesion string, respaldo bool) {
	prompt := fmt.Sprintf("\n%s%s", division.dividendo.construir(), division.divisor.construir());
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	prompt = fmt.Sprintf("%s|%s", strings.Repeat(" ", 7 + len(division.dividendo.contenido)), strings.Repeat("-", 5 + len(division.divisor.contenido)));
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	prompt = fmt.Sprintf("%s|%s", division.operaciones[0].construir(), division.cociente.construir());
	fmt.Println(prompt);
	if respaldo {
		archivoAgregar(sesion, prompt);
	}
	if len(division.operaciones) > 1 {
		for i := 1; i < len(division.operaciones); i++ {
			prompt = division.operaciones[i].construir();
			fmt.Println(prompt);
			if respaldo {
				archivoAgregar(sesion, prompt);
			}
		}
	}
}
