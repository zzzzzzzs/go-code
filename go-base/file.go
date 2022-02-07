package go_base

import (
	"fmt"
	"io/ioutil"
	"os"
)

//读取到file中，再利用ioutil将file直接读取到[]byte中, 这是最优
func Read3()  (string){
	f, err := os.Open("file/test")
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}

func main() {
	Read3()
}