package logstream

type ProcessorBuilder struct {
	processor Processor
}

func (pb *ProcessorBuilder) SetProcessor(processor Processor) *ProcessorBuilder {
	pb.processor = processor

	return pb
}

func (pb *ProcessorBuilder) Build() Processor {
	return pb.processor
}

func NewProcessorBuilder() *ProcessorBuilder {
	return &ProcessorBuilder{}
}
