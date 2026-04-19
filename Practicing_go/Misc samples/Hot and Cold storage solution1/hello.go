package HCS

import (
	"fmt"
	"math/rand"
)

// // implement a service, including:
// -  hot storage and cold storage
// - storage as a map[id][string,string]
// - operations: store(), retrieve()
// - hot storage offload map values to cold one in 100ms

func main() {
	firstNamesList := [...]string{"Emma", "Liam", "Olivia", "Noah", "Ava", "William", "Sophia", "James", "Isabella", "Oliver", "Charlotte", "Benjamin", "Amelia", "Elijah", "Mia", "Lucas", "Harper", "Mason", "Evelyn", "Logan"}
	lastNamesList := [...]string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Wilson"}

	result := func() string {
		if 0 == NewStorage() {
			return "success"
		}
		return "failure"
	}()
	fmt.Println("Creating the storage:", result)
	// PrintTheStorageContent()
	N := 5
	fmt.Print("Storing ", N, " items there...")
	result = func(itemsNumber int) string {
		// fmt.Println("N:", itemsNumber)
		for range itemsNumber {
			firstName := firstNamesList[rand.Intn(len(firstNamesList))]
			lastName := lastNamesList[rand.Intn(len(lastNamesList))]
			// fmt.Println(firstName, lastName)
			if 0 != Store(FullName{FirstName: firstName, LastName: lastName}) {
				return "failure"
			}
		}
		return "success"
	}(N)
	fmt.Println("", result)
	// PrintTheStorageContent()

	dbgCacheInvalidationEmu()
	PrintTheStorageContent()
	fmt.Println("Storage size retrieved:", StorageSize(), "items")

	sampleValue := FullName{"John", "Smith"}
	Update(2, sampleValue)
	PrintTheStorageContent()

}
