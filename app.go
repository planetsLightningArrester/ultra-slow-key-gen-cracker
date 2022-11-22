package main

import (
	"math/rand"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.KeySleep = 0

	// Possible key-gen values (not all letters are here, like O)
	strings := []string{"A", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	key_gen_size := 24

	for {
		key := ""
		for i := 0; i < key_gen_size; i++ {
			key += strings[rand.Intn(len(strings))]
		}
		robotgo.WriteAll(key)
		robotgo.KeyToggle("lctrl", "down")
		robotgo.KeyToggle("v", "down")
		robotgo.KeyToggle("lctrl", "up")
		robotgo.KeyToggle("v", "up")

		robotgo.MilliSleep(200)
		robotgo.KeyTap("enter")
		robotgo.KeyToggle("lctrl", "down")
		robotgo.KeyToggle("a", "down")
		robotgo.KeyToggle("lctrl", "up")
		robotgo.KeyToggle("a", "up")
		robotgo.KeyTap("backspace")
	}
}
