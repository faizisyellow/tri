package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	Position int
	Done     bool
}

// ByPri implements sort.Interface For []Item based on
// the priority & position field.
type ByPri []Item

func SaveItems(filename string, items []Item) error {

	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {

	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, fmt.Errorf("error when reading file: %v", err)
	}

	var items []Item

	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i := range items {

		items[i].Position = i + 1
	}

	return items, nil
}

func (i *Item) SetPriority(pri int) {

	switch pri {
	case 1:
		i.Priority = 1
	case 2:
		i.Priority = 2
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {

	if i.Priority == 1 {
		return "(1)"
	}

	if i.Priority == 3 {
		return "(3)"
	}

	return " "
}

func (i *Item) PrettyDone() string {

	if i.Done {
		return "X"
	}

	return ""
}

func (i *Item) Label() string {

	return strconv.Itoa(i.Position) + "."
}

func (s ByPri) Len() int {

	return len(s)
}

func (s ByPri) Swap(i, j int) {

	s[i], s[j] = s[j], s[i]
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s ByPri) Less(i, j int) bool {

	if s[i].Done != s[j].Done {

		return s[i].Done
	}

	if s[i].Priority == s[j].Priority {
		return s[i].Position < s[j].Position
	}

	return s[i].Priority < s[j].Priority
}
