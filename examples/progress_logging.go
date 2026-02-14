package main

import (
	"fmt"
	"math/rand"
	"time"

	versalog "github.com/VersaLog/VersaLog.go/VersaLog"
)

func processLine(log *versalog.VersaLog, line int) {
	time.Sleep(100 * time.Millisecond)

	if rand.Float64() < 0.05 {
		log.Warning(
			fmt.Sprintf("Line %d took longer than expected", line),
		)
	}
}

func processFile(log *versalog.VersaLog, fileIndex int, totalFiles int) {
	log.Step(
		fmt.Sprintf("Processing file_%d.txt", fileIndex),
		fileIndex,
		totalFiles,
	)

	defer log.Timer(
		fmt.Sprintf("file_%d.txt", fileIndex),
	)()

	totalLines := 20

	for i := 1; i <= totalLines; i++ {
		processLine(log, i)

		log.Progress(
			fmt.Sprintf("file_%d.txt", fileIndex),
			i,
			totalLines,
		)
	}
}

func main() {
	log := versalog.NewVersaLog(
		"detailed",
		false,
		true,
		"BATCH",
		false,
		false,
		false,
		nil,
		false,
		false,
	)

	totalFiles := 5

	log.Info("Batch Start")

	defer log.Timer("Total Batch")()

	for i := 1; i <= totalFiles; i++ {
		processFile(log, i, totalFiles)

		log.Progress(
			"Overall Progress",
			i,
			totalFiles,
		)
	}

	log.Info("Batch Finished")
}
