package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//接收参数
	var fileRead, fileWrite, key string
	flag.StringVar(&fileRead, "r", "", "要加密/解密的文件名")
	flag.StringVar(&key, "k", "", "加密密钥")
	flag.StringVar(&fileWrite, "w", "", "要保存的文件名，默认\"原文件名_new\"")
	flag.Parse()

	var err error
	if fileWrite == "" {
		//自动生成要保存的文件名
		fileWrite = makeFileWriteName(fileRead)
	}

	//对文件 加密/解密
	err = cryptFile(fileRead, fileWrite, key)
	if err != nil {
		log.Fatal(err)
	}
}

func makeFileWriteName(read string) string {
	var saveName string

	idx := strings.Index(read, ".")
	saveName = read[0:idx] + "_new" + read[idx:]

	return saveName
}

func cryptFile(fileRead, fileWrite, key string) error {
	//打开要读的文件
	f, err := os.Open(fileRead)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	w, err := os.Create(fileWrite)
	if err != nil {
		return err
	}
	defer func(w *os.File) {
		err := w.Close()
		if err != nil {
			panic(err)
		}
	}(w)

	//获取加密对象
	stream, err := getStream(key)
	if err != nil {
		return err
	}

	//读取文件
	buff := make([]byte, 4096)
	for {
		_, err := f.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		//对文件加密
		cnt, err := encryptContent(stream, buff)
		if err != nil {
			return err
		}

		//写入文件
		_, err = w.Write(cnt)
		if err != nil {
			return err
		}

	}

	return nil
}

func encryptContent(stream cipher.Stream, buff []byte) ([]byte, error) {
	content := make([]byte, len(buff))
	stream.XORKeyStream(content, buff)

	return content, nil
}

func getStream(key string) (cipher.Stream, error) {
	k := md5.Sum([]byte(key))

	block, err := aes.NewCipher(k[:])
	if err != nil {
		return nil, err
	}

	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	return stream, nil
}
