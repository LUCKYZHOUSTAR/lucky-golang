package cache2go

import (
	"fmt"
	"testing"
)

/**
Golang单元测试对文件名和方法名，参数都有很严格的要求。
http://blog.csdn.net/samxx8/article/details/46894587
　　例如：
　　1、文件名必须以xx_test.go命名
　　2、方法必须是Test[^a-z]开头
　　3、方法参数必须 t *testing.T
　　之前就因为第 2 点没有写对，导致找了半天错误。现在真的让人记忆深刻啊，小小的东西当初看书没仔细。
　　下面分享一点go test的参数解读。来源
*/
func TestCache(t *testing.T) {
	fmt.Println("hi")
}
