package interact

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// 清空屏幕
func clear() error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// 停顿
func pause(sec int64) {
	time.Sleep(time.Second * time.Duration(sec))
}

// 读入字符串
func readLine(s *string) (err error) {
	fmt.Scan()
	// 从stdin中取内容直到遇到换行符，停止
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	*s = str[:len(str)-2]

	return nil
}

// 读入整型
func readInt(i *int32) (err error) {
	var s string
	readLine(&s)
	j, err := strconv.Atoi(s)
	*i = int32(j)
	return
}

// 读入整型
func readInt64(i *int64) (err error) {
	var s string
	readLine(&s)
	j, err := strconv.Atoi(s)
	*i = int64(j)
	return
}

// 读入浮点型
func readFloat(i *float32) (err error) {
	var s string
	readLine(&s)
	j, err := strconv.ParseFloat(s, 32)
	*i = float32(j)
	return
}
