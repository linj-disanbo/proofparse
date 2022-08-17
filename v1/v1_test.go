package v1

import (
	parser "github.com/33cn/proofparse"
	"reflect"
	"testing"
)

var (
	Testdata1  ="[{\"data\":[{\"data\":[{\"format\":\"hash\",\"type\":\"image\",\"value\":\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\"},{\"format\":\"hash\",\"type\":\"image\",\"value\":\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"}],\"key\":\"\",\"label\":\"相册\",\"type\":1},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"\",\"label\":\"照片描述\",\"type\":0}],\"key\":\"\",\"label\":\"相册\",\"type\":3},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证名称\",\"label\":\"存证名称\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"basehash\",\"label\":\"basehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"prehash\",\"label\":\"prehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证类型\",\"label\":\"存证类型\",\"type\":0}],\"key\":\"\",\"label\":\"ext\",\"type\":3}]"
	Testdata2  ="[{\"data\":[{\"data\":[{\"format\":\"hash\",\"type\":\"image\",\"value\":\"ff\"},{\"format\":\"hash\",\"type\":\"image\",\"value\":\"s\"}],\"key\":\"\",\"label\":\"相册\",\"type\":1},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"\",\"label\":\"照片描述\",\"type\":0}],\"key\":\"\",\"label\":\"相册\",\"type\":3},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证名称\",\"label\":\"存证名称\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"basehash\",\"label\":\"basehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"prehash\",\"label\":\"prehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证类型\",\"label\":\"存证类型\",\"type\":0}],\"key\":\"\",\"label\":\"ext\",\"type\":3}]"
)

func TestNewProofV1(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want parser.BaseProof
	}{
		{"Testdata1",args{data:Testdata1},&ProofV1{ViewData:Testdata1}},
		{"Testdata2",args{data:Testdata2},&ProofV1{ViewData:Testdata2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProofV1(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProofV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProofV1_ToTx(t *testing.T) {
	type fields struct {
		ViewData string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Testdata1",fields{Testdata1},Testdata1,false},
		{"Testdata2",fields{Testdata2},Testdata2,false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProofV1{
				ViewData: tt.fields.ViewData,
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

func TestProofV1_ToView(t *testing.T) {
	type fields struct {
		ViewData string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Testdata1",fields{Testdata1},Testdata1,false},
		{"Testdata2",fields{Testdata2},Testdata2,false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProofV1{
				ViewData: tt.fields.ViewData,
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