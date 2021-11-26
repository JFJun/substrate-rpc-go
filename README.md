# Substrate rpc go sdk

## 简介
	这个包其实是使用[itering/scale.go](https://github.com/itering/scale.go)这个包去代替[JFJun/go-substrate-rpc-client](https://github.com/JFJun/go-substrate-rpc-client)里面的metadata解析。

## 如何使用
	```
	import srg "github.com/JFJun/substrate-rpc-go"

	func main(){
		client,err:=srg.NewSubstrateAPI(url,coinType)
	}

	```
### 如何获取coinType
	可以从https://github.com/itering/scale.go/tree/master/network 下查找对应的网络json文件，填写对应的文件名字，就表示启动对应的网络，
	例如：
	```
	import srg "github.com/JFJun/substrate-rpc-go"

	func main(){
		// 启动acala网络
		coinType:="acala"
		// url 为acala的节点
		client,err:=srg.NewSubstrateAPI(url,coinType)
	}
	```
