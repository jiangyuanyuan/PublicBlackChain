package BLC

type TXInput struct {
	TxHash   []byte
	Vout     int64
	SciptSig string
}

//判断是否是该地址的钱
func (TXInput *TXInput) UnLockWithAddress(address string) bool {
	return TXInput.SciptSig == address
}
