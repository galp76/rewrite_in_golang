package main

import (
	"fmt"
	"strings"
)

type Line struct {
	sepIzquierdo string
	padIzquierdo int
	contenido string
	padDerecho int
	sepDerecho string
}

func (line Line) nuevo(sepIzq string, padIzq int, contenido string, padDer int, sepDer string) Line {
	var resultado Line;
	resultado.sepIzquierdo = sepIzq;
	resultado.padIzquierdo = padIzq;
	resultado.contenido = contenido;
	resultado.padDerecho = padDer;
	resultado.sepDerecho = sepDer;

	return resultado;
}

func (line Line) construir() string {
	resultado := fmt.Sprintf("%s%s%s%s%s", 
		line.sepIzquierdo,
		strings.Repeat(" ", line.padIzquierdo),
		line.contenido,
		strings.Repeat(" ", line.padDerecho),
		line.sepDerecho);

	return resultado;
}

func (line *Line) prefix(nuevo string) {
	line.contenido = fmt.Sprintf("%s%s", nuevo, line.contenido);
	line.padIzquierdo -= len(nuevo);
}

func (line *Line) postfix(nuevo string) {
	line.contenido = fmt.Sprintf("%s%s", line.contenido, nuevo);
	line.padIzquierdo -= len(nuevo);
}
