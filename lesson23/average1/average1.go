package main

import "fmt"

type report struct {
	sol         int
	location    location
	temperature temperature
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// func (r report) average() celsius {
// 	return r.temperature.average()
// }

func main() {
	t := temperature{high: -1.0, low: -78.0}
	fmt.Printf("average %vº C\n", t.average())

	report := report{sol: 15, temperature: t}
	fmt.Printf("average %vº C\n", report.temperature.average())
	// 关于方法转发：
	// 在 Go 语言中，只有通过嵌入字段（anonymous fields）才能实现方法的自动转发
	// 而在这个代码中，temperature 是作为命名字段存在的（第 8 行），不是嵌入字段
	// fmt.Printf("average %vº C\n", report.average())
}
