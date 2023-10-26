package lib

func GetArg(args []string, index int) string {
	if len(args) > index {
		return args[index]
	} else {
		return ""
	}
}
