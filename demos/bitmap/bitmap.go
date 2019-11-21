package bitmap

type BitMap []byte

func NewBitMap(length uint) BitMap {
	return make(BitMap, length/8 + 1)
}

func (b BitMap) Set(value uint) {

}
