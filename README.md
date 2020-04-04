# go-gp2y1026au0f
Demo code for Sharp GP2Y1026AU0F dust sensor on Raspberry Pi

There's Arduino C demo code on [sharpsensoruser/sharp\-sensor\-demos](https://github.com/sharpsensoruser/sharp-sensor-demos/tree/master/sharp_gp2y1026au0f_demo). I've ported to Go.

## How to run

```
$ GOARM=6 GOARCH=arm GOOS=linux go build -ldflags="-s -w"
$ scp go-gp2y1026au0f raspberrypi:
$ ssh raspberrypi ./go-gp2y1026au0f
```
