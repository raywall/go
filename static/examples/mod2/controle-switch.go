package main

import "fmt"

func main() {
    dia := 3
    switch dia {
        case 1, 2:
            fmt.Println("Início da semana")
        case 3, 4, 5:
            fmt.Println("Meio da semana")
        default:
            fmt.Println("Fim de semana")
    }

    // Switch com expressão
    switch x := dia \* 2; x {
    case 6:
        fmt.Println("Dia 3 dobrado")
    default:
        fmt.Println("Outro valor")
    }
}