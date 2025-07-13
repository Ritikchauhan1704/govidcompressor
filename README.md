# GoVidCompressor

🚧 **Under Construction** 🚧

A high-performance Go application for batch video compression that reduces file sizes while maintaining original quality.

## Features

- **Batch Processing**: Compress all videos in a directory automatically
- **Quality Preservation**: Lossless compression maintains original video quality
- **Multi-format Support**: Handles various video formats (MP4, AVI, MOV, MKV, etc.)
- **Concurrent Processing**: Leverages Go's goroutines for faster compression
- **Progress Tracking**: Real-time compression progress and statistics
- **Cross-platform**: Works on Windows, macOS, and Linux

## Quick Start

```bash
# Install dependencies
go mod tidy

# Run the compressor
go run main.go -input ./videos -output ./compressed

# Or build and run
go build -o govidcompressor
./govidcompressor -input ./videos -output ./compressed
```

## Usage

```bash
govidcompressor [options]
  -input    string    Input directory containing videos (default: "./videos")
  -output   string    Output directory for compressed videos (default: "./compressed")
  -format   string    Output format (mp4, avi, mov) (default: "mp4")
  -threads  int       Number of concurrent workers (default: CPU cores)
```

## Requirements

- Go 1.19 or higher
- FFmpeg installed on system

## Installation

1. Clone the repository
2. Install FFmpeg on your system
3. Run `go mod tidy`
4. Build with `go build`
