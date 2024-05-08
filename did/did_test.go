package did

import (
	"testing"

	mcodec "github.com/multiformats/go-multicodec"
)

func TestParseDIDKey(t *testing.T) {
	str := "did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z"
	d, err := Parse(str)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d.String() != str {
		t.Fatalf("expected %v to equal %v", d.String(), str)
	}
}

func TestDecodeDIDKeyED25519(t *testing.T) {
	str := "did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z"
	d0, err := Parse(str)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d0.Algorithm() != mcodec.Ed25519Pub {
		t.Fatalf("expected algorithm to be %s", mcodec.Ed25519Pub)
	}
	d1, err := Decode(d0.Bytes())
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d1.String() != str {
		t.Fatalf("expected %v to equal %v", d1.String(), str)
	}
}

func TestDecodeDIDKeySecp256k1(t *testing.T) {
	str := "did:key:zQ3shq4bgfyqUGzQiXXneg4xtQBh4t8vmb8bREHveJVqj2DGW"
	d0, err := Parse(str)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d0.Algorithm() != mcodec.Secp256k1Pub {
		t.Fatalf("expected algorithm to be %s", mcodec.Secp256k1Pub)
	}
	d1, err := Decode(d0.Bytes())
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d1.String() != str {
		t.Fatalf("expected %v to equal %v", d1.String(), str)
	}
}

func TestParseDIDWeb(t *testing.T) {
	str := "did:web:up.web3.storage"
	d, err := Parse(str)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d.String() != str {
		t.Fatalf("expected %v to equal %v", d.String(), str)
	}
}

func TestDecodeDIDWeb(t *testing.T) {
	str := "did:web:up.web3.storage"
	d0, err := Parse(str)
	if err != nil {
		t.Fatalf("%v", err)
	}
	d1, err := Decode(d0.Bytes())
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d1.String() != str {
		t.Fatalf("expected %v to equal %v", d1.String(), str)
	}
}

func TestEquivalence(t *testing.T) {
	u0 := DID{}
	u1 := Undef
	if u0 != u1 {
		t.Fatalf("undef DID not equivalent")
	}

	d0, err := Parse("did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z")
	if err != nil {
		t.Fatalf("%v", err)
	}

	d1, err := Parse("did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z")
	if err != nil {
		t.Fatalf("%v", err)
	}

	if d0 != d1 {
		t.Fatalf("two equivalent DID not equivalent")
	}
}
