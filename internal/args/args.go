package args

func HasArg(args []string, arg string) bool {
	for _, a := range args {
		if a == arg {
			return true
		}
	}
	return false
}
