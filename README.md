# go-hls-hosting

![test](https://github.com/TakuSemba/go-hls-hosting/workflows/test/badge.svg)
![release](https://github.com/TakuSemba/go-hls-hosting/workflows/release/badge.svg)


**go-hls-hosting** generates **LIVE**, **CHASE** playlists from your given **VOD** playlist, and allows you to access to those **VOD**, **LIVE**, **CHASE** playlists individually.

gif...

## Get Started

To get started, run main.go with master playlist path and listenen port.
If you haven't prepared your master playlist, you can use one in [sampledata](https://github.com/TakuSemba/go-hls-hosting/tree/master/sampledata) directory.

```sh
go run main.go --path your/hls/master/playlist.m3u8 --port 1323
```

Once you've run it, you can access **VOD**, **LIVE**, **CHASE** playlists respectively with the following path.

| Type | master playlist | media playlist | segment |
|:---|:---|:---|:---|
| VOD | /vod/playlist.m3u8 | /vod/:index/playlist.m3u8 | /vod/:index/:path |
| LIVE | /live/playlist.m3u8 | /live/:index/playlist.m3u8 | /live/:index/:path |
| CHASE | /chase/playlist.m3u8 | /chase/:index/playlist.m3u8 | /chase/:index/:path |

### With Docker

You can also start **go-hls-hosting** using Docker.

### With Binary

You can download binary from [release](https://github.com/TakuSemba/go-hls-hosting/releases) page.

## Integrate with ngrok

With [ngrok](https://ngrok.com/), you can publish playlists.


gif...


## Author

* **Taku Semba**
    * **Github** - (https://github.com/takusemba)
    * **Twitter** - (https://twitter.com/takusemba)
    * **Facebook** - (https://www.facebook.com/takusemba)

## Licence
```
Copyright 2017 Taku Semba.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
