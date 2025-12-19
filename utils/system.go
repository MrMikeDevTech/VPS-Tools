package utils

import (
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/term"
)

type SystemInfo struct {
	Username           string
	Host               string
	OperativeSystem    string
	Platform           string
	PlatformVersion    string
	Uptime             time.Duration
	KernelVersion      string
	Architecture       string
	CPUModel           string
	CPUPercent         float64
	TotalMemory        uint64
	MemoryUsagePercent float64
	UsedMemory         uint64
	FreeMemory         uint64
	DiskUsagePercent   float64
	UsedDisk           uint64
	TotalDisk          uint64
	Date               string
	Time               string
}

func GetCurrentUsername() string {
	u, err := user.Current()

	if err != nil {
		return "unknown"
	}

	username := u.Username

	if strings.Contains(username, "\\") {
		parts := strings.Split(username, "\\")
		return parts[len(parts)-1]
	}

	return username
}

func GetInfo() SystemInfo {
	// ======================
	// CPU
	// ======================
	cpuPercent, _ := cpu.Percent(0, false)
	cpuInfo, _ := cpu.Info()

	cpuModel := "Unknown"
	if len(cpuInfo) > 0 {
		cpuModel = cpuInfo[0].ModelName
	}

	// ======================
	// Memory
	// ======================
	vm, _ := mem.VirtualMemory()

	totalMemory := vm.Total / 1024 / 1024
	usedMemory := vm.Used / 1024 / 1024

	// ======================
	// Disk
	// ======================
	diskUsage, _ := disk.Usage("/")

	totalDiskGB := diskUsage.Total / 1024 / 1024 / 1024
	usedDiskGB := diskUsage.Used / 1024 / 1024 / 1024

	// ======================
	// Host / System
	// ======================
	hostInfo, _ := host.Info()

	username := GetCurrentUsername()
	platform := strings.ToLower(hostInfo.Platform)

	operativeSystem := "Unknown"
	switch {
	case strings.Contains(platform, "windows"):
		operativeSystem = "Windows"
	case strings.Contains(platform, "debian"):
		operativeSystem = "Debian"
	case strings.Contains(platform, "ubuntu"):
		operativeSystem = "Ubuntu"
	case strings.Contains(platform, "centos"):
		operativeSystem = "CentOS"
	case strings.Contains(platform, "fedora"):
		operativeSystem = "Fedora"
	case strings.Contains(platform, "darwin"), strings.Contains(platform, "mac"):
		operativeSystem = "MacOS"
	}

	// ======================
	// Date and Time
	// ======================
	currentTime := time.Now()
	date := currentTime.Format("02/01/2006")
	timeStr := currentTime.Format("15:04:05")

	// ======================
	// Build response
	// ======================
	return SystemInfo{
		Username:           username,
		Host:               hostInfo.Hostname,
		OperativeSystem:    operativeSystem,
		Platform:           hostInfo.Platform,
		PlatformVersion:    hostInfo.PlatformVersion,
		Uptime:             time.Duration(hostInfo.Uptime) * time.Second,
		KernelVersion:      hostInfo.KernelVersion,
		Architecture:       hostInfo.KernelArch,
		CPUModel:           cpuModel,
		CPUPercent:         cpuPercent[0],
		TotalMemory:        totalMemory,
		UsedMemory:         usedMemory,
		FreeMemory:         vm.Free,
		MemoryUsagePercent: vm.UsedPercent,
		TotalDisk:          totalDiskGB,
		UsedDisk:           usedDiskGB,
		DiskUsagePercent:   diskUsage.UsedPercent,
		Date:               date,
		Time:               timeStr,
	}
}

func GetTerminalWidth() int {
	fd := int(os.Stdout.Fd())
	if term.IsTerminal(fd) {
		if w, _, err := term.GetSize(fd); err == nil {
			return w
		}
	}
	return 50
}
