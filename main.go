package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func cepHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")


	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		http.Error(w, "Erro ao consultar o ViaCEP", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "CEP inválido ou não encontrado", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erro ao ler a resposta do ViaCEP", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func main() {
	http.HandleFunc("/cep", cepHandler) 

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
