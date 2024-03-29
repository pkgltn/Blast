// Code generated by rlpgen. DO NOT EDIT.

package types

import "github.com/ethereum/go-ethereum/rlp"
import "io"

func (obj *StateAccount) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	w.WriteUint64(obj.Nonce)
	w.WriteUint64(uint64(obj.Flags))
	if obj.Fixed == nil {
		w.Write(rlp.EmptyString)
	} else {
		if obj.Fixed.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		w.WriteBigInt(obj.Fixed)
	}
	if obj.Shares == nil {
		w.Write(rlp.EmptyString)
	} else {
		if obj.Shares.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		w.WriteBigInt(obj.Shares)
	}
	if obj.Remainder == nil {
		w.Write(rlp.EmptyString)
	} else {
		if obj.Remainder.Sign() == -1 {
			return rlp.ErrNegativeBigInt
		}
		w.WriteBigInt(obj.Remainder)
	}
	w.WriteBytes(obj.Root[:])
	w.WriteBytes(obj.CodeHash)
	w.ListEnd(_tmp0)
	return w.Flush()
}
