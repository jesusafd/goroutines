package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var mutex sync.Mutex

// Para este ejercico haremos uso de una varible global llamada juego
// la cual sera verdadero mientra nadie adivine el numero y falso cuando
// alguien lo adivine, indicara el estado del juego
// Se hara uso de mutex para la modificacion de este dato
var juego bool

// El jugador generar un numero aleatorio y lo enviara pro el canal mientra
// el juego siga, cuado se adivien el numero se cerrar el canal
func jugador(canal chan int) {
	var numero int
	for juego {
		numero = rand.Intn(5)
		canal <- numero
	}
	close(canal)
}

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	// Iniciamos el juego colocando el estado de juego en true y ejcutando
	// las goruotines
	juego = true
	go jugador(canal1) // Jugador1 que adivinara el numero
	go jugador(canal2) // Jugador2 que generar un numero

	var numero2 int
	// Se usa range para leer los canale mietras sigan abiertos
	for numero1 := range canal1 {
		numero2 = <-canal2
		// El jugador 1 intenta adivinar el numero del jugador 2
		if numero1 == numero2 {
			// En caso de que se adivine el numero se cambiara el estado del juego
			// para lo cual haremos uso de mutex
			fmt.Println("El jugador adivino el numero", numero1, "==", numero2)
			mutex.Lock()
			juego = false
			mutex.Unlock()
			break
		} else {
			fmt.Println(numero1, "!=", numero2)
		}
	}
}
