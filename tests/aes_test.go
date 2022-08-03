package tests

import (
	utils "github.com/titrxw/smart-home-server/app/Utils"
	"testing"
)

func TestAes(t *testing.T) {
	t.Run("testAes", func(t *testing.T) {
		var data = "qwer"
		var key = "12345678912345678909876543212345"
		result, _ := utils.Encrypt(data, key)

		result1, _ := utils.Decrypt(result, key)

		if result1 == data {
			t.Skipped()
		}
		t.Failed()
	})
}
