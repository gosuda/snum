package sort

import (
	"bytes"
	"encoding/hex"
	"errors"

	"gosuda.org/snum/snum"
)

func New[T snum.ISnum](num T) *SnumSort {
	return &SnumSort{
		Snum: *snum.New(num),
	}
}

type SnumSort struct {
	snum.Snum
}

func (t *SnumSort) UnmarshalJSON(bt []byte) error {
	if len(bt) < DEF_lenTotalMin+2 || bt[0] != '"' || bt[len(bt)-1] != '"' {
		return errors.New("invalid format")
	}
	enc, err := hex.DecodeString(string(bt[1 : len(bt)-1]))
	if err != nil {
		return err
	}
	return t.Decode(enc)
}

func (t *SnumSort) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('"')
	buf.WriteString(hex.EncodeToString(t.Encode()))
	buf.WriteByte('"')
	return buf.Bytes(), nil
}
