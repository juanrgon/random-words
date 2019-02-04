package main

import "testing"

func Test_load_words(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := loadWords()
			first := got[0]
			last := got[len(got)-1]
			if first != "home" {
				t.Errorf("Expected the first common word to be home")
			}
			if last != "happier" {
				t.Errorf("Expected the last common word to be happier")
			}
		})
	}
}
