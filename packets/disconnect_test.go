package packets

import (
	"bytes"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
)

func TestDisconnectEncode(t *testing.T) {
	require.Contains(t, expectedPackets, Disconnect)
	for i, wanted := range expectedPackets[Disconnect] {
		require.Equal(t, uint8(14), Disconnect, "Incorrect Packet Type [i:%d]", i)

		pk := new(DisconnectPacket)
		copier.Copy(pk, wanted.packet.(*DisconnectPacket))

		require.Equal(t, Disconnect, pk.Type, "Mismatched Packet Type [i:%d]", i)
		require.Equal(t, Disconnect, pk.FixedHeader.Type, "Mismatched FixedHeader Type [i:%d]", i)

		var b bytes.Buffer
		err := pk.Encode(&b)

		require.NoError(t, err, "Error writing buffer [i:%d]", i)
		require.Equal(t, len(wanted.rawBytes), len(b.Bytes()), "Mismatched packet length [i:%d]", i)
		require.EqualValues(t, wanted.rawBytes, b.Bytes(), "Mismatched byte values [i:%d]", i)
	}
}

func TestDisconnectDecode(t *testing.T) {
	pk := newPacket(Disconnect).(*DisconnectPacket)

	var b = []byte{}
	err := pk.Decode(b)
	require.NoError(t, err, "Error unpacking buffer")
	require.Empty(t, b)
}

func BenchmarkDisconnectDecode(b *testing.B) {
	pk := newPacket(Disconnect).(*DisconnectPacket)
	pk.FixedHeader.decode(expectedPackets[Disconnect][0].rawBytes[0])

	for n := 0; n < b.N; n++ {
		pk.Decode(expectedPackets[Disconnect][0].rawBytes[2:])
	}
}
