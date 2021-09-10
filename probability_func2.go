package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "p_data.txt"
	for j := 1; j <= 2000; j++ {
		n := j
		sum := float64(0)
		var bounds string
		count := 0
		isFirst := true
		for i := 1; i <= n; i++ {
			x0 := (float64(i - 1)) * 100 / float64(n)
			x1 := (float64(i)) * 100 / float64(n)
			y0 := 0.2325*(x0)*(x0)*(x0)*(x0)/float64(4) - float64(35)*(x0)*(x0)*(x0)/float64(3) + 177.2*x0*x0 + float64(85540)*x0
			y1 := 0.2325*(x1)*(x1)*(x1)*(x1)/float64(4) - float64(35)*(x1)*(x1)*(x1)/float64(3) + 177.2*x1*x1 + float64(85540)*x1
			p := y1 - y0
			sum = sum + p
			if len(bounds) == 0 {
				if isFirst {
					bounds = fmt.Sprintf("0,%d", int(sum))
				} else {
					bounds = fmt.Sprintf("%d", int(sum))
				}
			} else {
				bounds = fmt.Sprintf("%s,%d", bounds, int(sum))
			}
			count++
			if count == 100 {
				appendFile(fileName, fmt.Sprintf("%d|%s", j, bounds))

				count = 0
				bounds = ""
				isFirst = false
			}
		}
		appendFile(fileName, fmt.Sprintf("%d|%s", j, bounds))
	}
	yt := 0.2325*(100)*(100)*(100)*(100)/float64(4) - float64(35)*(100)*(100)*(100)/float64(3) + 177.2*100*100 + float64(85540)*100
	fmt.Println(yt)
}

func appendFile(fileName string, str string) error {
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(fd)
	fmt.Fprintln(w, str)
	return w.Flush()
}
