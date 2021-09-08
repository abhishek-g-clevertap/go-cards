package main

import "fmt"

func main()  {
	//cards:=newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	fileCards:=newDeckFromFile("my_cards")
	//fileCards.print()
	fileCards.shuffle()
	fileCards.print()

	numbers:=[]int{1,2,3,4,5,6,7,8,9,10};
	for i:=range numbers{
		if i%2 == 0{
			fmt.Println("Event")
		}else{
			fmt.Println("Odd")
		}
	}
}

