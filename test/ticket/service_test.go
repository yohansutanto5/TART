package user_test

import (
	"app/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert_Normal_1(t *testing.T) {
	tickets := []*model.Ticket{
		{},
	}

	for _, ticket := range tickets {

		err := ticketService.Insert(ticket)
		assert.Nil(t, err)
	}

}

func TestGetList_Normal_1(t *testing.T) {
	users, _ := ticketService.GetList()
	fmt.Println(len(users))
	assert.Greater(t, len(users), 2)
}
