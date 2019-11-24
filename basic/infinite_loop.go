package "main"

import "fmt"

func main() {
    sum := 0
    for {
        sum++ // se repite interminablemente
    }
    // nunca se imprime
    fmt.Println(sum)
}

