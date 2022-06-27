// Package fabric -----------------------------
// @file      : main.go
// @author    : cayden
// @contact   : cuiran2001@163.com
// @time      : 2022/6/27 16:44
// -------------------------------------------

package main

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric/protos/common"
)

func main() {
	fmt.Println("开始实例化fabric和fabric-sdk-go对象...")

	// mock 使用fabric/protos代码
	tmpBlock := &common.Block{}

	tmpBlock.Data = &common.BlockData{
		Data: [][]byte{[]byte("bytes1"), []byte("bytes2")},
	}

	tmpbs := []byte(tmpBlock.String())

	//mock 使用fabric-sdk-go/pkg/client/resgment代码
	tmpClient := &resmgmt.Client{}
	_, err := tmpClient.CreateConfigSignatureFromReader(nil, bytes.NewBuffer(tmpbs))
	if err != nil {
		fmt.Println("如期签名报错，继续执行...")
	}
	_ = tmpClient
	fmt.Println("分别实例化fabric和fabric-sdk-go中对象成功")
}
