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

### GET `/preview` 
视频流预览页面

### GET `/preview/[streamPath]?type=[hdl|hls|ws|wt|rtc]`
预览指定的视频流，比如 `preview/live/test?type=ws`, 使用 ws 方式预览 live/test 直播流 

### GET `/preview/[filepath]`
（开发中）预览制定的录像文件，比如 `preview/record/flv/xxx.flv` 

## 使用WebTransport注意事项

- 本地测试需要本地启动https服务，并配置有效的证书
- 由于证书与域名绑定，所以需要host里面配置对应的域名 例如：`127.0.0.1  monibuca.com`
