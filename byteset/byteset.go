package byteset

const ONEBYTE int = 1 << 8

type ByteSet []byte

func (b ByteSet) BitLen() byte {
	return b[0]
}

func (b ByteSet) DEC() uint64 {
	var res uint64
	byLen := len(b)
	for idx, v := range b {
		res += uint64(int(v) * (ONEBYTE << (byLen - idx)))
	}
	return res
}
func (b ByteSet) BIN() string {
	binStr := ""
	DEC := b.DEC()
	bitLen := int(b.BitLen())
	for i := bitLen; i >= 0; i-- {
		if (DEC>>i)&1 == 1 {
			binStr += "1"
		} else {
			binStr += "0"
		}
	}
	return binStr
}

func NewByteSetByDEC(bitsize byte, value uint64) ByteSet {
	res := []byte{}
	res = append(res, bitsize)
	for value > 0 {
		res = append(res)

	}

	return res
}
