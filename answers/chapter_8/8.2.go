/*
构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；先打印所有的饮料，然后打印原名和翻译后的名字。接下来按照英文名排序后再打印出来。
备注：小弟不才，因为饮料翻译比较困难，所以改为了水果
*/
package chapter_8

import (
	"fmt"
	"sort"
)

func main() {
	fruit := map[string]string{"peach": "桃子", "Lemon": "柠檬", "Pear": "梨子"}
	for key, val := range fruit {
		fmt.Printf("%s:%s", key, val)
		fmt.Printf("\n")
	}

	fmt.Printf("\n")

	for key := range fruit {
		fmt.Printf(key)
		fmt.Printf("\n")
	}

	fmt.Printf("\n")

	for _, val := range fruit {
		fmt.Printf(val)
		fmt.Printf("\n")
	}

	fmt.Printf("\n")

	arr := make([]string, len(fruit))
	i := 0
	for key := range fruit {
		arr[i] = key
		i++
	}

	sort.Strings(arr)

	for _, val := range arr {
		fmt.Printf("%s:%s", val, fruit[val])
		fmt.Printf("\n")
	}

}
