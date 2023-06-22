package logstream

type ResponseBuilder struct {
	response      *Response
	outputChannel chan *Response
}

func (rb *ResponseBuilder) SetMessageType(messageType MessageType) *ResponseBuilder {
	rb.response.messageType = messageType

	return rb
}

func (rb *ResponseBuilder) SetOutputChannel(outputChannel chan *Response) *ResponseBuilder {
	rb.outputChannel = outputChannel

	return rb
}

func (rb *ResponseBuilder) StreamResponse(r *Response) {
	rb.outputChannel <- r
}

func (rb *ResponseBuilder) SetMessage(message string) *ResponseBuilder {
	rb.response.message = message

	return rb
}

func (rb *ResponseBuilder) SetTargets(targets []int) *ResponseBuilder {
	rb.response.targets = append(rb.response.targets, targets...)

	return rb
}

func (rb *ResponseBuilder) GetTargets() []int {
	return rb.response.targets
}

func (rb *ResponseBuilder) Build() *Response {
	return rb.response
}

func (rb *ResponseBuilder) Clear() *ResponseBuilder {
	rb.response = &Response{}

	return rb
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		response: &Response{},
	}
}
