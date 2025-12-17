package cmd

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Muestra los componentes del sistema",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("🖥️  Detalles del sistema")
		fmt.Println("────────────────────────")

		// ─── HOST / OS ───────────────────────────────
		hostInfo, _ := host.Info()

		fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
		fmt.Printf("Sistema: %s %s\n", hostInfo.Platform, hostInfo.PlatformVersion)
		fmt.Printf("Kernel: %s\n", hostInfo.KernelVersion)
		fmt.Printf("Arquitectura: %s\n", runtime.GOARCH)
		fmt.Println()

		// ─── CPU ─────────────────────────────────────
		cpuInfo, _ := cpu.Info()
		cpuCores, _ := cpu.Counts(false)
		cpuThreads, _ := cpu.Counts(true)

		if len(cpuInfo) > 0 {
			fmt.Println("🧠 CPU")
			fmt.Printf("Modelo: %s\n", cpuInfo[0].ModelName)
			fmt.Printf("Cores físicos: %d\n", cpuCores)
			fmt.Printf("Threads: %d\n", cpuThreads)
			fmt.Printf("Frecuencia: %.0f MHz\n", cpuInfo[0].Mhz)
			fmt.Println()
		}

		// ─── MEMORIA ─────────────────────────────────
		vm, _ := mem.VirtualMemory()

		fmt.Println("🧬 Memoria RAM")
		fmt.Printf("Total: %d MB\n", vm.Total/1024/1024)
		fmt.Println()

		// ─── DISCO ───────────────────────────────────
		partitions, _ := disk.Partitions(false)

		fmt.Println("💾 Discos")
		for _, p := range partitions {
			usage, err := disk.Usage(p.Mountpoint)
			if err != nil {
				continue
			}

			fmt.Printf(
				"- %s (%s): %d GB total\n",
				p.Mountpoint,
				p.Fstype,
				usage.Total/1024/1024/1024,
			)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(detailsCmd)
}
