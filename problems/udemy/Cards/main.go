package main

func main() {

	cards := getDeck() //function will excecute and return a value of type deck , cards will be hold value of type deck so it can access print function

	//cards.saveToFile("testFile")

	//cards := newDeckFormFile("testFile")
	cards.shuffle()
	cards.print()

}
