package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func getTerminalWidth() int {
	fd := int(os.Stdout.Fd())
	if term.IsTerminal(fd) {
		if w, _, err := term.GetSize(fd); err == nil {
			return w
		}
	}
	return 50
}

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

type Title struct {
	Text  string
	ASCII string
}

func printTitle(width int) {
	titles := []Title{
		{
			Text: "Welcome to MrMikeDev VPS",
			ASCII: `
 ____  _                           _     _                 __  __      __  __ _ _        _____              __      _______   _____ 
|  _ \(_)                         (_)   | |               |  \/  |    |  \/  (_) |      |  __ \             \ \    / /  __ \ / ____|
| |_) |_  ___ _ ____   _____ _ __  _  __| | ___     __ _  | \  / |_ __| \  / |_| | _____| |  | | _____   __  \ \  / /| |__) | (___  
|  _ <| |/ _ \ '_ \ \ / / _ \ '_ \| |/ _' |/ _ \   / _' | | |\/| | '__| |\/| | | |/ / _ \ |  | |/ _ \ \ / /   \ \/ / |  ___/ \___ \ 
| |_) | |  __/ | | \ V /  __/ | | | | (_| | (_) | | (_| | | |  | | |  | |  | | |   <  __/ |__| |  __/\ V /     \  /  | |     ____) |
|____/|_|\___|_| |_|\_/ \___|_| |_|_|\__,_|\___/   \__,_| |_|  |_|_|  |_|  |_|_|_|\_\___|_____/ \___| \_/       \/   |_|    |_____/ 
`,
		},
		{
			Text: "MrMikeDev VPS",
			ASCII: `
 __  __      __  __ _ _        _____              __      _______   _____
|  \/  |    |  \/  (_) |      |  __ \             \ \    / /  __ \ / ____|
| \  / |_ __| \  / |_| | _____| |  | | _____   __  \ \  / /| |__) | (___
| |\/| | '__| |\/| | | |/ / _ \ |  | |/ _ \ \ / /   \ \/ / |  ___/ \___ \
| |  | | |  | |  | | |   <  __/ |__| |  __/\ V /     \  /  | |     ____) |
|_|  |_|_|  |_|  |_|_|_|\_\___|_____/ \___| \_/       \/   |_|    |_____/
`,
		},
		{
			Text: "MrMikeDev",
			ASCII: `
 __  __      __  __ _ _        _____
|  \/  |    |  \/  (_) |      |  __ \
| \  / |_ __| \  / |_| | _____| |  | | _____   __
| |\/| | '__| |\/| | | |/ / _ \ |  | |/ _ \ \ / /
| |  | | |  | |  | | |   <  __/ |__| |  __/\ V /
|_|  |_|_|  |_|  |_|_|_|\_\___|_____/ \___| \_/
`,
		},
		{
			Text: "MMD VPS",
			ASCII: `
 __  __ __  __ ____    __      _______   _____
|  \/  |  \/  |  _ \   \ \    / /  __ \ / ____|
| \  / | \  / | | | |   \ \  / /| |__) | (___
| |\/| | |\/| | | | |    \ \/ / |  ___/ \___ \
| |  | | |  | | |_| |     \  /  | |     ____) |
|_|  |_|_|  |_|____/       \/   |_|    |_____/
`,
		},
		{
			Text: "VPS",
			ASCII: `
 __      _______   _____
 \ \    / /  __ \ / ____|
  \ \  / /| |__) | (___
   \ \/ / |  ___/ \___ \
    \  /  | |     ____) |
     \/   |_|    |_____/
`,
		},
	}

	for _, t := range titles {
		lines := strings.Split(strings.Trim(t.ASCII, "\n"), "\n")
		artWidth := maxLineLength(lines)

		if artWidth <= width {
			padding := (width - artWidth) / 2
			for _, line := range lines {
				fmt.Println(strings.Repeat(" ", padding) + line)
			}
			return
		}
	}

	text := "Bienvenido a MrMikeDev VPS"
	padding := max(0, (width-len(text))/2)
	fmt.Println(strings.Repeat(" ", padding) + text)
}

func maxLineLength(lines []string) int {
	max := 0
	for _, l := range lines {
		if len(l) > max {
			max = len(l)
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var welcomeCmd = &cobra.Command{
	Use:   "welcome",
	Short: "Mensaje de bienvenida del VPS",
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := user.Current()

		username := u.Username
		if strings.Contains(username, "\\") {
			parts := strings.Split(username, "\\")
			username = parts[len(parts)-1]
		}

		now := time.Now().Format("2006-01-02 15:04:05")
		width := getTerminalWidth()

		printFrame(width)
		printTitle(width)
		printFrame(width)

		fmt.Printf("\n👋 Bienvenido, %s\n", username)
		fmt.Printf("📅 Fecha y hora: %s\n", now)
	},
}

func init() {
	rootCmd.AddCommand(welcomeCmd)
}
