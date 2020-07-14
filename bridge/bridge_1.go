package bridge

type Device interface {
	IsEnabled() bool
	EnAble()
	DisAble()
	GetVolume() uint64
	SetVolume(volume uint64)
	GetChannel() uint64
	SetChannel(channel uint64)
}

type RemoteControl struct {
	 Device
}

func NewRemoteControl(device Device) *RemoteControl {
	return &RemoteControl{
		Device: device,
	}
}

func (r *RemoteControl) TogglePower() {
	if r.IsEnabled() {
		r.DisAble()
		return
	}

	r.EnAble()
}

func (r *RemoteControl) VolumeDown() {
	r.SetVolume(r.GetVolume() - 10)
}

func (r *RemoteControl) VolumeUp() {
	r.SetVolume(r.GetVolume() + 10)
}

func (r *RemoteControl) ChannelDown() {
	r.SetChannel(r.GetChannel() - 1)
}

func (r *RemoteControl) ChannelUp() {
	r.SetChannel(r.GetChannel() + 1)
}

func (r *RemoteControl) Mute() {
	r.SetVolume(0)
}

type AdvanceRemoteControl interface {
	Mute()
}

