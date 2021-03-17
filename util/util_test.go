package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestThing(t *testing.T) {
	var n Notifier

	a := &NotifierMock{}
	a.On("Send", "string").Return(nil)
	n = a
	assert.Nil(t, Thing(&n))

}

func TestThing2(t *testing.T) {
	a := &NotifierMock{}
	a.On("Send", "string").Return(nil)
	assert.Nil(t, Thing2(a))
}

// Mock for notification testing
type NotifierMock struct {
	mock.Mock
}

func (m *NotifierMock) Send(msg string) error {
	m.Called(msg)
	return nil
}
