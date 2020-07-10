package proxy

import (
	"fmt"
	"time"
)

type ThirdPartyVideo interface {
	GetVideoList() []*Video
	VideoInfo(id uint64) *Video
	DownloadVideo(id uint64) bool
}

type Video struct {
	Id uint64 `json:"id"`
	DownloadTime time.Time
}

type VideoClient struct {}

func (v *VideoClient) GetVideoList() []*Video {
	fmt.Println("get video list")
	videos := make([]*Video, 5)
	return videos
}
func (v *VideoClient) VideoInfo(id uint64) *Video {
	fmt.Printf("get video info: id: %d\n", id)
	return &Video{Id: id}
}
func (v *VideoClient) DownloadVideo(id uint64) bool {
	fmt.Printf("start download video. id: %d\n", id)
	time.Sleep(time.Microsecond * 100)
	fmt.Printf("download video success. id: %d\n", id)
	return true
}


type VideoProxy interface {
	ThirdPartyVideo
	VideoCounter()
}


type VideoManager struct {
	VideoMap map[uint64]int
	VideoList []*Video
	VideoClient VideoClient
}

func NewVideoManager() *VideoManager {
	return &VideoManager{
		VideoMap: make(map[uint64]int),
		VideoList: make([]*Video, 0),
		VideoClient: VideoClient{},
	}
}

func (v *VideoManager) VideoCounter() {}

func (v *VideoManager) GetVideoList() []*Video {
	fmt.Printf("get video list from manager\n")
	if len(v.VideoList) > 0 {
		return v.VideoList
	}

	v.VideoList = v.VideoClient.GetVideoList()
	return v.VideoList
}

func (v *VideoManager) VideoInfo(id uint64) *Video {
	if index, ok := v.VideoMap[id]; ok {
		return v.VideoList[index]
	}
	video := v.VideoClient.VideoInfo(id)
	v.VideoList = append(v.VideoList, video)
	v.VideoMap[id] = len(v.VideoList)

	return video
}

func (v *VideoManager) DownloadVideo(id uint64) bool {
	fmt.Printf("download video from video manager")
	if _, ok := v.VideoMap[id]; ok {
		return true
	}
	return v.VideoClient.DownloadVideo(id)
}

