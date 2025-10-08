package testing_utils


import "testing"

func TestFullNameMethod(t *testing.T) {
	tests := []struct{
		name string
		FirstName string
		LastName string
		result string
	}{
		{
			"first",
			"Sergo",
			"Legend",
			"Sergo Legend",
		},
		{
			"second",
			"Sanya",
			"Ulivir",
			"Sanya Ulivir",
		},
	}


	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user := User{
				FirstName: test.FirstName,
				LastName: test.LastName,
			}

			if fullName := user.FullName(); fullName != test.result {
				t.Errorf("User with expected result %s, got %s", test.result, fullName)
			}
		})
	}
}