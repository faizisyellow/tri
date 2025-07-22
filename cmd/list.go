package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"faizisyellow.com/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  `Listing the todos`,
	Run:   ListRun,
}

var (
	doneOpt   bool
	allOpt    bool
	searchOpt string
)

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos ")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos ")
	listCmd.Flags().StringVarP(&searchOpt, "search", "s", "", "Search Text items")
}

func ListRun(cmd *cobra.Command, args []string) {

	data := dataFile
	if viper.GetString("datafile") != "" {
		data = viper.GetString("datafile")
	}

	items, err := todo.ReadItems(data)
	if err != nil {
		fmt.Println("error reading items: ", err)
		return
	}

	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	var bufItems []*todo.Item

	for _, i := range items {

		if allOpt || i.Done == doneOpt {
			bufItems = append(bufItems, &i)
		}
	}

	var found int

	for _, v := range bufItems {

		if searchOpt != "" && strings.Contains(v.Text, searchOpt) {
			fmt.Fprintln(w, v.Label(), "\t"+v.PrettyDone(), "\t", v.PrettyP()+"\t"+v.Text+"\t")
			found = +1
		} else {
			fmt.Fprintln(w, v.Label(), "\t"+v.PrettyDone(), "\t", v.PrettyP()+"\t"+v.Text+"\t")
		}
	}

	if found == 0 && searchOpt != "" {
		fmt.Fprint(w, "no matching any items\n")
	}

	w.Flush()

}
