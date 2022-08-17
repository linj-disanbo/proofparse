package v2

import (
	"github.com/33cn/proofparse"
	"github.com/33cn/proofparse/v1"
)

//ProofV2
type ProofV2 struct {
	v1.ProofV1
	TxData string //交易数据,格式任意展开服务需要定制化解析
}

//NewProofV2
func NewProofV2(data string) parse.BaseProof {
	p := &ProofV2{
		TxData: data,
	}
	return p
}

//ToTx
func (p *ProofV2) ToTx() (string, error) {
	return p.TxData, nil
}

//ToView
func (p *ProofV2) ToView() (string, error) {
	return p.TxData, nil
}
