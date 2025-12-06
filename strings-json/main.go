package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Strings()
	StringBuilder()
	Json()
}

func Strings() {
	fmt.Println("------------- STRINGS --------------")
	fmt.Println("")
	fmt.Println("declaring string")
	var s string = ""
	s = "hello"
	fmt.Println("s: ", s)
	fmt.Println("len of s: ", len(s))
	fmt.Println("indexing s, gives byte: ", s[3])
	fmt.Println("indexing s and converting to string: ", string(s[3]))
	fmt.Println("slicing s: ", s[1:3])

	fmt.Println("Concatinating string")
	a := "world"
	c := s + a
	fmt.Println("c: ", c)
	fmt.Print("String tab literals \t and backsapce literal \b we also have new line \n")

	fmt.Println("String -><- byte")
	bs := "I am a coder"
	fmt.Printf("string ->%v to byte ->%v to string again -> %v", bs, []byte(bs), string([]byte(bs)))

	b := "String for practice"
	fmt.Println("->", b)
	fmt.Println("string to upper: ", strings.ToUpper(b))
	fmt.Println("to lower: ", strings.ToLower(b))
	fmt.Println("has prefix S? ", strings.HasPrefix(b, "S"))
	fmt.Println("has suffix ice? ", strings.HasSuffix(b, "ice"))
	fmt.Println("replace all for with to ", strings.ReplaceAll(b, "for", "to"))
	fmt.Println("convert to tile: ", strings.Title(b))
	fmt.Println("trim String: ", strings.Trim(b, "String"))
	fmt.Printf("compare %v with %v: %v\n", b, c, strings.Compare(b, c))
	fmt.Printf("index of for in %v: %v\n", b, strings.Index(s, "for"))

	fmt.Println("")
	fmt.Println("converting string to other type")
	x := 2
	xs := strconv.Itoa(x)
	fmt.Println("x int string: ", xs)

	x, err := strconv.Atoi(xs)
	if err != nil {
		fmt.Println("error converting string to int")
	}
	fmt.Println("string to int ", x)
	fmt.Println("")

}

func StringBuilder() {
	fmt.Println("------------- STRING BUILDER --------------")
	fmt.Println("")
	var sb strings.Builder
	fmt.Println("string builder: ", sb)
	fmt.Println("writing string to sb")
	n, err := sb.WriteString("This is string builder")
	if err != nil {
		fmt.Println("Error writing to sb ", err)
	}
	fmt.Println("Wrote n line: ", n)
	fmt.Println("--> sb: ", sb.String())

	fmt.Println("Writing byte to sb")
	n, err = sb.Write([]byte(" this is byte"))
	if err != nil {
		fmt.Println("Error writing to sb ", err)
	}
	fmt.Println("Wrote n line: ", n)
	fmt.Println("--> sb: ", sb.String())
	fmt.Println("len of sb: ", sb.Len())
	fmt.Println("cap of sb: ", sb.Cap())
	sb.Reset()
	fmt.Println("--> reset ", sb.String())
	fmt.Println("")
}

func Json() {
	fmt.Println("------------- JSON --------------")
	fmt.Println("")
	type ToDo struct {
		TaskID   int    `json:"task_id"`
		TaskName string `json:"task_name"`
		Status   string `json:"status"`
	}

	task := ToDo{
		TaskID:   1,
		TaskName: "Go revision",
		Status:   "in-progress",
	}

	fmt.Println("converting task obj to json")
	jsontask, err := json.Marshal(&task)
	if err != nil {
		fmt.Println("error converting task obj to json: ", err)
	}
	fmt.Println("json marshal: ", jsontask)
	jsonIndent, err := json.MarshalIndent(&task, "", "\t")

	if err != nil {
		fmt.Println("error converting task obj to indent json: ", err)
	}
	fmt.Println("json marshal: ", string(jsonIndent))

	fmt.Println("converting indent json to obj again")
	task1 := ToDo{}

	if err := json.Unmarshal(jsonIndent, &task1); err != nil {
		fmt.Println("error converting json to obj: ", err)
	}
	fmt.Println("-->", task1)

}
