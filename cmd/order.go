/*
Copyright © 2025 Peter Mazzocco <petermazzocco@gmail.com>
*/
package cmd

import (
	"coffee/handlers/context"
	"coffee/handlers/goroutines"
	"coffee/handlers/make"
	"coffee/handlers/pointers"
	switchpkg "coffee/handlers/switch"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	goroutinesBool  bool
	orderQueueBool  bool
	switchBool      bool
	ctxBool         bool
	pickUpCtxBool   bool
	updateOrderBool bool
)

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Run the order function to order coffees",
	Long: `In order to execute some examples, run the order
	command with a flag for which operation you're looking to run.

	Run the coffee order -h to view a list of available operations.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case goroutinesBool:
			goroutines.CoffeeShopBatchOrder()
		case orderQueueBool:
			totalOrders, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid format for total orders:", err)
				return
			}
			make.CoffeeShopOrderQueue(totalOrders)
		case switchBool:
			switchpkg.CoffeeShopOutOfOrder(args[0])
		case ctxBool:
			context.CoffeeShopDelivery(args[0])
		case pickUpCtxBool:
			orderCode, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid format for order code:", err)
				return
			}
			context.CoffeeShopPickUpOrder(orderCode)
		case updateOrderBool:
			pointers.UpdateOrder()
		default:
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(orderCmd)
	orderCmd.Flags().BoolVarP(&goroutinesBool, "batch", "g", false, "Learning about goroutines? Place a batch order of coffees at once!")
	orderCmd.Flags().BoolVarP(&orderQueueBool, "queue", "q", false, "Learning about channels? Let's queue up orders for our coffee shop.")
	orderCmd.Flags().BoolVarP(&switchBool, "switch", "x", false, "Learning about switches? Let's see what our inventory is.")
	orderCmd.Flags().BoolVarP(&ctxBool, "delivery", "d", false, "Learning about context? A customer just placed a delivery order, can we make it in time?.")
	orderCmd.Flags().BoolVarP(&pickUpCtxBool, "pickup", "p", false, "Learning about context? A customer just placed a pickup order, do they have the right pick up code?.")
	orderCmd.Flags().BoolVarP(&updateOrderBool, "update", "u", false, "Learning about pointers? Let's update our preivous order.")

}
