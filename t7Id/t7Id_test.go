package t7Id

import (
	"github.com/spf13/viper"
	"testing"
)

func TestNew(t *testing.T) {
	viper.AddConfigPath("../test")

	for i := 0; i < 10; i++ {
		t.Log(New().Generate().Int64())
	}
}
