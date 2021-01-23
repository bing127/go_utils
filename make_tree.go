package main

type TreeConfig struct {
	id string
	pid string
	children string
}

type MakeTree interface {
	construct([]map[string]interface{}) []map[string]interface{}
}

func (c *TreeConfig) construct(nodes []map[string]interface{}) []map[string]interface{} {
	idMap := make(map[string]interface{})
	jsonTree := make([]map[string]interface{},0)

	var id,pid,children string
	if c.id != ""{
		id = c.id
	} else {
		id = "id"
	}

	if c.pid != ""{
		pid = c.pid
	} else {
		pid = "pid"
	}

	if c.children != ""{
		children = c.children
	} else {
		children = "children"
	}



	for _, node := range nodes {
		for k, v := range node {
			if k == id {
				value,ok := v.(string)
				if ok {
					idMap[value] = node
				}
			}
		}
	}

	for _, node := range nodes {
		pid,_ := node[pid].(string)
		parent,ok := idMap[pid].(map[string]interface{})
		if ok {
			p, _ := parent[children].([]map[string]interface{})
			parent[children] = append(p, node)
		} else {
			jsonTree = append(jsonTree, node)
		}
	}

	return jsonTree
}
