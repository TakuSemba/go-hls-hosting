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
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/vod/0/0.ts", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/vod/0/0.mp4", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/live/playlist.m3u8", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/live/0/playlist.m3u8", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/live/0/0.ts", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/live/0/0.mp4", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/chase/playlist.m3u8", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/chase/0/playlist.m3u8", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/chase/0/0.ts", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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
	request := httptest.NewRequest(http.MethodGet, "/chase/0/0.mp4", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
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

func TestM4sSegment(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/0.m4s", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetParamNames("index", "segment")
	context.SetParamValues("0", "0.m4s")
	handler := FakeHandler

	err := handler.VodSegment(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, MimeMP4, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "test-segment: mediaPlaylistIndex=0, segmentIndex=0", recorder.Body.String())
	}
}

func TestUnknownSegment(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/vod/0/0.aaa", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	context.SetParamNames("index", "segment")
	context.SetParamValues("0", "0.aaa")
	handler := FakeHandler

	err := handler.VodSegment(context)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, echo.MIMETextPlainCharsetUTF8, recorder.Header().Get(echo.HeaderContentType))
		assert.Equal(t, "failed to load Segment.", recorder.Body.String())
	}
}
