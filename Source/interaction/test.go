package main

import (
	"io/ioutil"
	"os/exec"
	"fmt"
)

func main()  {
	conts := `abcdefsadfsda\n
	sdafsdfsadfsa\n
	sdafasdfsadfsad\n
	sdafasdfasdfasdfsadfsadfasd\n`
	cmd := exec.Command("less")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stdin.Write([]byte(conts))
	stdin.Close()
	cmd.Start()
	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()
	cmd.Wait()
	fmt.Println(string(out_bytes))
	return 
}