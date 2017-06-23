/*
创建一个 map 来保存每周 7 天的名字，将它们打印出来并且测试是否存在 Tuesday 和 Hollyday
 */
package main

func main() {
	week := map[string]string{
		"one":"Sunday",
		"two":"Monday",
		"three":"Tuesday",
		"four":"Wednesday",
		"five":"Thursday",
		"six":"Friday",
		"seven":"Saturday",
	}
	for _,x := range week{
		print(x)
		print("\n")
	}

	print("\n")

	for _,ok  := range week{
		if ok == "Tuesday"{
			print("Tuesday is find in map of week")
			print("\n")
		}

		if ok == "Hollyday"{
			print("Hollyday is find in map of week")
		}
	}

}
