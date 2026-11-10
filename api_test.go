// SPDX-FileCopyrightText: the original author or authors
// SPDX-License-Identifier: Apache-2.0

package pbf_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	parser "m4o.io/pbf/v2"
	"m4o.io/pbf/v2/internal/encoder"
)

func TestBlobCompressionValuesAreExported(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		compression parser.BlobCompression
		want        encoder.BlobCompression
	}{
		{"RAW", parser.RAW, encoder.RAW},
		{"ZLIB", parser.ZLIB, encoder.ZLIB},
		{"LZMA", parser.LZMA, encoder.LZMA},
		{"LZ4", parser.LZ4, encoder.LZ4},
		{"ZSTD", parser.ZSTD, encoder.ZSTD},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, tt.compression)
		})
	}
}
