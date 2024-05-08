package response

import (
	"fmt"

	"github.com/selesy/go-ucanto/core/car"
	"github.com/selesy/go-ucanto/core/dag/blockstore"
	"github.com/selesy/go-ucanto/core/message"
	"github.com/selesy/go-ucanto/transport"
)

const ContentType = car.ContentType

func Decode(response transport.HTTPResponse) (message.AgentMessage, error) {
	roots, blocks, err := car.Decode(response.Body())
	if err != nil {
		return nil, fmt.Errorf("decoding response: %s", err)
	}
	bstore, err := blockstore.NewBlockReader(blockstore.WithBlocksIterator(blocks))
	if err != nil {
		return nil, fmt.Errorf("creating blockstore: %s", err)
	}
	return message.NewMessage(roots, bstore)
}
