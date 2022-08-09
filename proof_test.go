package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitValue(t *testing.T) {
	var s = "[{\"data\":[{\"data\":[{\"format\":\"hash\",\"type\":\"image\",\"value\":\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\"},{\"format\":\"hash\",\"type\":\"image\",\"value\":\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"}],\"key\":\"\",\"label\":\"相册\",\"type\":1},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"\",\"label\":\"照片描述\",\"type\":0}],\"key\":\"\",\"label\":\"相册\",\"type\":3},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证名称\",\"label\":\"存证名称\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"basehash\",\"label\":\"basehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"prehash\",\"label\":\"prehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证类型\",\"label\":\"存证类型\",\"type\":0}],\"key\":\"\",\"label\":\"ext\",\"type\":3}]"
	var tc = []struct {
		com     string
		version string
	}{
		{s, Version3},
		{"[{\"label\":\"相册\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":[{\"type\":\"image\",\"format\":\"hash\",\"value\":\"4a\"},{\"type\":\"image\",\"format\":\"hash\",\"value\":\"b4\"}],\"type\":1,\"key\":\"\",\"label\":\"相册\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"照片描述\"}]},{\"label\":\"ext\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"相册\"},\"type\":0,\"key\":\"存证名称\",\"label\":\"存证名称\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"null\"},\"type\":0,\"key\":\"basehash\",\"label\":\"basehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"null\"},\"type\":0,\"key\":\"prehash\",\"label\":\"prehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"相册\"},\"type\":0,\"key\":\"存证类型\",\"label\":\"存证类型\"}]}]", Version3},
	}
	var res = []string{
		"{\"ext\":{\"basehash\":\"null\",\"prehash\":\"null\",\"存证名称\":\"相册\",\"存证类型\":\"相册\"},\"相册\":{\"照片描述\":\"\",\"相册\":[\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\",\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"]}}",
		"{\"ext\":{\"basehash\":\"null\",\"prehash\":\"null\",\"存证名称\":\"相册\",\"存证类型\":\"相册\"},\"相册\":{\"照片描述\":\"\",\"相册\":[\"4a\",\"b4\"]}}",
	}
	for i, v := range tc {
		p := NewProof(v.com, "", "", v.version)
		err := p.ComleteDataToContent()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, p.Content, res[i])
	}
}

func TestMergeValue(t *testing.T) {
	var tmepStr = "[{\"label\":\"相册\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":[{\"type\":\"image\",\"format\":\"hash\",\"value\":\"\"}],\"type\":1,\"key\":\"\",\"label\":\"相册\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"照片描述\"}]},{\"label\":\"ext\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"存证名称\",\"label\":\"存证名称\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"null\"},\"type\":0,\"key\":\"basehash\",\"label\":\"basehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"null\"},\"type\":0,\"key\":\"prehash\",\"label\":\"prehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"存证类型\",\"label\":\"存证类型\"}]}]"
	var proof = "{\"ext\":{\"basehash\":\"null\",\"prehash\":\"null\",\"存证名称\":\"相册\",\"存证类型\":\"相册\"},\"相册\":{\"照片描述\":\"\",\"相册\":[\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\",\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"]}}"
	var nilExtProof="{\"相册\":{\"照片描述\":\"\",\"相册\":[\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\",\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"]}}"

	var tc = []struct {
		temp    string
		contend string
		version string
	}{
		{tmepStr, proof, Version3},
		{tmepStr, nilExtProof, Version3},
	}
	var res = []string{
		"[{\"data\":[{\"data\":[{\"format\":\"hash\",\"type\":\"image\",\"value\":\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\"},{\"format\":\"hash\",\"type\":\"image\",\"value\":\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"}],\"key\":\"\",\"label\":\"相册\",\"type\":1},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"\",\"label\":\"照片描述\",\"type\":0}],\"key\":\"\",\"label\":\"相册\",\"type\":3},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证名称\",\"label\":\"存证名称\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"basehash\",\"label\":\"basehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"prehash\",\"label\":\"prehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证类型\",\"label\":\"存证类型\",\"type\":0}],\"key\":\"\",\"label\":\"ext\",\"type\":3}]",
		"[{\"data\":[{\"data\":[{\"format\":\"hash\",\"type\":\"image\",\"value\":\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\"},{\"format\":\"hash\",\"type\":\"image\",\"value\":\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"}],\"key\":\"\",\"label\":\"相册\",\"type\":1},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"\",\"label\":\"照片描述\",\"type\":0}],\"key\":\"\",\"label\":\"相册\",\"type\":3},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"存证名称\",\"label\":\"存证名称\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"basehash\",\"label\":\"basehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"prehash\",\"label\":\"prehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"存证类型\",\"label\":\"存证类型\",\"type\":0}],\"key\":\"\",\"label\":\"ext\",\"type\":3}]",
	}
	for i, v := range tc {
		p := NewProof("", v.temp, v.contend, v.version)
		err := p.ContentToComleteData()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, p.ComleteData, res[i])
	}
}


