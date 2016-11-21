package audio

import (
	"errors"
	"strings"
	"time"

	"github.com/dh1tw/samplerate"
	"github.com/gordonklaus/portaudio"
)

const (
	INPUT  = 1
	OUTPUT = 2
)

const (
	MONO   = 1
	STEREO = 2
)

var bitMapToInt32 = map[int32]float32{
	8:  255,
	16: 32767,
	32: 2147483647,
}

var bitMapToFloat32 = map[int]float32{
	8:  256,
	16: 32768,
	32: 2147483648,
}

type AudioStream struct {
	DeviceName      string
	Direction       int
	Channels        int
	Samplingrate    float64
	Latency         time.Duration
	FramesPerBuffer int
	Device          *portaudio.DeviceInfo
	Stream          *portaudio.Stream
	Converter       samplerate.Src
	out             []float32
	in              []float32
}

type AudioMsg struct {
	Data  []byte
	Raw   []int16
	Topic string
}

type AudioDevice struct {
	AudioStream
	AudioInCh       chan AudioMsg
	AudioOutCh      chan AudioMsg
	AudioLoopbackCh chan AudioMsg
	EventCh         chan interface{}
}

// IdentifyDevice checks if the Audio Devices actually exist
func (as *AudioDevice) IdentifyDevice() error {
	devices, _ := portaudio.Devices()
	for _, device := range devices {
		if device.Name == as.DeviceName {
			as.Device = device
			return nil
		}
	}
	return errors.New("unknown audio device")
}

// GetChannel returns the integer representation of channels
func GetChannel(ch string) int {
	if strings.ToUpper(ch) == "MONO" {
		return MONO
	} else if strings.ToUpper(ch) == "STEREO" {
		return STEREO
	}
	return 0
}
