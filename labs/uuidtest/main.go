package main

import (
	"fmt"

	"github.com/satori/go.uuid"
	// "github.com/google/uuid"
)

func main() {
	// fmt.Println("vim-go")
	// id, _ := uuid.NewV4()

	// domainstring := []byte("aozsky.com")

	// uid := uint8(os.Getuid)

	// u2, err := uuid.NewV2([]byte(domain))
	// fmt.Printf("uuid.NewV4() id: %s\n", id)

	u2, _ := uuid.NewV2(uuid.DomainGroup)
	fmt.Printf("%s\n", u2)
	// u5, _ := uuid.FromString("聂倩倩")
	// u5new := uuid.NewV5(u5, "aozsky.com")
	// fmt.Printf("NewV5 with uuid.FromString() u5: %s\n", u5new)
}
