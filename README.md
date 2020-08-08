```
                _   _
               | | | |
 _ __ _ __ ___ | |_| |_   _ ______ ___  ___ _ ____   _____ _ __
| '__| '_ ` _ \| __| | | | |______/ __|/ _ \ '__\ \ / / _ \ '__|
| |  | | | | | | |_| | |_| |      \__ \  __/ |   \ V /  __/ |
|_|  |_| |_| |_|\__|_|\__, |      |___/\___|_|    \_/ \___|_|
                       __/ |
                      |___/
```

## About
rmtly-server is a server application written in golang that allows you to start application with a rest api


## Requirements
Ubuntu
```
sudo apt install libgtk-3-dev libcairo2-dev libglib2.0-dev libnotify-bin 
```
Arch Linux
```
sudo pacman -S gtk3 libnotify
```

## API
### Get all applications
```
http://localhost:3000/applications
```
### Get single application
```
http://localhost:3000/applications/org.gnome.gedit.desktop
```
### Run an application
```
http://localhost:3000/applications/run/org.gnome.gedit.desktop
```
### Get icon of an application
```
http://localhost:3000/applications/org.gnome.gedit.desktop/icon
```
### SignUp a device
```
http://localhost:3000/authentication/code
```
after requesting the code go to the terminal and scan the created qr code or type it then post a json with that code and an id 
```json
{
  "qrCode": "code",
  "deviceId": "random string"
}
```
to 
```
http://localhost:3000/authentication/signUp
``` 
after that you get an jwt and add it to your headers for the requests

### Configuration
The config file for this server is located at ```.config/rmtly-server/config.json``` this file is auto generated if not available.