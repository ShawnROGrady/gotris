package canvas

// WithBackground specifies the background for the canvas
func WithBackground(c Color) Option {
	return withBackground(c)
}

// Option represents a configuration option for the canvas
type Option interface {
	ApplyToCanvas(c *TermCanvas)
}

type withBackground Color

func (w withBackground) ApplyToCanvas(c *TermCanvas) {
	c.Background = Color(w)
}

// WithDebugMode specifies that the canvas should render in debug mode
func WithDebugMode() Option {
	return withDebugMode{}
}

type withDebugMode struct{}

func (w withDebugMode) ApplyToCanvas(c *TermCanvas) {
	c.debugMode = true
}

// WithWidth returns an option the specifies the width of the board
func WithWidth(width int) Option {
	return withWidth(width)
}

type withWidth int

func (w withWidth) ApplyToCanvas(c *TermCanvas) {
	c.width = int(w)
}

// WithHeight returns an option the specifies the height of the board
func WithHeight(height int) Option {
	return withHeight(height)
}

type withHeight int

func (w withHeight) ApplyToCanvas(c *TermCanvas) {
	c.height = int(w)
}
