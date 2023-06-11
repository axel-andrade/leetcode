package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func runProblem(problemNum int) {
	// Mapeie os números dos problemas para os diretórios correspondentes
	problems := map[int]string{
		20:  "20_valid_parentheses",
		28:  "28_implement_strstr",
		58:  "58_length_of_last_word",
		66:  "66_plus_one",
		125: "125_valid_palindrome",
		136: "136_single_number",
		217: "217_contains_duplicated",
		// Adicione mais mapeamentos aqui conforme necessário
	}

	// Verifique se o número do problema está mapeado
	problemDir, found := problems[problemNum]
	if !found {
		fmt.Println("Problema inválido")
		return
	}

	// Obtenha o caminho absoluto para o diretório do problema
	problemPath := filepath.Join(problemDir)

	// Execute o comando "go run" para o arquivo main.go no diretório do problema
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Defina o diretório de trabalho para o diretório do problema
	cmd.Dir = problemPath

	fmt.Printf("Executando o problema %s...\n", problemDir)

	// Execute o comando
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erro ao executar o problema %s: %v\n", problemDir, err)
	}

	fmt.Printf("Concluído problema %s.\n", problemDir)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run main.go <problema>")
		return
	}

	problemNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Problema inválido")
		return
	}

	runProblem(problemNum)
}
