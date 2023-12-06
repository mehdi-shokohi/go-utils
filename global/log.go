package globUtils
import(
	"fmt"
)
const (
	ErrAlert = "alert"
	ErrInfo = "info"
)
func Logger(level ,message string, args ...interface{}){
	fmt.Println(level, message)
}

func LoggerInfo(message string, args ...interface{}){
	 Logger(ErrInfo, message, args)
}

func LoggerAlert(message string, args ...interface{}){
	Logger(ErrAlert, message, args)
}