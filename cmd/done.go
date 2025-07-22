package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"faizisyellow.com/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Marks Item as Done",
	Run:     DoneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func DoneRun(cmd *cobra.Command, args []string) {

	data := dataFile
	if viper.GetString("datafile") != "" {
		data = viper.GetString("datafile")
	}

	items, err := todo.ReadItems(data)
	if err != nil {
		fmt.Println("error while reading items: ", err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")

		sort.Sort(todo.ByPri(items))
		todo.SaveItems(data, items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}
