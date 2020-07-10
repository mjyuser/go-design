package proxy

import "testing"

func TestProxy(t *testing.T) {
	manager := NewVideoManager()
	manager.GetVideoList()

	manager.VideoInfo(1)
	manager.DownloadVideo(1)
}