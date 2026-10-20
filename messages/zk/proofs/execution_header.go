// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/rlp"
)

// headerField is an execution block header field, named as eth_getBlockByHash
// returns it. A quantity encodes as an RLP integer, anything else as an RLP
// string.
type headerField struct {
	name     string
	quantity bool
}

// legacyHeaderFields are the 15 header fields every fork has, in RLP order.
// receiptsRoot is item 5 in every fork.
var legacyHeaderFields = []headerField{
	{name: "parentHash"},
	{name: "sha3Uncles"},
	{name: "miner"},
	{name: "stateRoot"},
	{name: "transactionsRoot"},
	{name: "receiptsRoot"},
	{name: "logsBloom"},
	{name: "difficulty", quantity: true},
	{name: "number", quantity: true},
	{name: "gasLimit", quantity: true},
	{name: "gasUsed", quantity: true},
	{name: "timestamp", quantity: true},
	{name: "extraData"},
	{name: "mixHash"},
	{name: "nonce"},
}

// optionalHeaderFields are the header fields later forks append, in fork
// order: London, Shanghai, Cancun (3), Prague and Amsterdam (2). A header
// carries a prefix of them.
var optionalHeaderFields = []headerField{
	{name: "baseFeePerGas", quantity: true},
	{name: "withdrawalsRoot"},
	{name: "blobGasUsed", quantity: true},
	{name: "excessBlobGas", quantity: true},
	{name: "parentBeaconBlockRoot"},
	{name: "requestsHash"},
	{name: "blockAccessListHash"},
	{name: "slotNumber", quantity: true},
}

// encodeExecutionHeader returns the RLP encoding of an execution block
// header, whose keccak256 is the block hash, from the JSON of an
// eth_getBlockByHash result. It encodes the fields by name in fork order
// rather than through libevm's header type, which ends at the Cancun fields.
// A node could report wrong fields, so callers must check the encoding's hash
// against a trusted block hash.
func encodeExecutionHeader(headerJSON []byte) ([]byte, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(headerJSON, &fields); err != nil {
		return nil, fmt.Errorf("failed to decode header: %w", err)
	}

	items := make([]any, 0, len(legacyHeaderFields)+len(optionalHeaderFields))
	for _, field := range legacyHeaderFields {
		raw, ok := fields[field.name]
		if !ok {
			return nil, fmt.Errorf("missing header field %q", field.name)
		}
		item, err := decodeHeaderField(field, raw)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	missing := ""
	for _, field := range optionalHeaderFields {
		raw, ok := fields[field.name]
		if !ok {
			if missing == "" {
				missing = field.name
			}
			continue
		}
		if missing != "" {
			return nil, fmt.Errorf("header field %q follows missing field %q", field.name, missing)
		}
		item, err := decodeHeaderField(field, raw)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return rlp.EncodeToBytes(items)
}

// decodeHeaderField decodes a hex JSON value into the type that RLP-encodes
// it: a big integer for quantities, bytes otherwise.
func decodeHeaderField(field headerField, raw json.RawMessage) (any, error) {
	if field.quantity {
		var quantity hexutil.Big
		if err := json.Unmarshal(raw, &quantity); err != nil {
			return nil, fmt.Errorf("invalid header field %q: %w", field.name, err)
		}
		return (*big.Int)(&quantity), nil
	}
	var data hexutil.Bytes
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("invalid header field %q: %w", field.name, err)
	}
	return []byte(data), nil
}
