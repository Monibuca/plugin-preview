# 预览插件

用于预览本地视频播放器插件

## 插件地址

https://github.com/Monibuca/plugin-preview

## 插件引入

```go
import (
    _ "m7s.live/plugin/preview/v4"
)
```

## 配置

无

## API

- `/preview/[streamPath]?type=[hdl|hls|ws|wt|rtc]` 可用于预览直播流
- `/preview/[filepath]` 可用于预览录像文件
