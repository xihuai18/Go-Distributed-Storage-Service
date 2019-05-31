// 命令行实用工具
package main

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
	"os/exec"
	"crypto/sha256"
	"io"
	"os"
	"encoding/base64"
	"time"
	"lib/trie"
	"strings"
)

func hashSHA256File(filePath string) (string, error){
    var hashValue string
    file, err := os.Open(filePath)
    if err != nil {
        return hashValue, err
    }
    defer file.Close()
    hash := sha256.New()
    if _, err := io.Copy(hash, file); err != nil {
        return hashValue,  err
    }
    hashInBytes := hash.Sum(nil)
    hashValue = base64.StdEncoding.EncodeToString(hashInBytes)
    return hashValue, nil
}

// the questions to ask
var simpleQs = []*survey.Question{
	{
		Name: "action",
		Prompt: &survey.Select{
			Message: "Select an action: ",
			Options: []string{"upload", "download", "delete", "query", "look up", "exit"},
		},
		Validate:  survey.Required,
	},
}
var simpleQs2 = []*survey.Question{
	{
		Name:     "fileName",
		Prompt:   &survey.Input{
			Message: "Input the file name or the prefix to look up: "},
		Validate: survey.Required,
	},
}

func main() {
	for {

		answers := struct {
			Action  string
			FileName string
			}{}
			
		err := survey.Ask(simpleQs, &answers)
		
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if answers.Action == "exit" {
			return 
		}
		err = survey.Ask(simpleQs2, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		
		fmt.Printf("%s %s...\n", answers.Action, answers.FileName)
		
		switch answers.Action {
		case "upload":
			hash, err := hashSHA256File(answers.FileName)
			if err != nil{
				fmt.Println(err.Error())
				return 
			}
			cmd := exec.Command("curl", fmt.Sprintf(`127.0.2.1:12345/objects/%s`, answers.FileName), "-XPUT", "-T", fmt.Sprintf("%s", answers.FileName), "-H", fmt.Sprintf("Digest: SHA-256=%s", hash))
			out, err := cmd.Output()
			if err != nil{
				fmt.Println(err.Error())
				return 
			}
			fmt.Printf(string(out))
			
		case "download":
			cmd := exec.Command("curl", fmt.Sprintf("127.0.2.1:12345/objects/%s", answers.FileName), "-XGET", "-o", fmt.Sprintf("%s", answers.FileName))
			out, err := cmd.Output()
			if err != nil{
				fmt.Println(err.Error())
				return 
			}
			fmt.Printf(string(out))
			
		case "delete":
			flag := false
			prompt := &survey.Confirm{
				Message: fmt.Sprintf("Are you sure to delete %s?", answers.FileName),
			}
			survey.AskOne(prompt, &flag, nil)
			if flag == true {
				cmd := exec.Command("curl", fmt.Sprintf("127.0.2.1:12345/objects/%s", answers.FileName), "-XDELETE")
				out, err := cmd.Output()
				if err != nil{
					fmt.Println(err.Error())
					return 
				}
				fmt.Printf(string(out))
			}
			
		case "query":
			cmd := exec.Command("curl", fmt.Sprintf("127.0.2.1:12345/versions/%s", answers.FileName))
			out, err := cmd.Output()
			if err != nil{
				fmt.Println(err.Error())
				return 
			}
			fmt.Printf(string(out)+"\n")

		case "look up":
			cmd := exec.Command("curl", fmt.Sprintf("127.0.2.1:12345/lookup/"))
			out_bytes, err := cmd.Output()
			if err != nil{
				fmt.Println(err.Error())
				return 
			}
			objs := strings.Split(string(out_bytes), "\n")
			tree := trie.NewTrie()
			for i := range(objs) {
				tree.Insert(objs[i])
			}
			suffix := tree.AllStartWith(answers.FileName)
			fmt.Println("look up result: \n"+strings.Join(suffix, "\n")+"\n")

		default:
			fmt.Println("Error action")
		}
		time.Sleep(1 * time.Second)
	}
}