/*
Copyright © 2025 Peter Mazzocco <petermazzocco@gmail.com>
*/
package cmd

import (
	"coffee/handlers/interfaces"
	ioPkg "coffee/handlers/io"
	jsonPkg "coffee/handlers/json"
	"coffee/handlers/make"
	"coffee/handlers/maps"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	slicesBool         bool
	mapsBool           bool
	newMenuItemBool    bool
	deleteMenuItemBool bool
	interfacesBool     bool
	jsonBool           bool
	receiptBool        bool
	readRecptBool      bool
)

// managerCmd represents the manager command
var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "A list of functions that mirror managing a Coffee Shop",
	Long: `Want to learn about Intefaces, make(), maps and more?

The Manager cmd will give you the tools to manage your coffee shop
and learn Golang.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case slicesBool:
			make.CoffeeShopSeatingChart()
		case mapsBool:
			maps.CoffeeShopMenu()
		case newMenuItemBool:
			price, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				fmt.Println("Invalid price format:", err)
				return
			}
			maps.AddNewMenuItem(args[0], price)
		case deleteMenuItemBool:
			maps.RemoveMenuItem(args[0])
		case jsonBool:
			jsonPkg.UnmarshalJson()
		case interfacesBool:
			interfaces.CoffeeShopChangeImport()
		case receiptBool:
			ioPkg.CreateReceipt()
		case readRecptBool:
			ioPkg.ReadReceipt()
		}

	},
}

func init() {
	rootCmd.AddCommand(managerCmd)
	managerCmd.Flags().BoolVarP(&slicesBool, "seating", "s", false, "Learning about slices? Let's check our seating capacity.")
	managerCmd.Flags().BoolVarP(&newMenuItemBool, "new-item", "n", false, "Learning about maps? Let's add a new item to the menu.")
	managerCmd.Flags().BoolVarP(&deleteMenuItemBool, "delete-item", "r", false, "Learning about maps? Let's remove an item from the menu.")
	managerCmd.Flags().BoolVarP(&interfacesBool, "import", "i", false, "Learning about interfaces? Let's change our coffee shop's import.")
	managerCmd.Flags().BoolVarP(&mapsBool, "menu", "m", false, "Learning about maps? Let's check our available items on the menu.")
	managerCmd.Flags().BoolVarP(&jsonBool, "json", "j", false, "Learning about JSON? Let's see if we can get a list from our new distributor for their coffee selection.")
	managerCmd.Flags().BoolVarP(&receiptBool, "new-receipt", "c", false, "Learning about receipts? Let's create a receipt.")
	managerCmd.Flags().BoolVarP(&readRecptBool, "read-receipt", "o", false, "Learning about receipts? Let's read a receipt.")
}
