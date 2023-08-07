package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

// https://zhuanlan.zhihu.com/p/165419292
// 示例一
/*
type Person struct {
	Name string
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
	datas := []string{`
    { 
      "type": "person",
      "name":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)

		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person", p)

		case "cat":
			var cat Cat
			mapstructure.Decode(m, &cat)
			fmt.Println("cat", cat)
		}
	}
}

 */

// 示例二
/*
type Person struct {
	Name string `mapstructure:"username"`
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
	datas := []string{`
    { 
      "type": "person",
      "username":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "Age": 1,
      "breed": "Ragdoll"
    }
  `,
		`
    {
      "type": "cat",
      "Name": "rooooose",
      "age": 2,
      "breed": "shorthair"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person", p)

		case "cat":
			var cat Cat
			mapstructure.Decode(m, &cat)
			fmt.Println("cat", cat)
		}
	}
}


 */

// 示例三

type Person struct {
	Name string
}

type Friend1 struct {
	Person
}

//type Friend2 struct {
//	Person `mapstructure:",squash"`
//}

type Friend2 struct {
	Person "mapstructure"
}

func main() {
	datas := []string{`
    { 
      "type": "friend1",
      "person": {
        "name":"dj"
      }
    }
  `,
		`
    {
      "type": "friend2",
      "name": "dj2"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("m----",m)
		switch m["type"].(string) {
		case "friend1":
			var f1 Friend1
			mapstructure.Decode(m, &f1)
			fmt.Println("friend1", f1)

		case "friend2":
			var f2 Friend2
			mapstructure.Decode(m, &f2)
			fmt.Println("friend2", f2)
		}
	}
}