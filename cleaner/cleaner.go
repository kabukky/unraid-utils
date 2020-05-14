package cleaner

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Clean will search the given dirs recusively and delete all empty folders that are older than stalePeriod
func Clean(dirs []string, stalePeriod time.Duration) {
	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println("Could not walk path", path, ":", err)
				return nil
			}
			if info.IsDir() && path != dir {
				// Check if not newly created
				if info.ModTime().Before(time.Now().Add(-1 * stalePeriod)) {
					// Check if empty
					files, err := ioutil.ReadDir(path)
					if err != nil {
						log.Println("Could not read files for", path, ":", err)
						return nil
					}
					if len(files) == 0 {
						// Delete dir
						err := os.Remove(path)
						if err != nil {
							log.Println("Could not remove empty dir", path, ":", err)
							return nil
						}
						log.Println("Successfully removed empty dir", path)
					}
				}
			}
			return nil
		})
		if err != nil {
			log.Println("Could not walk ", dir, ":", err)
		}
	}
	log.Println("Done cleaning dirs.")
}
