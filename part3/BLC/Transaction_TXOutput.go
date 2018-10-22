package BLC

type TXOutput struct {
	Value        int64
	ScriptPubKey string
}

//判断是否是该地址的钱
func (TXOutput *TXOutput) UnLockWithAddress(address string) bool {
	return TXOutput.ScriptPubKey == address
}
