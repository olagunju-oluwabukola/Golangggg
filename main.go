package main

import "fmt"

// variable declarations
func variables() {
	var name string
	var cost float32
	var num int
	var ask bool

	_ = name
	_ = cost
	_ = num
	_ = ask

	averageOpenRate, displayMessage, temperaturInit := 0.23, "is the average open rate", 58
	tempFloat := float64(temperaturInit)
	const premuimPlan = "plan a"

	fmt.Println(averageOpenRate, displayMessage)
	fmt.Println(tempFloat)
	fmt.Println("premium plan is", premuimPlan)
}

// String concatenation
func concact(s1 string, s2 string) string {
	return s1 + s2
}

// Conditionals
func conditionals() {
	const height = 4
	if height > 4 {
		fmt.Println("you're super tall")
	} else if height < 4 {
		fmt.Println("you're short")
	} else {
		fmt.Println("height is probably =", height)
	}

	messageLength := 10
	maxMessageLength := 20
	name := "bukola"

	fmt.Println(name, "is trying to send a message length of", messageLength, "but the max length is", maxMessageLength)

	if messageLength < maxMessageLength {
		fmt.Println("response: message sent")
	} else if messageLength > maxMessageLength {
		fmt.Println("response: message is too long")
	} else {
		fmt.Println("try again")
	}
}

// Subtraction
func sub(x int, y int) int {
	return x - y
}

// Addition
func add(u, w int) int {
	return u + w
}

// Examples
func examples() {
	// length
	myName := "bukola"
	if len(myName) > 4 {
		fmt.Println("my name is greater than 4")
	}

	// Seconds in an hour cal
	const numberOfMinutes = 60
	const numberOfSeconds = 60
	const numberofSecondsInHour = numberOfMinutes * numberOfSeconds
	fmt.Println("there are", numberofSecondsInHour, "seconds in an hour")

	//
	const name = "bukola"
	const openRate = 30.5
	msg := fmt.Sprintf("hi %s your open rate is %.2f percent", name, openRate)
	fmt.Println(msg)
	fmt.Println("hi", name, " your open rate is", openRate)

	//
	congrats := "Happy birthdasy"
	fmt.Println(congrats)
}
// reassigning values
func sentMessages (){
toBeSent := 400
const sent = 20
toBeSent = calSent(toBeSent, sent)
fmt.Println("you have", toBeSent, "messages to be sent")

}

func calSent(toBeSent, sent int) int{
	toBeSent = toBeSent - sent
	return toBeSent
}

// ignoring unused variables --are ignored with underscore(-)

func nameInput (){
	firstName, _ := getNames()
	fmt.Println("hello", firstName)
}

func getNames () (string, string){
	return "Bukola", "Olagunju"
}
func main() {
	examples()
	conditionals()
	variables()
	fmt.Println(concact("hey", "there "))
	fmt.Println(sub(10, 2))
	fmt.Println(add(10, 2))
	sentMessages()
	nameInput()
}
