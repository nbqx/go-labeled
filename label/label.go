package label

import (
	"strings"
	"strconv"
	"os"
	"os/exec"
	"syscall"
	"errors"
	"sort"

	"code.google.com/p/go.text/unicode/norm"
)

type Label struct {
	colors map[string] int
}

func (self *Label) List() string {
	keys := make([]string,0,len(self.colors))
	for k := range self.colors {
		keys = append(keys,k)
	}
	sort.Strings(keys)
	return strings.Join(keys,"\n")
}

func (self *Label) getColorNum(name string) (string, error){
	if n,ok := self.colors[name]; ok {
		color := strconv.Itoa(n)
		return color, nil
	}else{
		return "", errors.New("ColorNotFoundError")
	}
}

func (self *Label) Get(name string) (string, error) {
	name = strings.ToLower(name)
	if color,err := self.getColorNum(name); err == nil {
		cmd := exec.Command("mdfind","kMDItemFSLabel","==",color)
		cmd.Stderr = os.Stderr
		buf,err := cmd.Output()

		if exitError,ok := err.(*exec.ExitError); ok {
			if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
				if waitStatus.ExitStatus() == 1 {
					return "", exitError
				}
			}
		}

		// NFD->NFC
		return norm.NFC.String(string(buf)), nil

	}else{
		if err.Error()=="ColorNotFoundError" {
			colorList := "Label Colors are: \n"+self.List()+"\n"
			return colorList, nil
		}else{
			return "", err
		}
	}
}

func NewLabel() *Label {
	return &Label{colors: map[string] int {
		"grey": 1,
		"gray": 1,
		"green": 2,
		"purple": 3,
		"blue": 4,
		"yellow": 5,
		"red": 6,
		"orange": 7,
	}}
}
