package streamrecorder

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	// DefaultTimeout is the default timeout for network issues
	DefaultTimeout = 5 * time.Second

	// DefaultPrefix is the default prefix for the filename of the recorded stream
	DefaultPrefix = "stream"

	// DefaultSuffix is the default suffix for the filename of the recorded stream
	DefaultSuffix = ".mp3"

	// DefaultRetriesCount is the default number of attempts to recover the stream
	DefaultRetriesCount = 60
)

var logger = log.New(os.Stdout, "streamrecorder: ", log.LstdFlags)

type DownloaderOptions struct {
	// Prefix and suffix for the filename of the recorded stream
	// Default: stream
	Prefix string

	// Suffix for the filename of the recorded stream
	// Default: .mp3
	Suffix string

	// In case of network issues, how long to wait before retrying
	// Default: 5 seconds
	Timeout time.Duration

	// How many times to attempt to recover the stream
	// Default: 60 retries
	// Set to -1 to disable
	Retries int
}

// DownloadStream downloads a stream from the given URL and saves it
// to files. It will retry the download if it fails, up to the number of
// retries specified in the options. The stream will be saved to files
// with the given prefix and suffix.
func DownloadStream(url string, options *DownloaderOptions) {
	if options == nil {
		options = &DownloaderOptions{}
	}

	timeout := options.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	retries := options.Retries
	if retries == 0 {
		retries = DefaultRetriesCount
	}

	prefix := options.Prefix
	if prefix == "" {
		prefix = DefaultPrefix
	}

	suffix := options.Suffix
	if options.Suffix == "" {
		suffix = DefaultSuffix
	}

	var retry, part int
	for {
		if retry > 0 {
			logger.Printf("[INFO] Attempt %d of %d", retry, retries)
		}

		filename := fmt.Sprintf("%s_%d%s", prefix, part, suffix)
		success := downloadAndFlushStream(url, filename)

		if success {
			logger.Printf("[INFO] Stream was recorded to %s", filename)
			part++
		}

		retry++
		if retry >= retries {
			logger.Printf("[ERROR] Stream recovery failed after %d attempts. Stopping.", retries+1)
			break
		}

		logger.Printf("[WARNING] Stream recovery failed. Retrying in %s...", timeout)

		time.Sleep(timeout)
	}

	logger.Printf("[INFO] Stream was recorder to %d files", part)
}

func downloadAndFlushStream(url, filename string) bool {
	resp, err := http.Get(url)
	if err != nil {
		logger.Printf("[ERROR] Error connecting to stream: %s", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Printf("[ERROR] Received non-200 status code: %d", resp.StatusCode)
		return false
	}

	file, err := os.Create(filename)
	if err != nil {
		logger.Printf("[ERROR] Error creating file: %s", err)
		return false
	}
	defer file.Close()

	logger.Printf("[INFO] Recording stream to %s", filename)

	buffer := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Printf("[ERROR] Error reading from stream: %s", err)
			return true
		}

		_, writeErr := file.Write(buffer[:n])
		if writeErr != nil {
			logger.Printf("[ERROR] Error writing to file: %s", writeErr)
			return true
		}

		flushErr := file.Sync()
		if flushErr != nil {
			logger.Printf("[ERROR] Error flushing file: %s", flushErr)
			return true
		}
	}

	return true
}
