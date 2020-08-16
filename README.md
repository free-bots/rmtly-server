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

## Build
Run `make` or `make build` and at `<project directory>/bin` the binary called rmtly-server can be used.

## Run
After building the project run `make run` or open the rmtly-server file in the build directory.

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
to `http://localhost:3000/authentication/signUp` 
after that you get an jwt and add it to your headers for the requests

### Configuration
The config file for this server is located at `.config/rmtly-server/config.json` this file is auto generated if not available.

#### Default Config
```
{
  "application": { "cacheExpiresInMillis": 10000 },
  "image": {
    "cacheExpiresInMillis": 10000,
    "maxImagesInCache": 100, //number of images in the cache
    "imageQuality": 512 // size in pixels
  },
  "security": {
    "expirationInDays": 99, // jwt expiration time
    "secret": "authenticationCode",
    "keyFile": ""
  },
  "network": { "address": "0.0.0.0:3000" }
}
```


## License
> MIT License

> Copyright (c) 2020 free-bots

> Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
