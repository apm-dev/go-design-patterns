package strategy

import "io"

type Output interface {
	Print() error
	SetLogger(io.Writer)
	SetWriter(io.Writer)
}

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (p *PrintOutput) SetLogger(w io.Writer) {
	p.LogWriter = w
}

func (p *PrintOutput) SetWriter(w io.Writer) {
	p.Writer = w
}
