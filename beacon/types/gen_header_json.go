// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum-arbitrum/common"
)

var _ = (*headerMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (h Header) MarshalJSON() ([]byte, error) {
	type Header struct {
		Slot          common.Decimal `gencodec:"required" json:"slot"`
		ProposerIndex common.Decimal `gencodec:"required" json:"proposer_index"`
		ParentRoot    common.Hash    `gencodec:"required" json:"parent_root"`
		StateRoot     common.Hash    `gencodec:"required" json:"state_root"`
		BodyRoot      common.Hash    `gencodec:"required" json:"body_root"`
	}
	var enc Header
	enc.Slot = common.Decimal(h.Slot)
	enc.ProposerIndex = common.Decimal(h.ProposerIndex)
	enc.ParentRoot = h.ParentRoot
	enc.StateRoot = h.StateRoot
	enc.BodyRoot = h.BodyRoot
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (h *Header) UnmarshalJSON(input []byte) error {
	type Header struct {
		Slot          *common.Decimal `gencodec:"required" json:"slot"`
		ProposerIndex *common.Decimal `gencodec:"required" json:"proposer_index"`
		ParentRoot    *common.Hash    `gencodec:"required" json:"parent_root"`
		StateRoot     *common.Hash    `gencodec:"required" json:"state_root"`
		BodyRoot      *common.Hash    `gencodec:"required" json:"body_root"`
	}
	var dec Header
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Slot == nil {
		return errors.New("missing required field 'slot' for Header")
	}
	h.Slot = uint64(*dec.Slot)
	if dec.ProposerIndex == nil {
		return errors.New("missing required field 'proposer_index' for Header")
	}
	h.ProposerIndex = uint64(*dec.ProposerIndex)
	if dec.ParentRoot == nil {
		return errors.New("missing required field 'parent_root' for Header")
	}
	h.ParentRoot = *dec.ParentRoot
	if dec.StateRoot == nil {
		return errors.New("missing required field 'state_root' for Header")
	}
	h.StateRoot = *dec.StateRoot
	if dec.BodyRoot == nil {
		return errors.New("missing required field 'body_root' for Header")
	}
	h.BodyRoot = *dec.BodyRoot
	return nil
}
