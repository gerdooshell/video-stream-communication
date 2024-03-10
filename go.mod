module github.com/gerdooshell/video-stream-communication

go 1.21.6

require (
	github.com/gerdooshell/video-stream-communication/src/server-stream latest
)

replace github.com/gerdooshell/video-stream-communication/src/server-stream => ./src/server-stream
