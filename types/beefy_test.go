package types_test

import (
	"testing"

	. "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	. "github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	. "github.com/centrifuge/go-substrate-rpc-client/v4/types/test_utils"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestBeefySignature(t *testing.T) {
	empty := types.NewOptionBeefySignatureEmpty()
	assert.True(t, empty.IsNone())
	assert.False(t, empty.IsSome())

	sig := types.NewOptionBeefySignature(types.BeefySignature{})
	sig.SetNone()
	assert.True(t, sig.IsNone())
	sig.SetSome(types.BeefySignature{})
	assert.True(t, sig.IsSome())
	ok, _ := sig.Unwrap()
	assert.True(t, ok)
	AssertRoundtrip(t, sig)
}

func makeCommitment() (*Commitment, error) {
	data, err := Encode([]byte("Hello World!"))
	if err != nil {
		return nil, err
	}

	payloadItem := PayloadItem{
		ID:   [2]byte{'m', 'h'},
		Data: data,
	}

	commitment := Commitment{
		Payload:        []PayloadItem{payloadItem},
		BlockNumber:    5,
		ValidatorSetID: 0,
	}

	return &commitment, nil
}

func makeLargeCommitment() (*Commitment, error) {
	data := MustHexDecodeString("0xb5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c")

	payloadItem := PayloadItem{
		ID:   [2]byte{'m', 'h'},
		Data: data,
	}

	commitment := Commitment{
		Payload:        []PayloadItem{payloadItem},
		BlockNumber:    5,
		ValidatorSetID: 3,
	}

	return &commitment, nil
}

func TestCommitment_Encode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)
	AssertEncode(t, []EncodingAssert{
		{c, MustHexDecodeString("0x046d68343048656c6c6f20576f726c6421050000000000000000000000")},
	})
}

func TestLargeCommitment_Encode(t *testing.T) {
	c, err := makeLargeCommitment()
	assert.NoError(t, err)
	fmt.Println(len(c.Payload[0].Data))
	fmt.Println(EncodeToHex(c))
}

func TestCommitment_Decode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)

	AssertDecode(t, []DecodingAssert{
		{
			Input:    MustHexDecodeString("0x046d68343048656c6c6f20576f726c6421050000000000000000000000"),
			Expected: *c,
		},
	})
}

func TestCommitment_EncodeDecode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)

	AssertRoundtrip(t, *c)
}

func TestSignedCommitment_Decode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)

	s := SignedCommitment{
		Commitment: *c,
		Signatures: []OptionBeefySignature{
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignature(sig1),
			NewOptionBeefySignature(sig2),
		},
	}

	AssertDecode(t, []DecodingAssert{
		{
			Input:    MustHexDecodeString("0x046d68343048656c6c6f20576f726c642105000000000000000000000004300400000008558455ad81279df0795cc985580e4fb75d72d948d1107b2ac80a09abed4da8480c746cc321f2319a5e99a830e314d10dd3cd68ce3dc0c33c86e99bcb7816f9ba012d6e1f8105c337a86cdd9aaacdc496577f3db8c55ef9e6fd48f2c5c05a2274707491635d8ba3df64f324575b7b2a34487bca2324b6a0046395a71681be3d0c2a00"),
			Expected: s,
		},
	})
}

func TestSignedCommitment_EncodeDecode(t *testing.T) {
	c, err := makeCommitment()
	assert.NoError(t, err)

	s := SignedCommitment{
		Commitment: *c,
		Signatures: []OptionBeefySignature{
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignature(sig1),
			NewOptionBeefySignature(sig1),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignatureEmpty(),
			NewOptionBeefySignature(sig1),
		},
	}

	AssertRoundtrip(t, s)
}

func TestBeefySignature_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[BeefySignature](t, 100)
	AssertDecodeNilData[BeefySignature](t)
	AssertEncodeEmptyObj[BeefySignature](t, 65)
}

var (
	optionBeefySignatureFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(o *OptionBeefySignature, c fuzz.Continue) {
			if c.RandBool() {
				*o = NewOptionBeefySignatureEmpty()
				return
			}

			var b BeefySignature
			c.Fuzz(&b)

			*o = NewOptionBeefySignature(b)
		}),
	}
)

func TestOptionBeefySignature_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[OptionBeefySignature](t, 100, optionBeefySignatureFuzzOpts...)
	AssertEncodeEmptyObj[OptionBeefySignature](t, 1)
}
