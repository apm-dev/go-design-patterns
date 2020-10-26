package shapes

import (
	"fmt"
	"github.com/apm-dev/go-design-patterns/behavioral/strategy"
	"os"
)

const (
	TextStrategy  = "text"
	ImageStrategy = "image"
)

func NewPrinter(s string) (strategy.Output, error) {
	switch s {
	case TextStrategy:
		return &TextSquare{
			strategy.PrintOutput{LogWriter: os.Stdout},
		}, nil
	case ImageStrategy:
		return &ImageSquare{
			strategy.PrintOutput{LogWriter: os.Stdout},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
