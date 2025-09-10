package logstream

type Processor interface {
	Process([]string) []string
}

type InputProcessor struct{}

func (p *InputProcessor) ReceiveInputFromArray(instr []string) {
	for _, input := range instr {
		p.Process(input)
	}
}

func (p *InputProcessor) Process(_ string) {
	panic("implement me!")
}
