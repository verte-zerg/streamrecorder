# StreamRecorder

StreamRecorder is a simple Go command-line tool that allows for the recording of streaming media. With configurable parameters for timeouts, retries, and file naming. The implementation does not rely on any external dependencies, and is built using the standard library.

## Getting Started

### Installing

Clone the repository to your local machine:

```bash
go install github.com/verte-zerg/streamrecorder/cmd/streamrecorder
```

### Usage

To start recording a stream, run:
```bash
streamrecorder [options] <url>
```

Options include:
- `url`: The URL of the stream to record.
- `-prefix`: Prefix for the filename of the recorded stream. Default is `stream`.
- `-suffix`: Suffix (extension) for the filename of the recorded stream. Default is `.mp3`.
- `-timeout`: Timeout for network issues in seconds. Default is 5 seconds.
- `-attempts`: Number of attempts to recover the stream. For disabling, set to -1. Default is 60.

Examples:
```bash
# default options
streamrecorder https://example.com/stream

# custom options
streamrecorder -prefix my_stream -suffix .mp4 -timeout 10 -attempts 3 https://example.com/stream
```

## Contributing

Please feel free to contribute to the development of StreamRecorder. Whether it's bug reports, feature requests, or contributions to the code, all are welcome.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
