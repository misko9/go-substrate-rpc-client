// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signature_test

import (
	"testing"

	. "github.com/misko9/go-substrate-rpc-client/v4/signature"
	"github.com/misko9/go-substrate-rpc-client/v4/types"
	"github.com/misko9/go-substrate-rpc-client/v4/types/codec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testSecretPhrase = "little orbit comfort eyebrow talk pink flame ridge bring milk equip blood"
var testSecretSeed = "0x167d9a020688544ea246b056799d6a771e97c9da057e4d0b87024537f99177bc"
var testPubKey = "0xdc64bef918ddda3126a39a11113767741ddfdf91399f055e1d963f2ae1ec2535"
var testAddressSS58 = "5H3gKVQU7DfNFfNGkgTrD7p715jjg7QXtat8X3UxiSyw7APW"
var testKusamaAddressSS58 = "HZHyokLjagJ1KBiXPGu75B79g1yUnDiLxisuhkvCFCRrWBk"
var testPolkadotAddressSS58 = "15yyTpfXxzvqhCNniKWrMGeFrhjPNQxfy5ccgLUKGY1THbTW"

func TestKeyRingPairFromSecretPhrase_SubstrateAddress(t *testing.T) {
	p, err := KeyringPairFromSecret(testSecretPhrase, 42)
	assert.NoError(t, err)

	assert.Equal(t, KeyringPair{
		URI:       testSecretPhrase,
		Address:   testAddressSS58,
		PublicKey: codec.MustHexDecodeString(testPubKey),
	}, p)
}

func TestKeyRingPairFromSecretPhrase_PolkadotAddress(t *testing.T) {
	p, err := KeyringPairFromSecret(testSecretPhrase, 0)
	assert.NoError(t, err)

	assert.Equal(t, KeyringPair{
		URI:       testSecretPhrase,
		Address:   testPolkadotAddressSS58,
		PublicKey: codec.MustHexDecodeString(testPubKey),
	}, p)
}

func TestKeyRingPairFromSecretPhrase_KusamaAddress(t *testing.T) {
	p, err := KeyringPairFromSecret(testSecretPhrase, 2)
	assert.NoError(t, err)

	assert.Equal(t, KeyringPair{
		URI:       testSecretPhrase,
		Address:   testKusamaAddressSS58,
		PublicKey: codec.MustHexDecodeString(testPubKey),
	}, p)
}

func TestKeyRingPairFromSecretPhrase_InvalidSecretPhrase(t *testing.T) {
	_, err := KeyringPairFromSecret("foo", 42)
	assert.Error(t, err)
}

func TestKeyringPairFromSecretSeed(t *testing.T) {
	p, err := KeyringPairFromSecret(testSecretSeed, 42)
	assert.NoError(t, err)

	assert.Equal(t, KeyringPair{
		URI:       testSecretSeed,
		Address:   testAddressSS58,
		PublicKey: codec.MustHexDecodeString(testPubKey),
	}, p)
}

func TestKeyringPairFromSecretSeedAndNetwork(t *testing.T) {
	p, err := KeyringPairFromSecret(testSecretSeed, 42)
	assert.NoError(t, err)

	assert.Equal(t, KeyringPair{
		URI:       testSecretSeed,
		Address:   testAddressSS58,
		PublicKey: codec.MustHexDecodeString(testPubKey),
	}, p)
}

func TestSignAndVerify(t *testing.T) {
	data := []byte("hello!")

	sig, err := Sign(data, TestKeyringPairAlice.URI)
	assert.NoError(t, err)

	ok, err := Verify(data, sig, TestKeyringPairAlice.URI)
	assert.NoError(t, err)

	assert.True(t, ok)
}

func TestDecodeSubstrateMessage(t *testing.T) {
	hex := "0x046d68809de63e597f2ff84b4e287c829b93a3aad304b1aebd51ec2d7392d41047e951c723040000690000000000000004d80500000010d260924684152d72c9d81d028382b4847f895ac6081cdcaa93c37e00df3b275b7fbc96c10c6fc86743d3aaba0bf59278822b1610218858edfbaf77a9884ee031004f733874acfbdd9d3e45e9b1247647d02761d1673c303a7982cd6eb5d573eb3036415493a0d655937c64a9536b3113747ec972022b7250c89c30888a45a3908c013251044285c3000d54d0fc0572a0559473e7e223175daa44046b3523d4c4ef7a71f6a4287cc7030d8fd30f504ea0360e0a4c3047e33f8f5d4e4c11f867cbac700047e643aee47dacc161c8cb049597f30a0c74ed596cd78b94f21526c1d744835758ad4492002c5890a0f45990aa175a320a676ce89977cffbe2eb73b2cf16d62900"
	compactCommitment := types.CompactSignedCommitment{}

	// attempt to decode the SignedCommitments
	err := codec.DecodeFromHex(hex, &compactCommitment)
	require.NoError(t, err)

	signedCommitment := compactCommitment.Unpack()

	t.Logf("COMMITMENT %#v \n", signedCommitment.Commitment)

	for _, v := range signedCommitment.Signatures {
		if v.HasValue {
			_, sig := v.Unwrap()
			t.Logf("SIGNATURE %#v \n", codec.HexEncodeToString(sig[:]))

		} else {
			t.Logf("NONE\n")

		}
	}
}
