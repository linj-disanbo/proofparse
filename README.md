# 存证解析库

## 背景
```
用于解析存证的数据
在创建存证的时候，区分版本
```
*  第一版本的存证内模板和存证内容混在一起，无需分离

*  第二版本的存证内容格式完全由项目方自定义,不做任何处理

*  第三版本的存证可以直接将模板和存证内容分离,
,例

```
[{"label":"相册","key":"","type":3,"data":[{"data":[{"type":"image","format":"hash","value":"4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5"},{"type":"image","format":"hash","value":"b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4"}],"type":1,"key":"","label":"相册"},{"data":{"type":"text","format":"string","value":""},"type":0,"key":"","label":"照片描述"}]},{"label":"ext","key":"","type":3,"data":[{"data":{"type":"text","format":"string","value":"相册"},"type":0,"key":"存证名称","label":"存证名称"},{"data":{"type":"text","format":"string","value":"null"},"type":0,"key":"basehash","label":"basehash"},{"data":{"type":"text","format":"string","value":"null"},"type":0,"key":"prehash","label":"prehash"},{"data":{"type":"text","format":"string","value":"相册"},"type":0,"key":"存证类型","label":"存证类型"}]}]

data的内容为两种
一种是嵌层一层完整数据，内容为数组，元素的key为"label""type""data"
一种是存证具体数据,内容为数组或者map，元素的key为"format""type""value"

分离后
{"ext":{"basehash":"null","prehash":"null","存证名称":"相册","存证类型":"相册"},"相册":{"照片描述":"","相册":["4bb42ecfe15a99ee586031b7a5a20a5d15c4364780f5ea14d278095d620619f5","b45f6c7728e9b59f9ad381f36c29945317ba646e49ad3b73511ec1731c22dff4"]}}

模板为

[{"label":"相册","key":"","type":3,"data":[{"data":[{"type":"image","format":"hash","value":""}],"type":1,"key":"","label":"相册"},{"data":{"type":"text","format":"string","value":""},"type":0,"key":"","label":"照片描述"}]},{"label":"ext","key":"","type":3,"data":[{"data":{"type":"text","format":"string","value":""},"type":0,"key":"存证名称","label":"存证名称"},{"data":{"type":"text","format":"string","value":""},"type":0,"key":"basehash","label":"basehash"},{"data":{"type":"text","format":"string","value":""},"type":0,"key":"prehash","label":"prehash"},{"data":{"type":"text","format":"string","value":""},"type":0,"key":"存证类型","label":"存证类型"}]}]
```
