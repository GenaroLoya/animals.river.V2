package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Identifier string

const (
	Wolf   Identifier = "🐺"
	Goat   Identifier = "🐐"
	Carrot Identifier = "🥕"
	Cowboy Identifier = "🤠"
)

type Position string

const (
	Left  Position = "🠸"
	Right Position = "🠺"
)

type Entity struct {
	Identifier Identifier
	Position   Position
}

var entities = []Entity{
	{Goat, Left},
	{Wolf, Left},
	{Carrot, Left},
	{Cowboy, Left},
}

func printState(entities []Entity) {
	for _, entity := range entities {
		fmt.Printf("%s: %s  ", entity.Identifier, entity.Position)
	}
	fmt.Println()
}

var entitiesAntiStates = [][]Entity{{
	{Goat, Right},
	{Wolf, Right},
	{Carrot, Left},
	{Cowboy, Left},
}, {
	{Goat, Right},
	{Wolf, Left},
	{Carrot, Right},
	{Cowboy, Left},
}, {
	{Goat, Right},
	{Wolf, Right},
	{Carrot, Right},
	{Cowboy, Left},
}, {
	{Goat, Left},
	{Wolf, Left},
	{Carrot, Left},
	{Cowboy, Right},
}, {
	{Goat, Left},
	{Wolf, Right},
	{Carrot, Left},
	{Cowboy, Right},
}, {
	{Goat, Left},
	{Wolf, Left},
	{Carrot, Right},
	{Cowboy, Right},
}}

func moveEntity(entity *Entity) {
	if entity.Position == Right {
		entity.Position = Left
	} else {
		entity.Position = Right
	}
}

func isValidState(state []Entity) bool {
	for _, antiState := range entitiesAntiStates {
		if statesEqual(state, antiState) {
			return false
		}
	}
	return true
}

func statesEqual(state1, state2 []Entity) bool {
	if len(state1) != len(state2) {
		return false
	}
	for i := range state1 {
		if state1[i] != state2[i] {
			return false
		}
	}
	return true
}

func generateRandomValidMoves(state []Entity) ([]Entity, bool) {
	stack := make([][]Entity, 0) // Pila para almacenar estados
	stack = append(stack, state)

	for len(stack) > 0 {
		currentState := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Intentar realizar un movimiento aleatorio
		entityIndex := rand.Intn(len(currentState))
		moveEntity(&currentState[entityIndex])

		fmt.Println("Current state:")
		printState(currentState)

		// Verificar si el nuevo estado es válido
		if isValidState(currentState) {
			// Almacenar el nuevo estado en la pila
			stack = append(stack, currentState)
		} else {
			// Deshacer el movimiento no válido
			moveEntity(&currentState[entityIndex])
			fmt.Println("Invalid state, undoing move:")
			printState(currentState)
		}

		// Verificar si se ha encontrado un estado válido
		if isFinalState(currentState) {
			return currentState, true
		}
	}

	return nil, false
}

func isFinalState(state []Entity) bool {
	// Implementa aquí una condición que determine si se ha alcanzado el estado final deseado
	// Por ejemplo, si todos los elementos están en el lado derecho
	for _, entity := range state {
		if entity.Position != Right {
			return false
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Estado inicial
	initialState := make([]Entity, len(entities))
	copy(initialState, entities)

	fmt.Println("Estado inicial:")
	printState(entities)

	// Generar movimientos aleatorios y encontrar una configuración válida
	for true {
		_, validState := generateRandomValidMoves(initialState)
		if validState {
			fmt.Println("Se encontró un estado válido:")
			printState(initialState)
			break
		} else {
			fmt.Println("No se encontró un estado válido.")
			printState(initialState)
		}
	}
}
