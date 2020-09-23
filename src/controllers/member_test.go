package controllers_test

import (
	"reserve_service/controllers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMember(t *testing.T) {
	err := controllers.GetMember(nil)
	assert.Empty(t, err)
}
