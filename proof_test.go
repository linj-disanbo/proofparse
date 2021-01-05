package parse

import (
	"testing"
)

func TestSplitValue(t *testing.T) {
	var s = "[{\"label\":\"相册\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":[{\"type\":\"image\",\"format\":\"hash\",\"value\":\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\"},{\"type\":\"image\",\"format\":\"hash\",\"value\":\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"}],\"type\":1,\"key\":\"\",\"label\":\"相册\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"照片描述\"}]},{\"label\":\"ext\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"相册\"},\"type\":0,\"key\":\"存证名称\",\"label\":\"存证名称\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"null\"},\"type\":0,\"key\":\"basehash\",\"label\":\"basehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"null\"},\"type\":0,\"key\":\"prehash\",\"label\":\"prehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"相册\"},\"type\":0,\"key\":\"存证类型\",\"label\":\"存证类型\"}]}]"
	p:=NewProof(s,"","",Version2)
	err := p.ParseProof()
	if err != nil {
		t.Error(err)
	}

	t.Log(p.Value)
}

func TestMergeValue(t *testing.T) {
	var tmepStr = "[{\"label\":\"相册\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":[{\"type\":\"image\",\"format\":\"hash\",\"value\":\"\"}],\"type\":1,\"key\":\"\",\"label\":\"相册\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"照片描述\"}]},{\"label\":\"ext\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"存证名称\",\"label\":\"存证名称\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"basehash\",\"label\":\"basehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"prehash\",\"label\":\"prehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"存证类型\",\"label\":\"存证类型\"}]}]"
	var proof = "{\"ext\":{\"basehash\":\"null\",\"prehash\":\"null\",\"存证名称\":\"相册\",\"存证类型\":\"相册\"},\"相册\":{\"照片描述\":\"\",\"相册\":[\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\",\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"]}}"

	p:=NewProof("",tmepStr,proof,Version2)
	err := p.ParseProof()
	if err != nil {
		t.Error(err)
	}

	t.Log(p.ComleteData)
}
