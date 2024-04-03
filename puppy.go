package puppy

import "github.com/ezraorbit/dog"

func Bark() string {
	return "Yah boy!"
}

func Barks() string {
	return "Woof! Woof! Okay Bark! Bark!"
}

func grownUp() string {
	return "I'm all grown up! BAAAAAAARK!!!!"
}

func BigBark() string {
	return dog.WhenGrownUps(grownUp())
}
