package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"faizisyellow.com/tri/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  `Listing the todos`,
	Run:   ListRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ListRun(cmd *cobra.Command, args []string) {

	items, err := todo.ReadItems(dataFile)
	if err != nil {
		fmt.Println("error reading items: ", err)
		return
	}

	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	for _, i := range items {
		fmt.Fprintln(w, i.Label(), i.PrettyP()+"\t"+i.Text+"\t")
	}

	w.Flush()
}
