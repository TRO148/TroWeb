package troWeb

import "strings"

// 路由树节点，主要是树的操作，通过matchChild与matChildren搜索子节点，insert与search完成树基本操作
type node struct {
	pattern  string  //匹配路由
	part     string  //路由中的一部分
	children []*node //子节点
	isWild   bool    //是否精确匹配，part含有:或*时为true
}

// matchChild 匹配子节点，返回第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		//如果子节点的part与要匹配的part相同，或者子节点的part为通配符，则返回该子节点
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 匹配子节点，返回所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		//如果子节点的part与要匹配的part相同，或者子节点的part为通配符，则返回该子节点
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 插入节点，height说是第几层就是第几层，遇到没有的就自己填上，遇到有的就继续向下插入
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height { //递归结束条件
		n.pattern = pattern
		return
	}

	part := parts[height]       //获取当前层级的part
	child := n.matchChild(part) //获取当前层级的part对应的子节点
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1) //递归插入子节点
}

// search 查询，不断匹配part，查询到路径的那个点
func (n *node) search(parts []string, height int) *node {
	//只要出现*，则不再进行后续搜索
	if len(parts) == height || strings.HasPrefix(n.part, "*") { //递归结束条件
		if n.pattern == "" { //如果当前节点的pattern为空，则说明没有匹配到
			return nil
		}
		return n
	}

	part := parts[height]             //获取当前层级的part
	children := n.matchChildren(part) //获取所有对应的子节点，包括*与:

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
