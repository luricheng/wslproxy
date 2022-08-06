# wslproxy
a simple tool for proxying port of linux running in wsl2 (which now use a hyper-v nat network).

## How it works
it make a reverse proxy from host_to_your_win to 127.0.0.1 (in wsl2, you can use 127.0.0.1 to access your wsl2's server)

this commod can make a reverse proxy from 1234(windows's port) to 127.0.0.1:4567(wsl2 server's port)
```
.\wslproxy.exe -f 1234 -t 4567
```