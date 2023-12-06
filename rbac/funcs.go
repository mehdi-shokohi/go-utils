package rbac

func CheckInArray(arr []string, target string) bool {
	for _,v:=range arr{
		if v==target{
			return true
		}
	}

	return false
}