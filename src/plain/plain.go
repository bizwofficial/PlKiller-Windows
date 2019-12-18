package plain

import "os"

//IfArgs func
//	Usage:	plain.IfArgs()
//	Return:	string typed
//	Info:	True if there's an argument, or false.
func IfArgs() string {
	if len(os.Args) == 2 {
		return "true"
	} else if len(os.Args) == 3 && os.Args[2] == "-v" {
		return "verbosed"
	} else if len(os.Args) == 4 && os.Args[2] == "-im" {
		return "specific"
	} else if len(os.Args) == 5 && os.Args[2] == "-im" && os.Args[4] == "-v" {
		return "verbosed,specific"
	} else {
		return "false"
	}
}

//GetArgs func
//	Usage:	plain.GetArgs()
//	Return:	String typed
//	Info:	Get the argument in shell if it exists or will return a void string.
func GetArgs() (string, bool, string, bool) {
	if IfArgs() == "verbosed" {
		return os.Args[1], true, "", false
	} else if IfArgs() == "true" {
		return os.Args[1], false, "", false
	} else if IfArgs() == "specific" {
		return os.Args[1], false, os.Args[3], true
	} else if IfArgs() == "verbosed,specific" {
		return os.Args[1], true, os.Args[3], true
	} else {
		return "", false, "", true
	}
}

func IsIn(iti string, ast []string) bool {
	symbol := false
	for _, each := range ast {
		if iti == each {
			symbol = true
			break
		}
	}
	return symbol
}
