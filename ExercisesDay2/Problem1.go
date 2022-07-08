package ExercisesDay2

import (
	"encoding/json"
	"fmt"
)

func count(word string, freq map[string]int) {
	done := make(chan bool, len(word))
	for i := 0; i < len(word); i++ {
		i := i
		go func() {
			freq[string(word[i])]++
			done <- true
		}()
	}
	for i := 0; i < len(word); i++ {
		<-done
	}
}

func Problem1() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	freq := make(map[string]int)
	done := make(chan bool, len(words))

	for i := 0; i < len(words); i++ {
		i := i
		go func() {
			count(words[i], freq)
			done <- true
		}()

	}
	for i := 0; i < len(words); i++ {
		<-done
	}

	result, err := json.MarshalIndent(freq, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(result))
	}
}
