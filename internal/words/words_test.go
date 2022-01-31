package words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWords(t *testing.T) {
	words := []string{"perky"}
	assert.Equal(t, 1, len(Has(words, "rk")))
	assert.Equal(t, 1, len(Excludes(words, "ans")))
	assert.Equal(t, 1, len(Orders(words, "???k?")))
	assert.Equal(t, 1, len(Orders(Excludes(Has(words, "rk"), "ans"), "???k?")))

}

func TestNotOrders(t *testing.T) {
	words := []string{"perky", "sugar", "could"}
	notOrders := []string{"???k?"}
	assert.Equal(t, 2, len(NotOrders(words, notOrders)))
	assert.Equal(t, false, doesntMatch("perky", "???k?"))
}
