<p align="center"><a href="https://github.com/portapps/twitch-portable" target="_blank"><img width="100" src="https://github.com/portapps/twitch-portable/blob/master/res/papp.png"></a></p>

<p align="center">
  <a href="https://github.com/portapps/twitch-portable/releases/latest"><img src="https://img.shields.io/github/release/portapps/twitch-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/portapps/twitch-portable/releases/latest"><img src="https://img.shields.io/github/downloads/portapps/twitch-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/portapps/twitch-portable"><img src="https://img.shields.io/appveyor/ci/portapps/twitch-portable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://goreportcard.com/report/github.com/portapps/twitch-portable"><img src="https://goreportcard.com/badge/github.com/portapps/twitch-portable?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/portapps/twitch-portable"><img src="https://img.shields.io/codacy/grade/b4286f4c64ba4788915b163a560b1c42.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://beerpay.io/portapps/portapps"><img src="https://img.shields.io/beerpay/portapps/portapps.svg?style=flat-square" alt="Beerpay"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

[Twitch](https://app.twitch.tv/) portable app made with 🚀 [Portapps](https://github.com/portapps).<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

## Installation

There are different kinds of artifacts :

* `twitch-portable-win{32,64}-x.x.x-x-setup.exe` : Full portable release of Twitch as a setup. **Recommended way**!
* `twitch-portable-win{32,64}-x.x.x-x.7z` : Full portable release of Twitch as a 7z archive.
* `twitch-portable-win{32,64}.exe` : Only the portable binary (must be renamed `twitch-portable.exe`)
* `TwitchSetup.exe` : Can be found in the [untouched repository](https://github.com/portapps/untouched/releases). 

### Fresh installation

Install `twitch-portable-win{32,64}-x.x.x-x-setup.exe` where you want then run `twitch-portable.exe`.

### App already installed

If **you have already installed Twitch from the original setup**, do the same thing as a fresh installation and move :

* Content of folder where Twitch is installed (default to `%APPDATA%\Twitch\*`) to the root folder except `Bin` folder.

Run `twitch-portable.exe` and then you can [remove](https://support.microsoft.com/en-us/instantanswers/ce7ba88b-4e95-4354-b807-35732db36c4d/repair-or-remove-programs) Twitch from your computer.

### Upgrade

**For an upgrade**, simply download and install the [latest twitch-portable setup](https://github.com/portapps/twitch-portable/releases/latest).

## How can i help ?

All kinds of contributions are welcomed :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
But we're not gonna lie to each other, I'd rather you buy me a beer or two :beers:!

[![Beerpay](https://beerpay.io/portapps/portapps/badge.svg?style=beer-square)](https://beerpay.io/portapps/portapps)
or [![Paypal](https://cdn.rawgit.com/portapps/portapps/master/res/paypal.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG)

## License

MIT. See `LICENSE` for more details.<br />
Rocket icon credit to [Squid Ink](http://thesquid.ink).
