package cmd

import (
	"fmt"
	"strings"

	"github.com/MrMikeDevTech/vps-tools/utils"
	"github.com/spf13/cobra"
)

func printFrame(width int) {
	block := []string{
		"  ______ ",
		" |______|",
	}

	blockWidth := len(block[0])
	count := width / blockWidth
	if count < 1 {
		count = 1
	}

	for _, line := range block {
		full := strings.Repeat(line, count)
		fmt.Println(full[:min(len(full), width)])
	}
}

func printTitle(width int) {
	title, artWidth := utils.GetTitle(width)
	lines := strings.Split(strings.Trim(title.ASCII, "\n"), "\n")
	padding := (width - artWidth) / 2

	for _, line := range lines {
		fmt.Println(strings.Repeat(" ", padding) + line)
	}
}

var welcomeCmd = &cobra.Command{
	Use:   "welcome",
	Short: "Mensaje de bienvenida del VPS",
	Run: func(cmd *cobra.Command, args []string) {
		sys := utils.GetInfo()
		username := sys.Username

		width := utils.GetTerminalWidth()

		printFrame(width)
		printTitle(width)
		printFrame(width)

		fmt.Printf("\n👋 Bienvenido, %s\n", username)
		fmt.Printf("📅 Fecha: %s\n", sys.Date)
		fmt.Printf("⏰ Hora: %s\n", sys.Time)
	},
}

func init() {
	rootCmd.AddCommand(welcomeCmd)
}
