package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestThing(t *testing.T) {
	var n Notifier

	a := NotifierMock{}
	a.On("Send", "string").Return(nil)

	n = &a

	assert.Nil(t, Thing(&n))

}

// Mock for notification testing
type NotifierMock struct {
	mock.Mock
}

func (m *NotifierMock) Send(msg string) error {
	m.Called(msg)
	return nil
}
