package visualizer

import (
	"testing"
)

func TestConsoleVisualizer(t *testing.T) {
	visualizer := NewConsoleVisualizer()
	for i := 0; i < 10; i++ {
		visualizer.Visualize(nil)
	}
}
