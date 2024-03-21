package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/verte-zerg/streamrecorder"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: streamrecorder [options] <url>")
		flag.PrintDefaults()
	}

	prefix := flag.String(
		"prefix",
		streamrecorder.DefaultPrefix,
		"Prefix for the filename of the recorded stream",
	)

	suffix := flag.String(
		"suffix",
		streamrecorder.DefaultSuffix,
		"Suffix (extension) for the filename of the recorded stream",
	)

	defaultTimeoutSeconds := int(streamrecorder.DefaultTimeout.Seconds())
	timeoutSeconds := flag.Int(
		"timeout",
		defaultTimeoutSeconds,
		"Timeout for network issues in seconds",
	)

	retries := flag.Int(
		"attempts",
		streamrecorder.DefaultRetriesCount,
		"Number of attempts to recover the stream. For disabling, set to -1",
	)

	flag.Parse()

	url := flag.Arg(0)

	if url == "" {
		fmt.Println("Error: URL is required")
		fmt.Println("Usage: stream-recorder [options] <url>")
		flag.PrintDefaults()
		return
	}

	timeout := time.Duration(*timeoutSeconds) * time.Second

	options := &streamrecorder.DownloaderOptions{
		Timeout: timeout,
		Retries: *retries,
		Prefix:  *prefix,
		Suffix:  *suffix,
	}

	streamrecorder.DownloadStream(url, options)
}
