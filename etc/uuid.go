package etc

import "github.com/nu7hatch/gouuid"



func UUIDGen() string {
	u, err := uuid.NewV4()
	if err!=nil{
		panic(err)
	}
	return u.String()
}
