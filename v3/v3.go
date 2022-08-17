package v3

import (
	"encoding/json"
	"fmt"
	parser "github.com/33cn/proofparse"
	"github.com/33cn/proofparse/v1"
)

//ProofV3
type ProofV3 struct {
	v1.ProofV1
	TxData   string //交易数据,格式任意展开服务需要定制化解析
	Template string //模板数据
}

//NewProofV3
func NewProofV3(data string) parser.BaseProof {
	p := &ProofV3{
		TxData: data,
	}
	return p
}

//ToTx
func (p *ProofV3) ToTx() (string, error) {
	if p.ViewData == "" {
		return "",fmt.Errorf("ProofV3 ViewData is nil ")
	}
	return p.splitValue(p.ViewData)
}

//ToView
func (p *ProofV3) ToView() (string, error) {
	if p.TxData != "" {
		if p.Template == "" {
			return "",fmt.Errorf("ProofV3 Template is nil ")
		}
		return p.mergeValue()
	}
	return "",fmt.Errorf("ProofV3 Content is nil ")
}

func (p *ProofV3) splitValue(s string) (string, error) {
	var in interface{}
	err := json.Unmarshal([]byte(s), &in)
	if err != nil {
		return "", err
	}
	var last = make(map[string]interface{})
	switch in.(type) {
	case []interface{}:
		last, err = splitArry(in.([]interface{}))
		if err != nil {
			return "", err
		}
	case map[string]interface{}:
		label, value, err := splitMap(in.(map[string]interface{}))
		if err != nil {
			return "", err
		}
		last[label] = value
	default:
		return "", fmt.Errorf("json types err")
	}
	out, _ := json.Marshal(last)
	return string(out), nil
}

// 分离数组
func splitArry(i []interface{}) (map[string]interface{}, error) {
	var last = make(map[string]interface{})
	for _, v := range i {
		m, ok := v.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("splitArry types err, not map[string]interface{}")
		}
		label, value, err := splitMap(m)
		if err != nil {
			return nil, err
		}
		last[label] = value
	}
	return last, nil
}

// 分离map
func splitMap(v map[string]interface{}) (string, interface{}, error) {
	var data interface{}
	var getValueOk bool
	label, ok := getLabel(v)
	if !ok {
		return "", nil, fmt.Errorf("get label err")
	}
	data, getValueOk = getValue(v[parser.ProofParaData])
	if !getValueOk {

		res, err := splitArry(v[parser.ProofParaData].([]interface{}))
		if err != nil {
			return "", nil, err
		}
		return label, res, nil
	} else {
		return label, data, nil
	}
}

// 获取label值
func getLabel(in map[string]interface{}) (string, bool) {
	v, ok := in[parser.ProofParaLabel].(string)
	return v, ok
}

// 获取value值
func getValue(in interface{}) (interface{}, bool) {
	switch in.(type) {
	case map[string]interface{}:
		if v, ok := in.(map[string]interface{})[parser.ProofParaValue]; ok {
			return v, true
		}
	case []interface{}:
		var values []interface{}
		for _, v := range in.([]interface{}) {
			value, ok := getValue(v)
			if ok {
				values = append(values, value)
			} else {
				return nil, false
			}
		}
		return values, true
	default:
		return nil, false
	}
	return nil, false
}

// 组装存证值
func (p *ProofV3)mergeValue() (string,error) {
	var temp interface{}
	err := json.Unmarshal([]byte(p.Template), &temp)
	if err != nil {
		return "",err
	}
	var proof interface{}
	err = json.Unmarshal([]byte(p.TxData), &proof)
	if err != nil {
		return "",err
	}
	pMap := proof.(map[string]interface{})
	t := temp.([]interface{})
	t, err = checkTemplateExt(t)
	if err != nil {
		return "",err
	}
	pBack, err := parseData(t, pMap)
	if err != nil {
		return "",err
	}
	pout, _ := json.Marshal(pBack)
	p.ViewData = string(pout)
	return p.ViewData,nil
}

// 组装data的值
func parseData(t []interface{}, p map[string]interface{}) (interface{}, error) {
	var res []map[string]interface{}
	for _, v := range t {
		m := v.(map[string]interface{})
		label := m[parser.ProofParaLabel].(string)
		plabel, ok := p[label]
		if !ok {
			res = append(res, m)
			continue
		}
		m[parser.ProofParaData], ok = parseValue(m[parser.ProofParaData], plabel)
		if !ok {
			return nil, fmt.Errorf("mergeValue err")
		}
		res = append(res, m)
	}
	return res, nil
}

// 组装value的值
func parseValue(data interface{}, proofData interface{}) (interface{}, bool) {
	switch data.(type) {
	//组装单个value
	case map[string]interface{}:
		if _, ok := data.(map[string]interface{})[parser.ProofParaValue]; ok {
			data.(map[string]interface{})[parser.ProofParaValue] = proofData
			return data, true
		} else {
			return nil, false
		}
	case []interface{}:
		//组装value数组
		if _, ok := data.([]interface{})[0].(map[string]interface{})[parser.ProofParaLabel]; !ok {
			var values []interface{}
			for _, v := range proofData.([]interface{}) {
				one := cloneMap(data.([]interface{})[0].(map[string]interface{}))
				one[parser.ProofParaValue] = v
				values = append(values, one)
			}
			return values, true
		}
		//递进下一层data
		pvalues, err := parseData(data.([]interface{}), proofData.(map[string]interface{}))
		if err != nil {
			return nil, false
		}
		return pvalues, true
	}

	return nil, false
}

// 拷贝一份map
func cloneMap(m map[string]interface{}) map[string]interface{} {
	cloneMap := make(map[string]interface{})
	for k, v := range m {
		cloneMap[k] = v
	}
	return cloneMap
}

func checkTemplateExt(t []interface{}) ([]interface{}, error) {
	//检查是否存在ext
	if m, ok := t[len(t)-1].(map[string]interface{}); !ok {
		return nil, fmt.Errorf("template err")
	} else {
		s, ok := m[parser.ProofParaLabel].(string)
		if !ok {
			return nil, fmt.Errorf("template err , nil label")
		}
		if s == parser.ProofParaLabelExt {
			return t, nil
		}
	}
	//不存在ext 则在末尾添加ext
	var tmpExtInfo interface{}
	json.Unmarshal([]byte(parser.TemplateExtInfo), &tmpExtInfo)
	t = append(t, tmpExtInfo)
	return t, nil
}
