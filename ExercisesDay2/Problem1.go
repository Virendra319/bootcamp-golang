package ExercisesDay2

import (
	"encoding/json"
	"fmt"
)

func count(word string, freq map[string]int) {
	for i := 0; i < len(word); i++ {
		freq[string(word[i])]++
	}
}
func Problem1() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	freq := make(map[string]int)

	for i := 0; i < len(words); i++ {
		count(words[i], freq)
	}
	result, err := json.MarshalIndent(freq, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(result))
	}
}
