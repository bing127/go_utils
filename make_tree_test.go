package go_utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMakeTree(t *testing.T) {
	r := []map[string]interface{}{
		{"cid": "1", "name": "动物", "pid": "0"},
		{"cid": "2", "name": "鸟类", "pid": "5"},
		{"cid": "3", "name": "无脊椎动物", "pid": "1"},
		{"cid": "4", "name": "哺乳动物", "pid": "5"},
		{"cid": "5", "name": "脊椎动物", "pid": "1"},
		{"cid": "6", "name": "喜鹊", "pid": "2"},
		{"cid": "7", "name": "蚯蚓", "pid": "3"},
		{"cid": "8", "name": "老虎", "pid": "4"},
		{"cid": "9", "name": "蛇", "pid": "3"},
	}

	config := TreeConfig{
		ID:       "cid",
		PID:      "pid",
		Children: "kids",
	}

	i := config.Construct(r)
	b, _ := json.Marshal(i)
	fmt.Println(string(b))
}
