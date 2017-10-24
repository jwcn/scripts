package main

/*
列出所有正在运行的进程,并打印每个进程执行的子进程个数
- 如果父进程有一个子进程,就打印 child ,如果多于一个,就打印children ;
- 进程列表要按照数字排序,这样就以 pid 0 开始,依次展示。
*/

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	cmd := exec.Command("/bin/ps", "-e", "-opid,ppid,comm")
	buf, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ppids := []int{}
	res := make(map[int][]int)

	response := strings.Split(string(buf), "\n")
	for k, v := range response {
		if k == 0 || k == len(response)-1 {
			continue
		}

		line := strings.Fields(v)
		pid, _ := strconv.Atoi(line[0])
		ppid, _ := strconv.Atoi(line[1])
		res[ppid] = append(res[ppid], pid)

	}

	for ppid, _ := range res {
		ppids = append(ppids, ppid)
	}
	sort.Ints(ppids)

	for _, ppid := range ppids {
		pv := res[ppid]
		var ch string
		l := len(pv)
		if l == 1 {
			ch = "child"
		} else {
			ch = "children"

		}

		fmt.Printf("Pid %d has %d %s: %d \n", ppid, l, ch, pv)
	}

}
