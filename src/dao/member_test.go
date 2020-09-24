package dao_test

import (
	"reserve_service/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertMember(t *testing.T) {
	err := dao.InsertMember(1, "test")
	assert.Equal(t, err, nil)
}

func TestGet(t *testing.T) {
	err, member := dao.GetMember(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, member.Name, "test")
}
