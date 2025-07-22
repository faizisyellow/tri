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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Item text args: pos, new text",
	Run:   UpdateRun,
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func UpdateRun(cmd *cobra.Command, args []string) {

	data := dataFile
	if viper.GetString("datafile") != "" {
		data = viper.GetString("datafile")
	}

	items, err := todo.ReadItems(data)
	if err != nil {
		log.Fatalln("error when reading items: ", err)
	}

	p, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if p > 0 && p < len(items) {
		items[p-1].Text = args[1]

		sort.Sort(todo.ByPri(items))
		todo.SaveItems(data, items)

		fmt.Printf("update %v to %v sucessfully\n", p, args[1])
	} else {
		log.Println(p, "doesn't match any items")
	}
}
