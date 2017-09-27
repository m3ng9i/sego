fork说明
============

[本fork](https://github.com/m3ng9i/sego)对[原sego](https://github.com/huichen/sego)代码做了以下修改：

1. 修改 segmenter.go: 删除log调用；载入字典出错时返回error，而不是写入log；
2. 修改 segmenter.go: 将LoadDictionary的参数由files string修改为files ...string；
3. 修改 segmenter.go: 加载字典文件时，增加对参数个数的判断；
4. 增加 stopwords.go：实现停用词处理功能，包括加载停用词词典、判断一个词是否为停用词、从Segment里删除包含停用词的元素；

sego作者：huichen <https://github.com/huichen>
本fork作者：m3ng9i <https://github.com/m3ng9i>

