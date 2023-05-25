package designpatterns

import (
	"fmt"
	"os"
	"strings"

	"enkya.org/playground/utils"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
}

func (j *Journal) RemoveEntry() {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "")
}

func (j *Journal) Save(filename string) {
	// ...
	_ = os.WriteFile(filename, []byte(j.String()), utils.WritePermission)
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	// ...
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), utils.WritePermission)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), utils.WritePermission)
}

func DesignPatterns() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug.")
	fmt.Println(j.String())

	SaveToFile(&j, "journal.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
