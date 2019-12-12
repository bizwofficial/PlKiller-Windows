package killlib

import (
	"os/exec"
	"sync"
)

func Kill(cpt_name string,flag bool,exe_name string,wg *sync.WaitGroup,state *[]string) {
	cmd:=new(exec.Cmd)
	if flag==false {
		cmd=exec.Command("shutdown","-s","-m","\\\\"+cpt_name,"-t","00")
	}else{
		cmd=exec.Command("taskkill","/S",cpt_name,"/IM",exe_name,"/F")
	}
	cmd.Run()
	*state = append(*state,cpt_name)
	defer wg.Done()
	return
}