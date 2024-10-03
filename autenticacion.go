package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func autenticacion() (bool, string) {
	var usuario string;
	for {
		fmt.Printf("\nIndique su usuario: ");
		var _, err = fmt.Scanln(&usuario);
		if err != nil {
			fmt.Println("Error procesando la información.");
			continue;
		} else {
			break;
		}
	}

	var clave string;
	for {
		fmt.Printf("\nIndique su clave: ");
		var _, err = fmt.Scanln(&clave);
		if err != nil {
			fmt.Println("Error procesando la información.");
			continue;
		} else {
			break;
		}
	}

	var verified bool;
	readFile, err := os.Open("users.txt");
    if err != nil {
        fmt.Println(err);
    }
    fileScanner := bufio.NewScanner(readFile);
    fileScanner.Split(bufio.ScanLines);
    for fileScanner.Scan() {
        var line = fileScanner.Text();
		var parts []string = strings.Split(line, ";");
		if parts[0] == usuario && parts[1] == clave {
			verified = true;
			break;
		}
    }
    readFile.Close();

	return verified, usuario;
}
