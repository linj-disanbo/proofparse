package v1

import "github.com/33cn/proofparse"

//ProofV1
type ProofV1 struct {
	ViewData string //完整数据
}

//NewProofV1
func NewProofV1(data string) parse.BaseProof {
	p := &ProofV1{
		ViewData: data,
	}
	return p
}

//ToTx
func (p *ProofV1) ToTx() (string, error) {
	return p.ViewData, nil
}

//ToView
func (p *ProofV1) ToView() (string, error) {
	return p.ViewData, nil
}