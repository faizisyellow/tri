package cmd

import (
	"fmt"

	"faizisyellow.com/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

// AddCmd represents the Add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   AddRun,
}

func init() {
	rootCmd.AddCommand(AddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	AddCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}

func AddRun(cmd *cobra.Command, args []string) {

	data := dataFile
	if viper.GetString("datafile") != "" {
		data = viper.GetString("datafile")
	}

	items, _ := todo.ReadItems(data)

	for _, v := range args {
		item := todo.Item{Text: v}
		item.SetPriority(priority)

		items = append(items, item)
	}

	err := todo.SaveItems(data, items)
	if err != nil {
		fmt.Println("error save todo: ", err)
	}
}