func TestSplitValueV4(t *testing.T) {
	var s = "{\"ext\":{\"basehash\":\"null\",\"prehash\":\"null\",\"存证名称\":\"相册\",\"存证类型\":\"相册\"},\"相册\":{\"照片描述\":\"\",\"相册\":[\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\",\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"]}}"
	var tc = []struct {
		com     string
		version string
	}{
		{s, Version4},
		}
	var res = []string{
		"{\"ext\":{\"basehash\":\"null\",\"prehash\":\"null\",\"存证名称\":\"相册\",\"存证类型\":\"相册\"},\"相册\":{\"照片描述\":\"\",\"相册\":[\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\",\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"]}}",
		}
	for i, v := range tc {
		p := NewProof(v.com, "", "", v.version)
		err := p.ComleteDataToContent()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, p.Content, res[i])
	}
}

func TestMergeValueV4(t *testing.T) {
	var tmepStr = "[{\"label\":\"相册\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":[{\"type\":\"image\",\"format\":\"hash\",\"value\":\"\"}],\"type\":1,\"key\":\"\",\"label\":\"相册\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"照片描述\"}]},{\"label\":\"ext\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"存证名称\",\"label\":\"存证名称\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"null\"},\"type\":0,\"key\":\"basehash\",\"label\":\"basehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"null\"},\"type\":0,\"key\":\"prehash\",\"label\":\"prehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"存证类型\",\"label\":\"存证类型\"}]}]"
	var proof = "{\"ext\":{\"basehash\":\"null\",\"prehash\":\"null\",\"存证名称\":\"相册\",\"存证类型\":\"相册\"},\"相册\":{\"照片描述\":\"\",\"相册\":[\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\",\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"]}}"
	var nilExtProof="{\"相册\":{\"照片描述\":\"\",\"相册\":[\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\",\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"]}}"

	var tc = []struct {
		temp    string
		contend string
		version string
	}{
		{tmepStr, proof, Version4},
		{tmepStr, nilExtProof, Version4},
	}
	var res = []string{
		"[{\"data\":[{\"data\":[{\"format\":\"hash\",\"type\":\"image\",\"value\":\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\"},{\"format\":\"hash\",\"type\":\"image\",\"value\":\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"}],\"key\":\"\",\"label\":\"相册\",\"type\":1},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"\",\"label\":\"照片描述\",\"type\":0}],\"key\":\"\",\"label\":\"相册\",\"type\":3},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证名称\",\"label\":\"存证名称\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"basehash\",\"label\":\"basehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"prehash\",\"label\":\"prehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证类型\",\"label\":\"存证类型\",\"type\":0}],\"key\":\"\",\"label\":\"ext\",\"type\":3}]",
		"[{\"data\":[{\"data\":[{\"format\":\"hash\",\"type\":\"image\",\"value\":\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\"},{\"format\":\"hash\",\"type\":\"image\",\"value\":\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"}],\"key\":\"\",\"label\":\"相册\",\"type\":1},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"\",\"label\":\"照片描述\",\"type\":0}],\"key\":\"\",\"label\":\"相册\",\"type\":3},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"存证名称\",\"label\":\"存证名称\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"basehash\",\"label\":\"basehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"prehash\",\"label\":\"prehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"存证类型\",\"label\":\"存证类型\",\"type\":0}],\"key\":\"\",\"label\":\"ext\",\"type\":3}]",
	}
	for i, v := range tc {
		p := NewProof("", v.temp, v.contend, v.version)
		err := p.ContentToComleteData()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, p.ComleteData, res[i])
	}
}
