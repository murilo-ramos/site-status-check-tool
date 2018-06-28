package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const cicloMonitoramentos = 4
const delay = 5

func main() {
	exibeIntroducao()

	for {
		var comando = leComando()

		/*if comando == 1 {
			fmt.Println("Monitorando...")
		} else if comando == 2 {
			fmt.Println("Exibindo logs...")
		} else if comando == 0 {
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		} else {
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}*/

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

		fmt.Println()
	}
}

func exibeIntroducao() {
	nome := "Murilo"
	versao := 1.1

	fmt.Println("Olá", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println()
}

func leComando() int {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scan(&comando)

	//fmt.Println("O valor da variável de comando é:", comando)
	fmt.Println()

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i := 0; i < cicloMonitoramentos; i++ {
		for _, site := range sites {
			fmt.Println("Testando site", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		if err == io.EOF {
			break
		}

		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}

	arquivo.Close()

	return sites
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//formato de data bizarra
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo Logs...")

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
