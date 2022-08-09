package v3

import (
	parser "github.com/33cn/proofparse"
	"github.com/33cn/proofparse/v1"
	"reflect"
	"testing"
)

func TestNewProofV3(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want parser.BaseProof
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProofV3(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProofV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProofV3_ToTx(t *testing.T) {
	type fields struct {
		ProofV1  v1.ProofV1
		TxData   string
		Template string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProofV3{
				ProofV1:  tt.fields.ProofV1,
				TxData:   tt.fields.TxData,
				Template: tt.fields.Template,
			}
			got, err := p.ToTx()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToTx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProofV3_ToView(t *testing.T) {
	type fields struct {
		ProofV1  v1.ProofV1
		TxData   string
		Template string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProofV3{
				ProofV1:  tt.fields.ProofV1,
				TxData:   tt.fields.TxData,
				Template: tt.fields.Template,
			}
			got, err := p.ToView()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToView() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToView() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProofV3_mergeValue(t *testing.T) {
	type fields struct {
		ProofV1  v1.ProofV1
		TxData   string
		Template string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProofV3{
				ProofV1:  tt.fields.ProofV1,
				TxData:   tt.fields.TxData,
				Template: tt.fields.Template,
			}
			got, err := p.mergeValue()
			if (err != nil) != tt.wantErr {
				t.Errorf("mergeValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mergeValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProofV3_splitValue(t *testing.T) {
	type fields struct {
		ProofV1  v1.ProofV1
		TxData   string
		Template string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProofV3{
				ProofV1:  tt.fields.ProofV1,
				TxData:   tt.fields.TxData,
				Template: tt.fields.Template,
			}
			got, err := p.splitValue(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("splitValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkTemplateExt(t *testing.T) {
	type args struct {
		t []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkTemplateExt(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkTemplateExt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkTemplateExt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cloneMap(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneMap(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLabel(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getLabel(tt.args.in)
			if got != tt.want {
				t.Errorf("getLabel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getLabel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getValue(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getValue(tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseData(t *testing.T) {
	type args struct {
		t []interface{}
		p map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseData(tt.args.t, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseValue(t *testing.T) {
	type args struct {
		data      interface{}
		proofData interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseValue(tt.args.data, tt.args.proofData)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_splitArry(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitArry(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitArry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitArry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitMap(t *testing.T) {
	type args struct {
		v map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := splitMap(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("splitMap() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("splitMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}