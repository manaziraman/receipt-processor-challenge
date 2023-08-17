package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

/**
* TestGenerateUUID tests the generateUUID function to ensure that it returns a valid UUID.
* It asserts that the function does not return an error and that the returned UUID is valid.
*
* @param t *testing.T - The testing object provided by the Go testing framework.
*/
func TestGenerateUUID(t *testing.T) {
	// Call the generateUUID function
	generatedUUID, err := generateUUID()

	// Assert that there was no error
	assert.NoError(t, err, "generateUUID() should not return an error")

	// Assert that the generated UUID is valid
	_, err = uuid.Parse(generatedUUID)
	assert.NoError(t, err, "generateUUID() should return a valid UUID")
}
