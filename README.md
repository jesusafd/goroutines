# Ejercicio para practica de goroutine
Realizar un juego en el cual existan dos persojes
* Jugador
* Enemigo


### Caracteristicas del juego
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