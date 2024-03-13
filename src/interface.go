package live_streaming

type StreamService interface {
	CreateStream()
	StartStream()
	StopStream()
	ListStream() error
	GetStreamDetails() error
}
