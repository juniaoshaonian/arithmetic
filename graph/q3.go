package graph

//【题目描述】
//
//给定一个有 n 个节点的有向无环图，节点编号从 1 到 n。请编写一个程序，找出并返回所有从节点 1 到节点 n 的路径。每条路径应以节点编号的列表形式表示。
//
//【输入描述】
//
//第一行包含两个整数 N，M，表示图中拥有 N 个节点，M 条边
//
//后续 M 行，每行包含两个整数 s 和 t，表示图中的 s 节点与 t 节点中有一条路径
//
//【输出描述】
//
//输出所有的可达路径，路径中所有节点的后面跟一个空格，每条路径独占一行，存在多条路径，路径输出的顺序可任意。
//
//如果不存在任何一条路径，则输出 -1。
//
//注意输出的序列中，最后一个节点后面没有空格！ 例如正确的答案是 1 3 5,而不是 1 3 5， 5后面没有空格！
//
//【输入示例】

type Node struct {
	number int
	Pre    []*Node
	Nexts  []*Node
	Side   [][]int
}

func GetGraphSide(sides [][]int) [][]int {
	// 找到所有节点
	nodeMap := make(map[int]*Node)
	for _, side := range sides {
		n1 := side[0]
		n2 := side[1]
		if v, ok := nodeMap[n1]; ok {
			v.Side = append(v.Side, side)
		} else {
			nodeMap[n1] = &Node{
				number: n1,
				Side: [][]int{
					side,
				},
			}
		}

		if v, ok := nodeMap[n2]; ok {
			v.Side = append(v.Side, side)
		} else {
			nodeMap[n2] = &Node{
				number: n2,
				Side: [][]int{
					side,
				},
			}
		}
	}
	var root *Node
	for number := range nodeMap {
		node := nodeMap[number]
		for _, side := range node.Side {
			preNum := side[0]
			nextNum := side[1]
			if nextNum != number {
				nextNode := nodeMap[nextNum]
				if nextNode != nil {
					node.Nexts = append(node.Nexts, nextNode)
				}
			}
			if preNum != number {
				preNode := nodeMap[preNum]
				if preNode != nil {
					node.Pre = append(node.Pre, preNode)
				}
			}
			if len(node.Pre) == 0 {
				root = node
			}
		}
	}
	ans := make([][]int,0,5)
	if root != nil {
		getLen(root,[]int{},&ans)
	}
	return ans
}

func getLen(n *Node, ans []int,valMap *[][]int) {
	ans = append(ans,n.number)
	if len(n.Nexts) == 0 {
		*valMap = append(*valMap,ans)
		return
	}
	for _, next := range n.Nexts {
		getLen(next, ans,valMap)
	}
}
