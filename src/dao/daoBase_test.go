package dao_test

import (
	"reserve_service/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDriver(t *testing.T) {
	assert.NotEmpty(t, dao.GetCommonDb())
}
