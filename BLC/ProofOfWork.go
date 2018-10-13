package BLC

type ProofOfWork struct {
	Block *Block
}

func NewProofOfWork(block *Block) *ProofOfWork {

	return &ProofOfWork{block}
}
func (p *ProofOfWork) Run() ([]byte, int64) {

	return nil, 0
}
