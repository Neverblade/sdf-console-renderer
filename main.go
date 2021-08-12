package main

import "sdf-console-renderer/internal/visualizer"

func main() {
	visualizer := visualizer.NewConsoleVisualizer()
	for i := 0; i < 10; i++ {
		visualizer.Visualize(nil)
	}
}
