package principal

import (
	"github.com/selesy/go-ucanto/ucan"
)

type Signer interface {
	ucan.Signer
	Code() uint64
	Verifier() Verifier
	Encode() []byte
}

type Verifier interface {
	ucan.Verifier
	Code() uint64
	Encode() []byte
}
