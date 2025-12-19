package cmd

import (
	"fmt"
	"strings"

	"github.com/MrMikeDevTech/vps-tools/utils"
	"github.com/spf13/cobra"
)

func formatLogoLine(line string, width int) string {
	lineLen := len(line)
	if lineLen >= width {
		return line
	}

	paddingLeft := (width - lineLen) / 2
	paddingRight := width - lineLen - paddingLeft

	return strings.Repeat(" ", paddingLeft) +
		line +
		strings.Repeat(" ", paddingRight)
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Muestra información del sistema",
	RunE: func(cmd *cobra.Command, args []string) error {
		sys := utils.GetInfo()

		underline := strings.Repeat("─", len(sys.Username)+len(sys.Host)+1)

		title_color := utils.Colors.Green
		subtle_color := utils.Colors.Blue

		info := []string{
			fmt.Sprintf("%s%s@%s%s", title_color, sys.Username, sys.Host, utils.Colors.Reset),
			fmt.Sprintf("%s%s%s", title_color, underline, utils.Colors.Reset),
			fmt.Sprintf("%sSistema:%s %s", subtle_color, utils.Colors.Reset, sys.Platform),
			fmt.Sprintf("%sKernel:%s %s", subtle_color, utils.Colors.Reset, sys.KernelVersion),
			fmt.Sprintf("%sArquitectura:%s %s", subtle_color, utils.Colors.Reset, sys.Architecture),
			fmt.Sprintf("%sTiempo activo:%s %s", subtle_color, utils.Colors.Reset, utils.UptimeFormat(sys.Uptime)),
			fmt.Sprintf("%sCPU:%s %s", subtle_color, utils.Colors.Reset, sys.CPUModel),
			fmt.Sprintf("%sMemoria:%s %dMB / %dMB",
				subtle_color,
				utils.Colors.Reset,
				sys.UsedMemory,
				sys.TotalMemory,
			),
		}

		currentLogo := utils.GetLogoByOS(sys.OperativeSystem)
		maxLines := max(len(currentLogo), len(info))
		const logoWidth = 30

		fmt.Println()
		for i := range maxLines {
			left := strings.Repeat(" ", logoWidth)
			right := ""

			if i < len(currentLogo) {
				left = utils.Colors.Cyan +
					formatLogoLine(currentLogo[i], logoWidth) +
					utils.Colors.Reset
			}

			if i < len(info) {
				right = info[i]
			}

			fmt.Printf("%s %s\n", left, right)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
