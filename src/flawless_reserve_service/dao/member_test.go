package dao_test

import (
	"flawless_reserve_service/flawless_reserve_service/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	member := dao.GetMember(1)
	assert.Equal(t, member.Name, "test")
}
