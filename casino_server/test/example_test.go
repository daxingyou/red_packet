package mongodb_test

import (
	"fmt"
	"github.com/name5566/leaf/db/mongodb"
	"gopkg.in/mgo.v2"
)

func Example() {
	c, err := mongodb.Dial("localhost", 51668)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// session
	s := c.Ref()
	defer c.UnRef(s)
	err = s.DB("test").C("counters").RemoveId("test")
	if err != nil && err != mgo.ErrNotFound {
		fmt.Println(err)
		return
	}

	// auto increment
	//创建自增字段
	err = c.EnsureCounter("test", "counters", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	//增长自增字段
	fmt.Println("开始测试自增字段")
	for i := 0; i < 3; i++ {
		id, err := c.NextSeq("test", "counters", "test")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id)
	}

	// index
	c.EnsureUniqueIndex("test", "counters", []string{"key1"})

	//nnn := "ceshi Name"
	//s.DB("test").C("counters").Insert(&bbproto.N{
	//	Name:&nnn,
	//})

	// Output:
	// 1
	// 2
	// 3
}
