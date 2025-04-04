// Projeto: Produtor x Consumidor com Buffer Limitado
// Equipe GoSync
// Integrantes:
// - Ana Beatriz Maciel Nunes
// - Fernando Luiz Da Silva Freire
// - Juliana Ballin Lima
// - Lucas Carvalho Dos Santos
// - Marcelo Heitor De Almeida Lira

package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// Função que simula um produtor
func produtor(id int, buffer chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        item := rand.Intn(100)
        buffer <- item
        fmt.Printf("Produtor %d produziu item %d\n", id, item)
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
    }
}

// Função que simula um consumidor
func consumidor(id int, buffer chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 4; i++ {
        item := <-buffer
        fmt.Printf("Consumidor %d consumiu item %d\n", id, item)
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    // Configurações iniciais
    bufferSize := 3
    numProdutores := 3
    numConsumidores := 2

    buffer := make(chan int, bufferSize)
    var wg sync.WaitGroup

    // Inicia produtores
    for i := 1; i <= numProdutores; i++ {
        wg.Add(1)
        go produtor(i, buffer, &wg)
    }

    // Inicia consumidores
    for i := 1; i <= numConsumidores; i++ {
        wg.Add(1)
        go consumidor(i, buffer, &wg)
    }

    // Aguarda conclusão
    wg.Wait()
    fmt.Println("Execução finalizada.")
}
