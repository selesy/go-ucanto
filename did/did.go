package did

import (
	"fmt"
	"strings"

	mbase "github.com/multiformats/go-multibase"
	mcodec "github.com/multiformats/go-multicodec"
	varint "github.com/multiformats/go-varint"
)

const Prefix = "did:"
const KeyPrefix = "did:key:"

var MethodOffset = varint.UvarintSize(uint64(mcodec.Multidid))

type DID struct {
	key  bool
	code mcodec.Code
	str  string
}

// Undef can be used to represent a nil or undefined DID, using DID{}
// directly is also acceptable.
var Undef = DID{}

func (d DID) Defined() bool {
	return d.str != ""
}

func (d DID) Bytes() []byte {
	if !d.Defined() {
		return nil
	}

	bytes := []byte(d.str)
	if d.code != mcodec.Identity {
		bytes = append(varint.ToUvarint(uint64(d.code)), bytes...)
	}

	return bytes
}

func (d DID) Key() []byte {
	return []byte(d.str)
}

func (d DID) Algorithm() mcodec.Code {
	return d.code
}

func (d DID) DID() DID {
	return d
}

// String formats the decentralized identity document (DID) as a string.
func (d DID) String() string {
	if d.key {
		key, _ := mbase.Encode(mbase.Base58BTC, d.Bytes())

		return KeyPrefix + key
	}

	return Prefix + d.str[MethodOffset:]
}

func Decode(bytes []byte) (DID, error) {
	code, read, err := varint.FromUvarint(bytes)
	if err != nil {
		return Undef, err
	}

	mcode := mcodec.Code(code)
	if mcode == mcodec.Ed25519Pub || mcode == mcodec.Secp256k1Pub {
		return DID{str: string(bytes[read:]), key: true, code: mcode}, nil
	} else if mcode == mcodec.Multidid {
		return DID{str: string(bytes), code: mcode}, nil
	}

	return Undef, fmt.Errorf("unsupported DID encoding: 0x%x", code)
}

func Parse(str string) (DID, error) {
	if !strings.HasPrefix(str, Prefix) {
		return Undef, fmt.Errorf("must start with 'did:'")
	}

	if strings.HasPrefix(str, KeyPrefix) {
		code, bytes, err := mbase.Decode(str[len(KeyPrefix):])
		if err != nil {
			return Undef, err
		}
		if code != mbase.Base58BTC {
			return Undef, fmt.Errorf("not Base58BTC encoded")
		}
		return Decode(bytes)
	}

	buf := make([]byte, MethodOffset)
	varint.PutUvarint(buf, uint64(mcodec.Multidid))
	suffix, _ := strings.CutPrefix(str, Prefix)
	buf = append(buf, suffix...)

	return DID{str: string(buf)}, nil
}
