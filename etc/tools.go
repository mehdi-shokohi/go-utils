package etc
func containsString(s []string, v string) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func CheckInArray(arr []string, target string) bool {
	for _,v:=range arr{
		if v==target{
			return true
		}
	}

	return false
}