package main

import (
	"fmt"
	"math"
	"time"
)

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

//explict & implicit return ---functions
func yearUntilEvent (age int)(
yearUntilAdult,  yearUntilLegal,  yearUntilUni int){
	yearUntilAdult = 18 - age
	if yearUntilAdult < 0{
		yearUntilAdult = 0
	}
yearUntilLegal = 21 - age
	if yearUntilLegal < 0{
		yearUntilLegal = 0
	}
	yearUntilUni = 16 - age
	if yearUntilUni < 0{
		yearUntilUni = 0
	}
return
}

func testAge(age int){
	fmt.Println("Age:", age)
	yearUntilAdult, yearUntilLegal, yearUntilUni := yearUntilEvent(age);
	fmt.Println("you're an adult in", yearUntilAdult, "years")
	fmt.Println("you're legal until", yearUntilLegal, "years")
	fmt.Println("youre not in uni until", yearUntilUni, "years")
	fmt.Println("==================")
}

//explicit returns
func cordinates() (x, y int){
	return x, y
}


//explicit returns

func swap (string1, string2 string) (string, string){
return string2, string1
}

func PrintSwap (){
	a, b := swap("hello", "world")
	fmt.Println(a,b)
}

func week (){
	today := time.Now().Weekday()
	// fmt.Println(today)

	switch time.Saturday{
	case today + 0:
		fmt.Println("today")
	case today + 1 :
		fmt.Println("tomorrow")
	case today - 2 :
		fmt.Println("in two days")
	default:
	fmt.Println("too far away")
	}

	fmt.Println(today)
}

func getTime (){
	t:=time.Now()
	// fmt.Println(t)
	switch {
	case t.Hour() > 12 && t.Hour() < 17:
		fmt.Println("good afternooon")
			case t.Hour() >= 17 && t.Hour()< 23:
		fmt.Println("good evening")
	case t.Hour() < 12:
		fmt.Println("good morning")
	default:
		fmt.Println("good whatever")
	}

}
//when defer is used, the value or function on which it was used is first checked and after, is passses into a stack, then it executes last after all the file before or after it in the same code block is exected - it's used in pushing the execution of a funtion to last statement

func deferTest (){
	fmt.Println("counting")

	for h := 0; h <= 10; h++{
		defer fmt.Println(h)
	}


	fmt.Println("done")
}

var num1, num2 int = 10,12
var sqrt = math.Sqrt(float64(num1*num1 + num2*num2))
var sqrtUint uint = uint(sqrt)
var namePrnt string = "ola" //works outside function
// var namePrint := "hello"works only withing a function block

// for loop
func loop(){
	loopSum := 0;
	for i :=0; i < 10; i++{
		loopSum += 1
		}
		fmt.Println(loopSum)
}


// reference types : pointers

func point(){
	point1, point2 := 90, 100

	l := &point1
	m := &point2
	fmt.Println(*l,*m)

	*l = *l *2
	*m = *m / 10
	fmt.Println("new value of point 1 is ",point1, "   new value of point2 is", point2)
}


// Agreegate dataTypes: array and structs

// structs can take multiple data types but array can't

type vertex struct {
	struct1 string
	struct2 int
	struct3 bool
}

func structOutput () {
	// fmt.Println(vertex{"bukola", 10, true})
	// v := vertex{struct1 :"bukola", struct2: 10,struct3:  true}
	v := vertex{"bukola",  10, true}
	v.struct1 = "ola"
	v.struct2 = v.struct2*2

	fmt.Println(v)
}



func structPointers (){
		vp  := vertex{"bukola", 10, true}
		structPoint := &vp
		structPoint.struct1 = "something"
		structPoint.struct2 = vp.struct2 * 10
		fmt.Println(vp)
		fmt.Println(structPoint)
}


// arrays

func myArrays (){
	var sArray [2] string
	sArray[0] = "hello"
	sArray[1]= "world"
	fmt.Println(sArray)
	fmt.Println(sArray[0], sArray[1])

	evenNum := [6] int {3,5,7,11,13,17}
	fmt.Println(evenNum   )
}

//slices
func sliceArray (){
	herNames := [6] string {"nola", "damola", "magret", "fola", "gbade", "simi"}
	fmt.Println(herNames)
	var sliceTest []string = herNames[2:5] //show only array between the index of 2 and 4
	fmt.Println(sliceTest)
}

func sliceReference (){
	referenceArray := [5] string {"john", "wesley", "Haggin", "Myles", "Abraham"}
	fmt.Println(referenceArray)

	var slice1 = referenceArray[2:4]
	var slice2 = referenceArray[1:5]
	fmt.Println(slice1, slice2 )

	slice2[1] = "something esle"
	fmt.Println(slice1, slice2 )
		fmt.Println(referenceArray)

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
	testAge(16)
	testAge(12)
	testAge(0)
	cordinates()
	PrintSwap()
	println(sqrt, sqrtUint)
	println( sqrtUint)
	loop()
	week()
	getTime()
	deferTest()
	 point()
	 structOutput()
	 structPointers()
	 myArrays()
	 sliceArray()
	 sliceReference()
}
