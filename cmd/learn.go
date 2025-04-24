/*
Copyright © 2025 Peter Mazzocco <petermazzocco@gmail.com>
*/
package cmd

import (
	bufferHandler "coffee/handlers/buffer"
	"coffee/handlers/context"
	"coffee/handlers/generics"
	"coffee/handlers/goroutines"
	"coffee/handlers/interfaces"
	ioPkg "coffee/handlers/io"
	jsonPkg "coffee/handlers/json"
	"coffee/handlers/make"
	"coffee/handlers/maps"
	mutexHandler "coffee/handlers/mutex"
	"coffee/handlers/pointers"
	selectPkg "coffee/handlers/select"
	"coffee/handlers/structs"
	switchpkg "coffee/handlers/switch"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// learnCmd represents the learn command
var learnCmd = &cobra.Command{
	Use:   "learn [concept]",
	Short: "Learn Go concepts with examples",
	Long: `Explore Go concepts like goroutines, interfaces, etc. with practical coffee shop examples.

Available concepts:
  buffer     - Learn about buffers through secret menu items
  context    - Learn about the context package with delivery and pickup examples
  generics   - Learn about Go generics with customer billing
  goroutines - Learn about goroutines with batch coffee orders
  interface  - Learn about interfaces with coffee imports
  io         - Learn about input/output operations with receipts
  json       - Learn about JSON marshal/unmarshal with distributor lists
  make       - Learn about make() with seating charts and order queues
  maps       - Learn about maps with coffee shop menus
  mutex      - Learn about mutexes with coffee machine usage
  pointers   - Learn about pointers by updating orders
  select     - Learn about select with delivery apps and payment processing
  structs    - Learn about structs with customer and order data
  switch     - Learn about switch statements with inventory checks`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		concept := args[0]

		switch concept {
		case "buffer":
			if len(args) < 2 {
				fmt.Println("Usage: coffee learn buffer [item-name]")
				os.Exit(1)
			}
			bufferHandler.SecretItem(args[1])

		case "context":
			if len(args) < 2 {
				fmt.Println("Available context examples:")
				fmt.Println("  delivery [seconds] - Learn context with delivery tracking")
				fmt.Println("  pickup [code]      - Learn context with order pickup")
				os.Exit(1)
			}

			ctxType := args[1]
			switch ctxType {
			case "delivery":
				if len(args) < 3 {
					fmt.Println("Usage: coffee learn context delivery [address]")
					os.Exit(1)
				}
				context.CoffeeShopDelivery(args[2])
			case "pickup":
				if len(args) < 3 {
					fmt.Println("Usage: coffee learn context pickup [code]")
					os.Exit(1)
				}
				orderCode, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Println("Invalid format for order code:", err)
					os.Exit(1)
				}
				context.CoffeeShopPickUpOrder(orderCode)
			case "database":
				context.OrderNewProducts()
			default:
				fmt.Printf("Unknown context example: %s\n", ctxType)
				os.Exit(1)
			}

		case "generics":
			generics.GiveCustomerBill()

		case "goroutines":
			goroutines.CoffeeShopBatchOrder()

		case "interface":
			interfaces.CoffeeShopChangeImport()

		case "io":
			if len(args) < 2 {
				fmt.Println("Available io examples:")
				fmt.Println("  create - Create a new receipt")
				fmt.Println("  read   - Read an existing receipt")
				os.Exit(1)
			}

			ioType := args[1]
			switch ioType {
			case "create":
				ioPkg.CreateReceipt()
			case "read":
				ioPkg.ReadReceipt()
			default:
				fmt.Printf("Unknown io example: %s\n", ioType)
				os.Exit(1)
			}

		case "json":
			jsonPkg.UnmarshalJson()

		case "make":
			if len(args) < 2 {
				fmt.Println("Available make examples:")
				fmt.Println("  seating     - Learn make() with seating charts")
				fmt.Println("  queue [num] - Learn make() with order queues")
				os.Exit(1)
			}

			makeType := args[1]
			switch makeType {
			case "seating":
				make.CoffeeShopSeatingChart()
			case "queue":
				if len(args) < 3 {
					fmt.Println("Usage: coffee learn make queue [number-of-orders]")
					os.Exit(1)
				}
				totalOrders, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Println("Invalid format for total orders:", err)
					os.Exit(1)
				}
				make.CoffeeShopOrderQueue(totalOrders)
			default:
				fmt.Printf("Unknown make example: %s\n", makeType)
				os.Exit(1)
			}

		case "maps":
			if len(args) < 2 {
				fmt.Println("Available maps examples:")
				fmt.Println("  view           - View the coffee shop menu")
				fmt.Println("  add [item] [price] - Add a new item to the menu")
				fmt.Println("  remove [item]  - Remove an item from the menu")
				os.Exit(1)
			}

			mapType := args[1]
			switch mapType {
			case "view":
				maps.CoffeeShopMenu()
			case "add":
				if len(args) < 4 {
					fmt.Println("Usage: coffee learn maps add [item-name] [price]")
					os.Exit(1)
				}
				price, err := strconv.ParseFloat(args[3], 64)
				if err != nil {
					fmt.Println("Invalid price format:", err)
					os.Exit(1)
				}
				maps.AddNewMenuItem(args[2], price)
			case "remove":
				if len(args) < 3 {
					fmt.Println("Usage: coffee learn maps remove [item-name]")
					os.Exit(1)
				}
				maps.RemoveMenuItem(args[2])
			default:
				fmt.Printf("Unknown maps example: %s\n", mapType)
				os.Exit(1)
			}

		case "mutex":
			mutexHandler.MachineInUse()

		case "pointers":
			pointers.UpdateOrder()

		case "select":
			if len(args) < 2 {
				fmt.Println("Available select examples:")
				fmt.Println("  delivery - Learn select with delivery apps")
				fmt.Println("  payment [amount] - Learn select with payment processing")
				os.Exit(1)
			}

			selectType := args[1]
			switch selectType {
			case "delivery":
				selectPkg.DeliveryApps()
			case "payment":
				if len(args) < 3 {
					fmt.Println("Usage: coffee learn select payment [amount]")
					os.Exit(1)
				}
				amount, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Println("Invalid format for payment amount:", err)
					os.Exit(1)
				}
				selectPkg.ReceivePayment(amount)
			default:
				fmt.Printf("Unknown select example: %s\n", selectType)
				os.Exit(1)
			}

		case "structs":
			structs.CreateNewCoffee()

		case "switch":
			if len(args) < 2 {
				fmt.Println("Usage: coffee learn switch [item-to-check]")
				os.Exit(1)
			}
			switchpkg.CoffeeShopOutOfOrder(args[1])

		default:
			fmt.Printf("Unknown concept: %s\n", concept)
			fmt.Println("Run 'coffee learn' without arguments to see available concepts")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(learnCmd)
}
