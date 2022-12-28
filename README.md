# Ejercicios para practica de goroutine

## Ejercicio 1: juego con goroutines
Realizar un juego en el cual existan dos persojes
* Jugador
* Enemigo


### Caracteristicas del ejercicio 1
Diseñar un jeugo con las siguientes caracteristicas, empleando goroutines:


* Los personajes se golpearan uno a otro
* Los personajes no tomaran turnos para golpearse, si no que una vez cumplido el tiempo de ataque el personaje atacara
* el tiempo de ataque sera de 1 segundo para cada juego
* El daño causado sera aleatorio en cada turno, el atque maximo no debe exceder 20 puntos de vida
* Cada que un personaje haga un ataque se debe imprimir el daño realizado y la vida total de personaje atacado
* Se debe imprimir que personaje gano y cual perdio
* Ambos personajes tendran un nivel de salud o puntos de vida
* Perdera el personaje que sus puntos de vida sean igual a 0 en el ultimo acaque


#### Finalidad el ejercicio 
Familizarese con el uso de goroutines y poner en practica conceptos basicos de estas como lo son:
* canales
* select
* mutex

## Ejercicio 2: adivinar el numero
Realizar un juego en el cual existan dos jugadores (goroutines), donde cada uno genrar un numero aleatorio y uno
tratara de adivnar el numero del otro jugador


### Caracteristicas del juego
* Cada jugador generara un nuemro, en este caso solo un adiviara el numero
* Cuando el jugadro adivine el numero ganara
* El numero aleatorio debera ser generado aleatoriamente
* Debe hacer uso de conceptos como range, mutex, canales y select

#### Finalidad el ejercicio 
Familizarese con el uso de goroutines y poner en practica conceptos basicos de estas como lo son:
* canales
* select
* mutex
* range
