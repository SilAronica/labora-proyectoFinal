package services

import (
	"testing"
)

func TestInsertLogInDb(t *testing.T) { //arreglo de todos los testCases.
	testCases := []struct {
		Name          string
		ExpectedError error
	}{
		{
			Name:          "Insert Log In DB",
			ExpectedError: nil,
		},
	}

	for i := range testCases { //ciclo for.
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel() //Subtarea se va a ejecutar de forma paralela.

		})
	}

}
