package padreader

import (
	"testing"

	"github.com/filecoin-project/specs-actors/actors/abi"

	"github.com/stretchr/testify/assert"
)

func TestComputePaddedSize(t *testing.T) {
	assert.Equal(t, abi.UnpaddedPieceSize(1040384), PaddedSize(1000000))

	assert.Equal(t, abi.UnpaddedPieceSize(1016), PaddedSize(548))
	assert.Equal(t, abi.UnpaddedPieceSize(1016), PaddedSize(1015))
	assert.Equal(t, abi.UnpaddedPieceSize(1016), PaddedSize(1016))
	assert.Equal(t, abi.UnpaddedPieceSize(2032), PaddedSize(1017))

	assert.Equal(t, abi.UnpaddedPieceSize(2032), PaddedSize(1024))
	assert.Equal(t, abi.UnpaddedPieceSize(4064), PaddedSize(2048))
}
