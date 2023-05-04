package test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/leancodebox/goose/jsonopt"
)

type Node struct {
	Id       int
	ParentId int
	Value    interface{}
	Weight   int
	Children []*Node
}

func ConvertToTree(nodes []Node) *Node {
	nodeMap := make(map[int]*Node, len(nodes))
	root := &Node{Id: -1} // 添加一个根节点方便查找
	nodeMap[root.Id] = root

	// 1. 遍历 nodes，将节点加入 nodeMap
	for i := range nodes {
		node := &nodes[i]
		n := &Node{
			Id:       node.Id,
			ParentId: node.ParentId,
			Value:    node.Value,
			Weight:   node.Weight,
		}
		nodeMap[n.Id] = n
	}

	// 2. 建立 parent-child 关系
	for _, node := range nodeMap {
		if node.Id == root.Id {
			continue
		}
		parent, ok := nodeMap[node.ParentId]
		if !ok {
			parent = root
		}
		parent.Children = append(parent.Children, node)
	}

	// 3. 对子节点按照权重从大到小排序
	for _, node := range nodeMap {
		sort.Slice(node.Children, func(i, j int) bool {
			return node.Children[i].Weight > node.Children[j].Weight
		})
	}

	// 4. 查找根节点
	return nodeMap[root.Children[0].Id]
}

func TestBuildTree(t *testing.T) {
	nodes := []Node{
		{Id: 1, ParentId: -1, Value: "A"},
		{Id: 7, ParentId: 3, Value: "G"},
		{Id: 2, ParentId: 1, Value: "B"},
		{Id: 3, ParentId: 1, Value: "C"},
		{Id: 4, ParentId: 2, Value: "D"},
		{Id: 5, ParentId: 2, Value: "E"},
		{Id: 6, ParentId: 3, Value: "F"},
	}
	root := ConvertToTree(nodes)

	fmt.Println(jsonopt.EncodeFormat(root))
}
