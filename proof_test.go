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
	var tmepStr = "[{\"label\":\"相册\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":[{\"type\":\"image\",\"format\":\"hash\",\"value\":\"\"}],\"type\":1,\"key\":\"\",\"label\":\"相册\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"照片描述\"}]},{\"label\":\"ext\",\"key\":\"\",\"type\":3,\"data\":[{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"存证名称\",\"label\":\"存证名称\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"basehash\",\"label\":\"basehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"prehash\",\"label\":\"prehash\"},{\"data\":{\"type\":\"text\",\"format\":\"string\",\"value\":\"\"},\"type\":0,\"key\":\"存证类型\",\"label\":\"存证类型\"}]}]"
	var proof = "{\"ext\":{\"basehash\":\"null\",\"prehash\":\"null\",\"存证名称\":\"相册\",\"存证类型\":\"相册\"},\"相册\":{\"照片描述\":\"\",\"相册\":[\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\",\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"]}}"

	var tc = []struct {
		temp    string
		contend string
		version string
	}{
		{tmepStr, proof, Version3},
		{"[{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"株高（cm）\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"根条数（条）\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"第三叶长（㎝）\"}],\"type\":3,\"key\":\"\",\"label\":\"秧苗素质\"},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"时 间\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"秧 龄\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"叶 龄\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"插深深度（cm）\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"行距×穴距（cm）\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"移栽方法\"}],\"type\":3,\"key\":\"\",\"label\":\"移 栽\"},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"总肥量\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"基 肥\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"蘖 肥\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"调节肥\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"穗 肥\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"type\":0,\"key\":\"\",\"label\":\"叶面追肥\"}],\"type\":3,\"key\":\"\",\"label\":\"本田施肥\"}]","{\"ext\":{\"basehash\":\"null\",\"prehash\":\"null\",\"存证名称\":\"eee\",\"存证类型\":\"秧苗素质\"},\"本田施肥\":{\"叶面追肥\":\"321sadf\",\"基 肥\":\"2\",\"总肥量\":\"1\",\"穗 肥\":\"213\",\"蘖 肥\":\"3\",\"调节肥\":\"4\"},\"秧苗素质\":{\"株高（cm）\":\"1\",\"根条数（条）\":\"2\",\"第三叶长（㎝）\":\"3\"},\"移 栽\":{\"叶 龄\":\"6\",\"插深深度（cm）\":\"7\",\"时 间\":\"4\",\"秧 龄\":\"5\",\"移栽方法\":\"9\",\"行距×穴距（cm）\":\"89\"}}",Version3},
	}
	var res = []string{
		"[{\"data\":[{\"data\":[{\"format\":\"hash\",\"type\":\"image\",\"value\":\"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5\"},{\"format\":\"hash\",\"type\":\"image\",\"value\":\"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4\"}],\"key\":\"\",\"label\":\"相册\",\"type\":1},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"\"},\"key\":\"\",\"label\":\"照片描述\",\"type\":0}],\"key\":\"\",\"label\":\"相册\",\"type\":3},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证名称\",\"label\":\"存证名称\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"basehash\",\"label\":\"basehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"null\"},\"key\":\"prehash\",\"label\":\"prehash\",\"type\":0},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"相册\"},\"key\":\"存证类型\",\"label\":\"存证类型\",\"type\":0}],\"key\":\"\",\"label\":\"ext\",\"type\":3}]",
		"[{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"1\"},\"type\":0,\"key\":\"\",\"label\":\"株高（cm）\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"2\"},\"type\":0,\"key\":\"\",\"label\":\"根条数（条）\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"3\"},\"type\":0,\"key\":\"\",\"label\":\"第三叶长（㎝）\"}],\"type\":3,\"key\":\"\",\"label\":\"秧苗素质\"},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"4\"},\"type\":0,\"key\":\"\",\"label\":\"时 间\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"5\"},\"type\":0,\"key\":\"\",\"label\":\"秧 龄\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"6\"},\"type\":0,\"key\":\"\",\"label\":\"叶 龄\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"7\"},\"type\":0,\"key\":\"\",\"label\":\"插深深度（cm）\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"89\"},\"type\":0,\"key\":\"\",\"label\":\"行距×穴距（cm）\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"9\"},\"type\":0,\"key\":\"\",\"label\":\"移栽方法\"}],\"type\":3,\"key\":\"\",\"label\":\"移 栽\"},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"1\"},\"type\":0,\"key\":\"\",\"label\":\"总肥量\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"2\"},\"type\":0,\"key\":\"\",\"label\":\"基 肥\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"3\"},\"type\":0,\"key\":\"\",\"label\":\"蘖 肥\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"4\"},\"type\":0,\"key\":\"\",\"label\":\"调节肥\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"213\"},\"type\":0,\"key\":\"\",\"label\":\"穗 肥\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"321sadf\"},\"type\":0,\"key\":\"\",\"label\":\"叶面追肥\"}],\"type\":3,\"key\":\"\",\"label\":\"本田施肥\"},{\"data\":[{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"eee\"},\"type\":0,\"key\":\"存证名称\",\"label\":\"存证名称\"},{\"data\":{\"format\":\"hash\",\"type\":\"text\",\"value\":\"null\"},\"type\":0,\"key\":\"basehash\",\"label\":\"basehash\"},{\"data\":{\"format\":\"hash\",\"type\":\"text\",\"value\":\"null\"},\"type\":0,\"key\":\"prehash\",\"label\":\"prehash\"},{\"data\":{\"format\":\"string\",\"type\":\"text\",\"value\":\"秧苗素质\"},\"type\":0,\"key\":\"存证类型\",\"label\":\"存证类型\"}],\"type\":3,\"key\":\"\",\"label\":\"ext\"}]",
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
