# go-encrypt
此工具用于对文件进行加解密操作。
## 加密算法
采用AES的CTR加密模式，用户传入自己需要的密钥，即可进行加解密操作。采用对称加密算法，可以减少加密的运算量，适合大文件加密操作。
## 示例
```
# 查看参数信息
go run encrypt.go -h
# -k string
#        加密密钥
#  -r string
#        要加密/解密的文件名
#  -w string
#        要保存的文件名，默认"原文件名_new"

# 对文件加密
go run encrypt.go -r test.txt -k 123456
# 此时已自动生成 test_new.txt 文件

# 对文件解密
run encrypt.go -r test_new.txt -k 123456 -w restore.txt
```

## 注意事项
- 如果密钥过于简单，有被破解的可能，因此建议使用复杂密钥
- 程序采用分批读取，因此读文件（r参数）和写文件（w参数）请不要重名
- 目前只是单线程操作，对于超大文件的加密耗时会比较长