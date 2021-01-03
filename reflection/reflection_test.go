package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

//Problem: write a function walk(x interface{}, fn func(string))
//which takes a struct x and calls fn for all strings fields found inside
func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Test"},
			[]string{"Test"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"TestName", "TestCity"},
			[]string{"TestName", "TestCity"},
		},
		{
			"Struct with non string fields",
			struct {
				Name string
				Age  int
			}{"TestName", 99},
			[]string{"TestName"},
		},
		{
			"Struct with nested fields",
			Person{
				"TestName",
				Profile{99, "TestCity"},
			},
			[]string{"TestName", "TestCity"},
		},
		{
			"Pointers to things",
			&Person{
				"TestName",
				Profile{99, "TestCity"},
			},
			[]string{"TestName", "TestCity"},
		},
		{
			"Slices",
			[]Profile{
				{33, "Manila"},
				{34, "Batangas"},
			},
			[]string{"Manila", "Batangas"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Manila"},
				{34, "Batangas"},
			},
			[]string{"Manila", "Batangas"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		testMap := map[string]string{
			"33": "Manila",
			"34": "Batangas",
		}

		var got []string
		walk(testMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Manila")
		assertContains(t, got, "Batangas")
	})

	t.Run("with channels", func(t *testing.T) {
		testChannel := make(chan Profile)

		go func() {
			testChannel <- Profile{35, "Cebu"}
			testChannel <- Profile{36, "Dauin"}
			close(testChannel)
		}()

		var got []string
		want := []string{"Cebu", "Dauin"}

		walk(testChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but didn't", haystack, needle)
	}
}
