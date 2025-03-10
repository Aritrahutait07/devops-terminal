package monitor

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	/*"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"*/)

// GetSystemStats fetches and returns CPU, Memory, and Disk usage
func GetSystemStats() {
	fmt.Println("🔍 Fetching system stats...\n")

	// CPU Usage
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatalf("Error fetching CPU usage: %v", err)
	}
	fmt.Printf("🖥️  CPU Usage: %.2f%%\n", cpuPercent[0])

	// Memory Usage
	vmStats, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("Error fetching memory stats: %v", err)
	}
	fmt.Printf("🟢 RAM Usage: %.2f%% (%v used out of %v)\n", vmStats.UsedPercent, formatBytes(vmStats.Used), formatBytes(vmStats.Total))

	// Disk Usage
	diskStats, err := disk.Usage("/")
	if err != nil {
		log.Fatalf("Error fetching disk stats: %v", err)
	}
	fmt.Printf("💾 Disk Usage: %.2f%% (%v used out of %v)\n", diskStats.UsedPercent, formatBytes(diskStats.Used), formatBytes(diskStats.Total))

	fmt.Println("\n✅ Monitoring complete!")
}

// formatBytes converts bytes to human-readable format
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
