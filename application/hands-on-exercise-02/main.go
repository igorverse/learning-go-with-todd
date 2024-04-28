package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	bs := []byte(s)
	people := []Person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	for i, person := range people {
		fmt.Println("#", i)
		fmt.Println("Name: ", person.First)
		fmt.Println("Last name: ", person.Last)
		fmt.Println("Age: ", person.Age)
		fmt.Println("Sayings: ", person.Sayings)
		fmt.Println("---------")
	}
}
