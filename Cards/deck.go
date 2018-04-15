package main
import "fmt"
type deck []string

func createDeck() deck{
	
	cards := deck{}
	cardType := []string {"Spades","Diamonds","Hearts","Clubs"}
	cardValue  := []string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","King","Queen","Jack"}

	for _,kind := range cardType {
		for _,value := range cardValue {
			cards = append(cards,value + " of "+ kind)
		}
	}
	return cards
}

func (d deck) print(){
	for _,card := range d {
		fmt.Println(card)
	}
}
