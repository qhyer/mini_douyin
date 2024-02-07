package ffmpeg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFrameAsJpeg(t *testing.T) {
	res, err := ReadFrameAsJpeg("https://lf3-cdn-tos.bytescm.com/obj/ttfe/ATSX/mainland/gongquhunjian_1080.min.mp4")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
