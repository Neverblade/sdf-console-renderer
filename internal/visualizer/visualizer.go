package visualizer

type Visualizer interface {
	// Visualize takes input pixel data and converts it into visual form.
	Visualize(pixels [][]uint8)
}
