package setup

import (
	"fmt"
	"testing"
)

func TestGetLogger(t *testing.T) {
	fmt.Println("Testing logger...")
	lgr := GetLogger()
	lgr.Print("все норм, крч")

}
