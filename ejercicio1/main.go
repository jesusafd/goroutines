package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex

/*--------------------------------------------Estructuras de datos para los personajes--------------------------------------------*/
type jugador struct {
	salud int
}

type enemigo struct {
	salud int
}

/*--------------------------------------------instanciamos a los personajes--------------------------------------------*/
// Los personajes se declararon de forma global para poder practicar el uso de
// mutex, el cual nos permite ejecutar de forma mas organizada las goroutines
// en este caso se utlizo para el daño generado
var j1 jugador = jugador{
	salud: 100,
}
var e1 enemigo = enemigo{
	salud: 100,
}

/*--------------------------------------------Ataques de los personajes--------------------------------------------*/
func ataqueEnemigo(victoria chan string) {
	for j1.salud > 0 && e1.salud > 0 {
		// Ejecutamoas un ataque a la vez evitando que al final ambos personajes ganen
		mutex.Lock()
		ataque := rand.Intn(20)
		j1.salud -= ataque
		if j1.salud < 0 {
			j1.salud = 0
		}
		fmt.Println("Jugador : -", ataque, "total vida:", j1.salud)
		mutex.Unlock()
		time.Sleep(time.Second)
	}
	// En caso de los puntos de salud del otro personaje sean cero quiere decir que gano el jugador
	// Usamos un canal para escribir la victoria o derrota del personaje
	if e1.salud > 0 {
		victoria <- "Gano"
	} else {
		victoria <- "Perdio"
	}
}

func ataqueJugador(victoria chan string) {
	for j1.salud > 0 && e1.salud > 0 {
		// Ejecutamoas un ataque a la vez evitando que al final ambos personajes ganen
		mutex.Lock()
		ataque := rand.Intn(10)
		e1.salud -= ataque
		if e1.salud < 0 {
			e1.salud = 0
		}
		fmt.Println("Enemigo : -", ataque, "total vida:", e1.salud)
		mutex.Unlock()
		time.Sleep(time.Second)
	}
	// En caso de los puntos de salud del otro personaje sean cero quiere decir que gano el enemigo
	// Usamos un canal para escribir la victoria o derrota del personaje
	if j1.salud > 0 {
		victoria <- "Gano"
	} else {
		victoria <- "Perdio"
	}
}

func main() {
	// Creamos los canales de cada jugador
	victoriaJugador := make(chan string)
	victoriaEnemigo := make(chan string)
	go ataqueJugador(victoriaJugador)
	go ataqueEnemigo(victoriaEnemigo)
	// Mediante un select leemos los canales de ambos personjaes
	for i := 0; i < 2; i++ {
		select {
		case vicJ := <-victoriaJugador:
			fmt.Println("El jugador", vicJ)
		case vicE := <-victoriaEnemigo:
			fmt.Println("El enemigo", vicE)
		}
	}
}
