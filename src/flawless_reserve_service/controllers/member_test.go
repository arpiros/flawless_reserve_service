package controllers_test

import (
	"flawless_reserve_service/flawless_reserve_service/controllers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMember(t *testing.T) {
	err := controllers.GetMember(nil)
	assert.Empty(t, err)
}
