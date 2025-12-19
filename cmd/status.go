package cmd

import (
	"fmt"

	"github.com/MrMikeDevTech/vps-tools/utils"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Muestra el estado del sistema",
	RunE: func(cmd *cobra.Command, args []string) error {
		sys := utils.GetInfo()

		fmt.Println("🖥️  Estado del sistema")

		fmt.Printf("CPU: %.2f%%\n", sys.CPUPercent)

		fmt.Printf(
			"RAM: %.2f%% (%d / %d MB)\n",
			sys.MemoryUsagePercent,
			sys.UsedMemory,
			sys.TotalMemory,
		)

		fmt.Printf(
			"Disco: %.2f%% (%d / %d GB)\n",
			sys.DiskUsagePercent,
			sys.UsedDisk,
			sys.TotalDisk,
		)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
