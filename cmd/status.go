package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Muestra el estado del sistema",
	RunE: func(cmd *cobra.Command, args []string) error {
		cpuPercent, _ := cpu.Percent(0, false)
		vm, _ := mem.VirtualMemory()
		d, _ := disk.Usage("/")

		fmt.Println("🖥️  Estado del sistema")
		fmt.Printf("CPU: %.2f%%\n", cpuPercent[0])
		fmt.Printf("RAM: %.2f%% (%v / %v MB)\n",
			vm.UsedPercent,
			vm.Used/1024/1024,
			vm.Total/1024/1024,
		)
		fmt.Printf("Disco: %.2f%% (%v / %v GB)\n",
			d.UsedPercent,
			d.Used/1024/1024/1024,
			d.Total/1024/1024/1024,
		)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
