package plain

import "os"

//IfArgs func
//	Usage:	plain.IfArgs()
//	Return:	string typed
//	Info:	True if there's an argument, or false.
func IfArgs() string {
	if len(os.Args) == 2 {
		return "true"
	} else if len(os.Args)==3 && os.Args[2]=="-v"{
		return "verbosed"
	} else {
		return "false"
	}
}

//GetArgs func
//	Usage:	plain.GetArgs()
//	Return:	String typed
//	Info:	Get the argument in shell if it exists or will return a void string.
func GetArgs() (string,bool) {
	if IfArgs()=="verbosed"{
		return os.Args[1],true
	}else if IfArgs()=="true"{
		return os.Args[1],false
	}else{
		return "",false
	}
}

func IsIn(iti string,ast []string) bool {
	symbol:=false
	for _,each:=range ast{
		if iti==each{
			symbol=true
			break
		}
	}
	return symbol
}