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
			y0 := (x0)*(x0)*(x0)*(x0)/float64(24) - (x0)*(x0)*(x0)*float64(25)/float64(3) + float64(87500)*x0
			y1 := (x1)*(x1)*(x1)*(x1)/float64(24) - (x1)*(x1)*(x1)*float64(25)/float64(3) + float64(87500)*x1
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
	yt := (100)*(100)*(100)*(100)/24 - (100)*(100)*(100)*25/3 + 87500*100
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
