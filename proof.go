package parse

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	OldVersion = "1.0.0" //最初版本（兼容）
	Version1   = "V1"    //最初版本
	Version2   = "V2"    //项目方的自定义版本 不做处理
	Version3   = "V3"    //对存证内容和模板分离 + 合并
	Version4   = "V4"    //不分离 + 合并

	ProofParaData  = "data"
	ProofParaValue = "value"
	ProofParaLabel = "label"

	ProofParaLabelExt = "ext"

	//TemplateExtInfo = "{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"存证名称\",\"label\":\"存证名称\"},{\"data\":{\"format\":\"hash\",\"type\":\"text\",\"value\":\"null\"},\"type\":0,\"key\":\"basehash\",\"label\":\"basehash\"},{\"data\":{\"format\":\"hash\",\"type\":\"text\",\"value\":\"null\"},\"type\":0,\"key\":\"prehash\",\"label\":\"prehash\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"存证类型\",\"label\":\"存证类型\"}],\"type\":3,\"key\":\"\",\"label\":\"ext\"}"
)

// BaseProof  基础的存证转换接口
type BaseProof interface {
	ToTx()(string,error)//use for backend
	ToView()(string,error)//use for externaldb
}

type defaultProof struct {


}

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

func FormatVersion(version string) (string, error) {
	switch {
	case version == "" || version == OldVersion:
		return Version1, nil
	case strings.Index(strings.ToUpper(version), Version1) != -1:
		return Version1, nil
	case strings.Index(strings.ToUpper(version), Version2) != -1:
		return Version2, nil
	case strings.Index(strings.ToUpper(version), Version3) != -1:
		return Version3, nil
	case strings.Index(strings.ToUpper(version), Version4) != -1:
		return Version4, nil
	default:
		return "", fmt.Errorf("version is err , version:" + version)
	}
}

// 将完整数据抽离成简化内容 需要完整数据
func (p *Proof) ComleteDataToContent() error {
	err := p.checkVersion()
	if err != nil {
		return err
	}
	switch p.Version {
	case Version1, Version2, Version4:
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
	case Version3, Version4:
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
	var err error
	p.Version, err = FormatVersion(p.Version)
	if err != nil {
		return err
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
	//t, err = checkTemplateExt(t)
	//if err != nil {
	//	return err
	//}
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

// 现在模板里无需增加ext了
//func checkTemplateExt(t []interface{}) ([]interface{}, error) {
//	//检查是否存在ext
//	if m, ok := t[len(t)-1].(map[string]interface{}); !ok {
//		return nil, fmt.Errorf("template err")
//	} else {
//		s, ok := m[ProofParaLabel].(string)
//		if !ok {
//			return nil, fmt.Errorf("template err , nil label")
//		}
//		if s == ProofParaLabelExt {
//			return t, nil
//		}
//	}
//	//不存在ext 则在末尾添加ext
//	var tmpExtInfo interface{}
//	json.Unmarshal([]byte(TemplateExtInfo), &tmpExtInfo)
//	t = append(t, tmpExtInfo)
//	return t, nil
//}
