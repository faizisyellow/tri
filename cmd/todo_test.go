package cmd

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"faizisyellow.com/tri/todo"
)

func TestTodoSort(t *testing.T) {

	t.Run("should order ascending by item's done and priority", func(t *testing.T) {

		items := []todo.Item{
			{
				Text:     "todos 1",
				Priority: 1,
				Position: 1,
				Done:     true,
			},
			{
				Text:     "todos 2",
				Priority: 1,
				Position: 2,
				Done:     false,
			},
			{
				Text:     "todos 3",
				Priority: 2,
				Position: 3,
				Done:     true,
			},
			{
				Text:     "todos 4",
				Priority: 3,
				Position: 4,
				Done:     true,
			},
			{
				Text:     "todos 5",
				Priority: 1,
				Position: 5,
				Done:     false,
			},
			{
				Text:     "todos 6",
				Priority: 1,
				Position: 6,
				Done:     true,
			},
		}

		expected := []todo.Item{
			{
				Text:     "todos 1",
				Priority: 1,
				Position: 1,
				Done:     true,
			},
			{
				Text:     "todos 6",
				Priority: 1,
				Position: 6,
				Done:     true,
			},
			{
				Text:     "todos 3",
				Priority: 2,
				Position: 3,
				Done:     true,
			},
			{
				Text:     "todos 4",
				Priority: 3,
				Position: 4,
				Done:     true,
			},
			{
				Text:     "todos 2",
				Priority: 1,
				Position: 2,
				Done:     false,
			},
			{
				Text:     "todos 5",
				Priority: 1,
				Position: 5,
				Done:     false,
			},
		}

		sort.Sort(todo.ByPri(items))

		if !reflect.DeepEqual(items, expected) {
			t.Errorf("want to sorted correctly but got: %v", items)
		}

		fmt.Println(items)

	})
}
