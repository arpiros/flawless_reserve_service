package dao_test

import (
	"flawless_reserve_service/flawless_reserve_service/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDriver(t *testing.T) {
	assert.NotEmpty(t, dao.GetCommonDb())
}
