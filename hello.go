package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const MONITORING_QUANTITY = 5
const MONITORING_INTERVAL = 5 * time.Second

func main() {
	showStartMessage()

	for {
		showOptions()
		option := scanOptions()
		fmt.Println()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func showStartMessage() {
	name := "Evandro"
	version := 1.1

	fmt.Println("Hello, ", name)
	fmt.Println("Program version: ", version)
	fmt.Println("================================")
}

func showOptions() {
	fmt.Println("Selecione uma opção para iniciar")
	fmt.Println()
	fmt.Println("1. Iniciar monitoramento")
	fmt.Println("2. Ver logs")
	fmt.Println("0. Sair")
	fmt.Println()
}

func scanOptions() int {
	var option int

	fmt.Scan(&option)

	return option
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	fmt.Println()
	sites, error := getFileSiteList()

	if error != nil {
		fmt.Println("")
		fmt.Println("Erro ao tentar ler arquivo de sites:")
		os.Exit(-1)
	}

	for i := 0; i < MONITORING_QUANTITY; i++ {
		for _, site := range sites {
			testSite(site)
		}

		time.Sleep(MONITORING_INTERVAL)
		fmt.Println()
		fmt.Println()
	}
}

func testSite(site string) {
	result, err := http.Get(site)

	if err != nil || result.StatusCode < 200 || result.StatusCode > 299 {
		fmt.Println("O site está fora do ar!")
		return
	}

	fmt.Println("O site ", site, " está online")
}

func getFileSiteList() ([]string, error) {
	file, error := os.ReadFile("sites.txt")

	if error != nil || file == nil {
		fmt.Println(error)
		return nil, error
	}

	return strings.Split(string(file), "\n"), nil
}
