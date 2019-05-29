package helpers

import "testing"

func TestContainsOr(t *testing.T) {

	type demo struct {
		Word     string
		Search   []string
		Expected bool
	}

	dData := []demo{
		demo{Word: "ära", Search: []string{"är", "ar"}, Expected: true},
		demo{Word: "Supercalifragidistico", Search: []string{"distico", "espialidoso"}, Expected: true},
		demo{Word: "Supercalifragidistico", Search: []string{"ditico", "califag"}, Expected: false},
		demo{Word: "visual code", Search: []string{"visual", "code"}, Expected: true},
		demo{Word: "visual code", Search: []string{"Visual", "Code"}, Expected: false},
	}

	var result bool
	for _, test := range dData {
		result = ContainsOr(test.Word, test.Search)
		if result != test.Expected {
			t.Errorf("Expected %v, got %v on '%s' - '%s'", test.Expected, result, test.Word, test.Search)
		}
	}
}

func TestStartsWith(t *testing.T) {

	type demo struct {
		Word     string
		Search   string
		Expected bool
	}

	dData := []demo{
		demo{Word: "ära", Search: "är", Expected: true},
		demo{Word: "ära", Search: "ar", Expected: false},
		demo{Word: "ära", Search: "ärab", Expected: false},
		demo{Word: " ära", Search: "är", Expected: false},
		demo{Word: "Hola", Search: "ho", Expected: false},
		demo{Word: "Hola", Search: "Ho", Expected: true},
		demo{Word: "¡Hola", Search: "¡Ho", Expected: true},
	}

	var result bool
	for _, test := range dData {
		result = StartsWith(test.Word, test.Search)
		if result != test.Expected {
			t.Errorf("Expected %v, got %v on '%s' - '%s'", test.Expected, result, test.Word, test.Search)
		}
	}
}
