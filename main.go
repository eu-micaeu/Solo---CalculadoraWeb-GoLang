package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define o tratamento da rota raiz
    http.HandleFunc("/", handler)

    // Define o tratamento da pasta "static" para servir arquivos estáticos
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Inicia o servidor e imprime uma mensagem informando que está ouvindo na porta 8080
    fmt.Println("Server is listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}

// Função que trata as requisições para a rota raiz
func handler(w http.ResponseWriter, r *http.Request) {
    // Serve o arquivo "index.html" ao usuário
    http.ServeFile(w, r, "index.html")
}
