package ExercisesDay2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func rate(id int, totalRating *int) {
	fmt.Println("Student -", id, "started rating")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	*totalRating += rand.Intn(6)
	fmt.Println("Student -", id, "finished rating")
}

func Problem2() {
	totalStudents := 200
	var totalRating int
	var wg sync.WaitGroup

	for i := 1; i <= totalStudents; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			rate(i, &totalRating)
		}()
	}
	wg.Wait()
	fmt.Println("totalRating:", totalRating)
	averageRating := float32(totalRating) / float32(totalStudents)
	fmt.Println(averageRating)
}
