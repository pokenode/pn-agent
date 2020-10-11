package system

import (
	"fmt"
	"testing"
)

func TestDocker(t *testing.T) {
	fmt.Println("Testing Docker...")
	_ = Docker()
	//PPrint(stats)
}
