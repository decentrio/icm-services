// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package beacon

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// newBeaconServer serves every request with the given status, consensus
// version header and body.
func newBeaconServer(t *testing.T, status int, fork string, body []byte) *BeaconClient {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "application/octet-stream", r.Header.Get("Accept"))
		if fork != "" {
			w.Header().Set("Eth-Consensus-Version", fork)
		}
		w.WriteHeader(status)
		_, _ = w.Write(body)
	}))
	t.Cleanup(server.Close)
	return NewBeaconClient(server.URL)
}

// Blocks and states of the supported forks must come back with the fork the
// node reported, so callers can pick the matching decoder.
func TestBeaconClientReturnsFork(t *testing.T) {
	body := []byte{0x01, 0x02, 0x03}
	for _, fork := range []Fork{ForkFulu, ForkGloas} {
		client := newBeaconServer(t, http.StatusOK, string(fork), body)

		block, err := client.Block(context.Background(), 10)
		require.NoError(t, err)
		require.Equal(t, fork, block.Fork)
		require.Equal(t, body, block.Data)

		state, err := client.State(context.Background(), 10)
		require.NoError(t, err)
		require.Equal(t, fork, state.Fork)
		require.Equal(t, body, state.Data)
	}
}

// Beacon nodes report fork names in lower case, but header values are
// matched case-insensitively.
func TestBeaconClientNormalizesForkCase(t *testing.T) {
	client := newBeaconServer(t, http.StatusOK, "Gloas", []byte{0x01})
	block, err := client.Block(context.Background(), 10)
	require.NoError(t, err)
	require.Equal(t, ForkGloas, block.Fork)
}

// Objects of any other fork, or with no fork header, must be rejected.
func TestBeaconClientRejectsUnsupportedForks(t *testing.T) {
	for _, fork := range []string{"electra", "deneb", ""} {
		client := newBeaconServer(t, http.StatusOK, fork, []byte{0x01})
		_, err := client.Block(context.Background(), 10)
		require.ErrorContains(t, err, "unsupported consensus fork", "fork %q", fork)
		_, err = client.State(context.Background(), 10)
		require.ErrorContains(t, err, "unsupported consensus fork", "fork %q", fork)
	}
}

// A 404 on a block means the slot was missed, which callers scanning slots
// must tell apart from failures. A 404 on a state means the node does not
// serve it, which is a failure.
func TestBeaconClientMissedSlot(t *testing.T) {
	client := newBeaconServer(t, http.StatusNotFound, "", nil)

	_, err := client.Block(context.Background(), 10)
	require.ErrorIs(t, err, ErrNoBlockAtSlot)

	_, err = client.State(context.Background(), 10)
	require.Error(t, err)
	require.NotErrorIs(t, err, ErrNoBlockAtSlot)
}
