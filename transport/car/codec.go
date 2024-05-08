package car

import (
	"github.com/selesy/go-ucanto/core/message"
	"github.com/selesy/go-ucanto/transport"
	"github.com/selesy/go-ucanto/transport/car/request"
	"github.com/selesy/go-ucanto/transport/car/response"
)

type carOutbound struct{}

func (oc *carOutbound) Encode(msg message.AgentMessage) (transport.HTTPRequest, error) {
	return request.Encode(msg)
}

func (oc *carOutbound) Decode(res transport.HTTPResponse) (message.AgentMessage, error) {
	return response.Decode(res)
}

func NewCAROutboundCodec() transport.OutboundCodec {
	return &carOutbound{}
}
