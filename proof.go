package parse

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	OldVersion = "1.0.0"
	Version1   = "V1"
	Version2   = "V2"
	Version3   = "V3"

	ProofParaData  = "data"
	ProofParaValue = "value"
	ProofParaLabel = "label"
)

type Proof struct {
	ComleteData string //完整数据
	Template    string //模板
	Content     string //存证内容
	Version     string //存证版本
}

func NewProof(data, template, value, version string) *Proof {
	//目前默认为v1
	if version == "" {
		version = Version1
	}
	p := &Proof{
		ComleteData: data,
		Template:    template,
		Content:     value,
		Version:     version,
	}
	return p
}

// 将完整数据抽离成简化内容 需要完整数据
func (p *Proof) ComleteDataToContent() error {
	err := p.checkVersion()
	if err != nil {
		return err
	}
	switch p.Version {
	case Version1, Version2:
		//存证内容和完整数据保持一致
		if p.ComleteData != "" {
			p.Content = p.ComleteData
			return nil
		}
		return fmt.Errorf("ComleteData is nil %s", p.Version)

	case Version3:
		//分离模板和内容
		if p.ComleteData != "" {
			return p.splitValue()
		}
		return fmt.Errorf("ComleteData is nil %s", p.Version)
	default:
		return fmt.Errorf("version err  version:%s", p.Version)
	}

}

// 将存证内容和模板合并成完整数据
func (p *Proof) ContentToComleteData() error {
	p.checkVersion()
	switch p.Version {
	case Version1, Version2:
		//存证内容和完整数据保持一致
		if p.Content != "" {
			p.ComleteData = p.Content
			return nil
		}
		return fmt.Errorf("Content is nil %s", p.Version)
	case Version3:
		//合并模板和内容
		if p.Content != "" {
			if p.Template == "" {
				return fmt.Errorf("Template is nil %s", p.Version)
			}
			return p.mergeValue()
		}
		return fmt.Errorf("Content is nil %s", p.Version)
	default:
		return fmt.Errorf("version err  version:%s", p.Version)
	}

}

func (p *Proof) checkVersion() error {
	p.Version = strings.ToUpper(p.Version)
	switch {
	case p.Version == OldVersion:
		p.Version = Version1
	case strings.Index(strings.ToUpper(p.Version), Version1) != -1:
		p.Version = Version1
	case strings.Index(strings.ToUpper(p.Version), Version2) != -1:
		p.Version = Version2
	case strings.Index(strings.ToUpper(p.Version), Version3) != -1:
		p.Version = Version3
	default:
		return fmt.Errorf("Version err version:%s", p.Version)
	}
	return nil
}

// 分离模板
func (p *Proof) splitValue() error {
	s := p.ComleteData
	var in interface{}
	err := json.Unmarshal([]byte(s), &in)
	if err != nil {
		return err
	}
	var last = make(map[string]interface{})
	switch in.(type) {
	case []interface{}:
		last, err = splitArry(in.([]interface{}))
		if err != nil {
			return err
		}
	case map[string]interface{}:
		label, value, err := splitMap(in.(map[string]interface{}))
		if err != nil {
			return err
		}
		last[label] = value
	default:
		return fmt.Errorf("json types err")
	}
	out, _ := json.Marshal(last)
	p.Content = string(out)
	return nil
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
	data, getValueOk = getValue(v[ProofParaData])
	if !getValueOk {
		res, err := splitArry(v[ProofParaData].([]interface{}))
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
	v, ok := in[ProofParaLabel].(string)
	return v, ok
}

// 获取value值
func getValue(in interface{}) (interface{}, bool) {
	switch in.(type) {
	case map[string]interface{}:
		if v, ok := in.(map[string]interface{})[ProofParaValue]; ok {
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
func (p *Proof) mergeValue() error {
	var temp interface{}
	err := json.Unmarshal([]byte(p.Template), &temp)
	if err != nil {
		return err
	}
	var proof interface{}
	err = json.Unmarshal([]byte(p.Content), &proof)
	if err != nil {
		return err
	}
	pMap := proof.(map[string]interface{})
	t := temp.([]interface{})
	pBack, err := parseData(t, pMap)
	if err != nil {
		return err
	}
	pout, _ := json.Marshal(pBack)
	p.ComleteData = string(pout)
	return nil
}

// 组装data的值
func parseData(t []interface{}, p map[string]interface{}) (interface{}, error) {
	var res []map[string]interface{}
	for _, v := range t {
		m := v.(map[string]interface{})
		label := m[ProofParaLabel].(string)
		plabel, ok := p[label]
		if !ok {
			res = append(res, m)
			continue
		}
		m[ProofParaData], ok = parseValue(m[ProofParaData], plabel)
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
		if _, ok := data.(map[string]interface{})[ProofParaValue]; ok {
			data.(map[string]interface{})[ProofParaValue] = proofData
			return data, true
		} else {
			return nil, false
		}
	case []interface{}:
		//组装value数组
		if _, ok := data.([]interface{})[0].(map[string]interface{})[ProofParaLabel]; !ok {
			var values []interface{}
			for _, v := range proofData.([]interface{}) {
				one := cloneMap(data.([]interface{})[0].(map[string]interface{}))
				one[ProofParaValue] = v
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
