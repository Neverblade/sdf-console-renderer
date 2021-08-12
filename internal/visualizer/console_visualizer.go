package visualizer

import (
	"fmt"
)

func NewConsoleVisualizer() Visualizer {
	return &consoleVisualizer{}
}

type consoleVisualizer struct{}

func (c *consoleVisualizer) Visualize(pixels [][]uint8) {
	ClearConsole()
	fmt.Println("Hello World!")
}
