package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bndw/port-scanner/pkg/nmap"
)

func TestEncodeDecodeScanResults(t *testing.T) {
	results := []nmap.PortStatus{
		{
			Port:     uint16(22),
			Protocol: "tcp",
			State:    "open",
			Name:     "ssh",
		},
	}

	_, err := encodeScanResults(results)
	assert.Nil(t, err)
}

func TestDecodeCreatedAt(t *testing.T) {
	const createdAt = "2021-05-29T18:54:23Z"
	_, err := decodeCreatedAt(createdAt)
	assert.Nil(t, err)
}
