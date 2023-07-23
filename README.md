# music-control
使用golang实现一个移动端控制电脑音乐播放的功能，可以控制音量的增减、暂停、播放、下一首、红心歌曲。

应用场景：在同一局域网下，电脑播放着歌曲而你不想起来切歌

后端使用到了[robotgo](https://github.com/go-vgo/robotgo)、[volume-go](https://github.com/itchyny/volume-go)，前端用到了[fyne](https://github.com/fyne-io/fyne)。项目整体为后端实现业务逻辑绑定路由，前端按钮发起请求

下载服务端和客户端

服务端：server.exe

客户端：control.apk

电脑端运行server.exe启动服务会回显ip及端口

移动端运行，输入ip及端口即可使用
