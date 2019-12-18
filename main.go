package main

import (
	"dispatch"
	"fmt"
	"plain"
)

func main() {
	filename, ifVerbosed, process, ifProcess := plain.GetArgs()
	if filename == "" {
		prompt()
		return
	}
	cptlist, err := plain.GetFile(filename)
	if err != nil {
		prompt()
		return
	}
	if ifVerbosed {
		reallist := dispatch.KillFromList(cptlist, ifProcess, process)
		var act_kil int = 0
		var not_kil int = 0
		actually_killed := ""
		not_killed := ""
		for _, each := range cptlist {
			if plain.IsIn(each, reallist) {
				actually_killed += " "
				actually_killed += each
				act_kil++
			} else {
				not_killed += " "
				not_killed += each
				not_kil++
			}
		}
		fmt.Printf("%d of %d killed.\n\nKilled devices: %s\nFailed devices: %s\n\n", act_kil, act_kil+not_kil, actually_killed, not_killed)
	} else {
		dispatch.KillFromList(cptlist, ifProcess, process)
	}
	fmt.Println("Success.")
	return
}

func prompt() {
	fmt.Println("Plkiller 2.1 for Windows")
	fmt.Println("\nUsage: plkiller <Configuration File Name> [-im <process name>] [-v]")
	fmt.Println("\nThis program is for volume remote controlling. For more details, visit https://github.com/bizwofficial/plkiller-windows")
	return
}
