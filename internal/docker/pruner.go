/*
Copyright © 2025 Elwan Mayencourt <mayencourt@elwan.ch>
*/
package docker

import (
	"log"
	"time"
)

func StartPruner(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		log.Printf("Pruner started, will run every %v", interval)

		for range ticker.C {
			pruneContainers()
		}
	}()
}

func pruneContainers() {
	containers, err := ListPruneContainers()
	if err != nil {
		log.Printf("Failed to list prune containers: %v", err)
		return
	}

	if len(containers) == 0 {
		return
	}

	log.Printf("Pruning %d containers...", len(containers))

	deleted := 0
	for _, containerName := range containers {
		if err := ForceDeleteContainer(containerName); err != nil {
			log.Printf("Failed to delete prune container %s: %v", containerName, err)
		} else {
			deleted++
		}
	}

	if deleted > 0 {
		log.Printf("Successfully pruned %d/%d containers", deleted, len(containers))
	}
}
