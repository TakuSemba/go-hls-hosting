package handle

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type FakeLoader struct{}

func (f *FakeLoader) LoadMasterPlaylist() ([]byte, error) {
	return []byte("test-master-playlist"), nil
}

func (f *FakeLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	return []byte("test-media-playlist: index=" + strconv.Itoa(index)), nil
}

func (f *FakeLoader) LoadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error) {
	return []byte("test-segment: mediaPlaylistIndex=" + strconv.Itoa(mediaPlaylistIndex) + ", segmentIndex=" + strconv.Itoa(segmentIndex)), nil
}

var FakeHandler = Handler{
	VodLoader:   &FakeLoader{},
	LiveLoader:  &FakeLoader{},
	ChaseLoader: &FakeLoader{},
}

func TestVodMasterPlaylist(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/playlist.m3u8")
	handler := FakeHandler

	err := handler.VodMasterPlaylist(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMpegUrl, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-master-playlist", recorder.Body.String())
	}
}

func TestVodMediaPlaylist(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/:index/playlist.m3u8")
	context.SetParamNames("index")
	context.SetParamValues("0")
	handler := FakeHandler

	err := handler.VodMediaPlaylist(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMpegUrl, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-media-playlist: index=0", recorder.Body.String())
	}
}

func TestVodTsSegment(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/:index/:segment")
	context.SetParamNames("index", "segment")
	context.SetParamValues("0", "0.ts")
	handler := FakeHandler

	err := handler.VodSegment(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMpeg2Ts, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-segment: mediaPlaylistIndex=0, segmentIndex=0", recorder.Body.String())
	}
}

func TestVodFmp4Segment(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/:index/:segment")
	context.SetParamNames("index", "segment")
	context.SetParamValues("0", "0.mp4")
	handler := FakeHandler

	err := handler.VodSegment(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMP4, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-segment: mediaPlaylistIndex=0, segmentIndex=0", recorder.Body.String())
	}
}

func TestLiveMasterPlaylist(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/playlist.m3u8")
	handler := FakeHandler

	err := handler.LiveMasterPlaylist(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMpegUrl, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-master-playlist", recorder.Body.String())
	}
}

func TestLiveMediaPlaylist(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/:index/playlist.m3u8")
	context.SetParamNames("index")
	context.SetParamValues("0")
	handler := FakeHandler

	err := handler.LiveMediaPlaylist(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMpegUrl, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-media-playlist: index=0", recorder.Body.String())
	}
}

func TestLiveTsSegment(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/:index/:segment")
	context.SetParamNames("index", "segment")
	context.SetParamValues("0", "0.ts")
	handler := FakeHandler

	err := handler.LiveSegment(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMpeg2Ts, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-segment: mediaPlaylistIndex=0, segmentIndex=0", recorder.Body.String())
	}
}

func TestLiveFmp4Segment(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/:index/:segment")
	context.SetParamNames("index", "segment")
	context.SetParamValues("0", "0.mp4")
	handler := FakeHandler

	err := handler.LiveSegment(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMP4, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-segment: mediaPlaylistIndex=0, segmentIndex=0", recorder.Body.String())
	}
}

func TestChaseMasterPlaylist(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/playlist.m3u8")
	handler := FakeHandler

	err := handler.ChaseMasterPlaylist(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMpegUrl, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-master-playlist", recorder.Body.String())
	}
}

func TestChaseMediaPlaylist(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/:index/playlist.m3u8")
	context.SetParamNames("index")
	context.SetParamValues("0")
	handler := FakeHandler

	err := handler.ChaseMediaPlaylist(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMpegUrl, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-media-playlist: index=0", recorder.Body.String())
	}
}

func TestChaseTsSegment(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/:index/:segment")
	context.SetParamNames("index", "segment")
	context.SetParamValues("0", "0.ts")
	handler := FakeHandler

	err := handler.ChaseSegment(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMpeg2Ts, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-segment: mediaPlaylistIndex=0, segmentIndex=0", recorder.Body.String())
	}
}

func TestChaseFmp4Segment(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/playlist.m3u8", nil)
	request.Header.Set(echo.HeaderContentType, MimeMpegUrl)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetPath("/vod/:index/:segment")
	context.SetParamNames("index", "segment")
	context.SetParamValues("0", "0.mp4")
	handler := FakeHandler

	err := handler.ChaseSegment(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMP4, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-segment: mediaPlaylistIndex=0, segmentIndex=0", recorder.Body.String())
	}
}
