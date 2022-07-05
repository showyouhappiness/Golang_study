package structs

import (
	"fmt"
	"testing"
	"time"
)

type Address struct {
	Street           string
	HouseNumber      uint32
	HouseNumberAddOn string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

type VCard struct {
	name      string
	BirtDate  time.Time
	Photo     string
	Addresses map[string]*Address
}

func TestVCard(t *testing.T) {
	addr1 := &Address{"Elfenstraat", 12, "", "", "2600", "Mechelen", "België"}
	addr2 := &Address{"Heideland", 28, "", "", "2640", "Mortsel", "België"}
	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2
	birthdt := time.Date(1996, 1, 1, 15, 4, 5, 0, time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := &VCard{"liyongjun", birthdt, photo, addrs}
	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Printf("My Addresses are:\n %v\n %v", addr1, addr2)
}
