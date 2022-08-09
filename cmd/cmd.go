package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	parse "github.com/33cn/proofparse"
)

var (
	// detail 模板原始数据详情
	detail = flag.String("d", "", "the template raw detail")
	// needBeautify 是否需要美化结果
	needBeautify = flag.Bool("b", false, "need to beautify the result")
)

func main() {
	flag.Parse()

	p := parse.NewProof(*detail, "", "", parse.Version3)
	err := p.ComleteDataToContent()
	if err != nil {
		log.Println(err)
		return
	}

	if *needBeautify {
		var buf bytes.Buffer
		err := json.Indent(&buf, []byte(p.Content), "", "    ")
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(buf.String())
		return
	}

	fmt.Println(p.Content)
}
