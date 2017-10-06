package environ

import (
	"testing"

	"github.com/drone/drone/model"
)

func TestStaticParsing(t *testing.T) {
	expected := [][]string{
		[]string{"PLUGIN_MIRROR", "localhost:5000"},
	}

	service := NewStatic([]string{"PLUGIN_MIRROR=localhost:5000"})
	actual, err := service.EnvironList(&model.Repo{})

	if err != nil {
		t.Error(err)
	}

	if len(actual) != len(expected) {
		t.Fatal("The service should have returned %d values, it returned %s values", len(expected), len(actual))
	}

	for i, ex := range expected {
		if actual[i].Name != ex[0] || actual[i].Value != ex[1] {
			t.Fatal("Actual doesn't equal expected")
		}
	}
}
