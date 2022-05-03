package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os" // Para exibir resultado de saída(exit) do sistema.
	"strconv"
	"strings"

	"time"

	//"reflect" - Para saber qual o tipo da váriavel.
	"bufio" // Para ajudar a ler um arquivo linha por linha
	"net/http"
)

const monitoramento = 5
const delay = 10

func main() {

	exibirIntroducao()

	for {

		exibeMenu()

		comandoLido := leComando()

		switch comandoLido {
		case 1:
			iniciarMonitoramento()
		case 2:
			iniciarLogs()
		case 0:
			saindo()
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}
		fmt.Println("")
		exibirEscolha()

		ccomando := leComandoNovamente()

		switch ccomando {
		case 1:
			fmt.Println("")
			fmt.Println("Voltando ao menu principal ...")
		case 2:
			saindo()
		default:
			comandoNãoIndentificado()
		}
	}
}

// Funções / Métodos

func iniciarMonitoramento() {
	fmt.Println("")
	fmt.Println("Iniciando monitoramento ...")

	site := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range site {
			fmt.Println("")
			fmt.Println("SITE", i)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Arquivo com erro!", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
			//fmt.Println("Arquivo com erro!", err)
		}

	}
	fmt.Println(sites)

	// io.EOF ele encontra o final do arquivo, então quando chegar no final do arquivo, ele automaticamente
	// vai parar e dar um Break

	//arquivo, err := ioutil.ReadFile("sites.txt")
	//fmt.Println(string(arquivo))

	arquivo.Close()

	return sites
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Arquivo com erro!", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println(site, " carregado com sucesso!")
		fmt.Println("StatusCode: ", resp.StatusCode)
	} else {
		fmt.Println(site)
		fmt.Println("ERRO: ", resp.StatusCode)
	}
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func iniciarLogs() {
	fmt.Println("")
	fmt.Println("Exibindo Logs ...")

	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("ERRO: ", err)
	}
	fmt.Println(string(arquivo))

}

func exibirIntroducao() {
	nome := "Guilherme"
	fmt.Println("Ola Sr.", nome)
	sistema := 18.1
	fmt.Println(sistema, " versão atual")
	fmt.Println("")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("")
	fmt.Println("O comando escolhido foi ", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("")
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibi Logs")
	fmt.Println("0 - sair")
}

func exibirEscolha() {
	fmt.Println("")
	fmt.Println("Deseja escolher outra opção?")
	fmt.Println("1 - Sim")
	fmt.Println("2 - Não")
}

func exibeMenuNovamente() {
	fmt.Println("")
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibi Logs")
	fmt.Println("0 - sair")
}

func leComandoNovamente() int {
	var ccomando int
	fmt.Println("")
	fmt.Println("")
	fmt.Scan(&ccomando)
	fmt.Println("")
	fmt.Println("Você selecionou ", ccomando)
	return ccomando
}

func saindo() {
	fmt.Println("")
	fmt.Println("Saindo ...")
	os.Exit(0)
}

func comandoNãoIndentificado() {
	fmt.Println("")
	fmt.Println("Comando não identificado!")
	os.Exit(-1)
}
