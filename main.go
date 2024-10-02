package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
	"strings"
)

func authentication() (bool, string) {
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

	time.Sleep(1 * time.Second);
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

func main() {
	verificado, usuario := authentication();
	time.Sleep(1 * time.Second);
	if verificado {
		fmt.Printf("\nUsuario %s validado.\n", usuario);
	} else {
		fmt.Println("\nHay un error con el usuario y/o la clave suministrados.");
		os.Exit(0);
	}
}
