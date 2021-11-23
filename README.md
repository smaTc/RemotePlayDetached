# Remote Play Detached (RPD) 
If you like my work and want to sponsor some beers for me, feel free to donate.
[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=U6DXYQPU7JLJL)

## Disclaimer
Don't join the Discord server if you did not read the README below, especially the troubleshooting section. If you're trying to do something like letting your friend play PS Remote/watch a movie stream or anything similar, do not ask for help.

## About
A simple launcher that is able to launch any external game or application with Steam Remote Play Together Support (Steam Overlay works as soon as the Fullscreen is available).

I used [Fyne](https://fyne.io/) for the GUI.

## Features
* Enable Remote Play Together for non-Steam applications
* Enable Remote Play Together for Steam applications that normally don't support (e.g. added later via mod, thanks to [AkiraJkr](https://github.com/AkiraJkr) for pointing that out)
* Basic File Explorer for binary selection
* CLI Support for integrations in other applications (e.g. Playnite, thanks to [darklinkpower](https://github.com/darklinkpower) for suggesting it)
* Importing/Editing your applications in a list with the option to append launch arguments

I created this Launcher to be able to easily start different non-Steam applications and use the Remote Play Together feature without copying files around all the time.
Works pretty well with applications like [Redream](https://redream.io/) and [RetroArch](https://www.retroarch.com/). Besides it seems that RPD also enables that feature for other Steam games if they are started with it. An example for this is Devil May Cry 4 with the [DDMK](https://github.com/serpentiem/ddmk)

## Download
You can find prebuilt binaries for Windows and Linux here on [GitHub](https://github.com/smaTc/RemotePlayDetached/releases/).

For Arch Linux or any other distro that is built upon Arch users, you can find the package in the [aur](https://aur.archlinux.org/packages/remoteplaydetached-bin/). You can skip the Installation process. Note that when you link the executable as described in Linux only you have to point to /usr/bin/remoteplaydetached, because the executable is being installed there. You could also create a link wherever you like it to be.
The aur package is not being maintained by the original author, but by @alx365 (discord: !LegendOfMiracles#1719)/

## Installation
1. Download a donor game for Remote Play on Steam (search on YouTube if you need help). Choose a game which has Remote Play Together enabled.
2. Delete or move the files in the folder of that game (you can just create a subfolder so you can still launch that game from RPD), but remember the original name of the main binary (file extension included!).
3. Copy the RPD binary into the folder and rename it to the name of the original binary.

Note: A `donor game` is a game installed via Steam of which the binaries are replaced. It "donates" its Steam privileges for RPD. 

### Linux only
* Instead of doing the installation as stated above you can use the launch options to point directly to the RPD binary
* To do this add `/path/to/RPD_binary # %command%`

Remember to set execute permissions on Linux!
Now just launch the original game in Steam and enjoy.

## Usage
1. Start your donor game over Steam.
2. Click `Import`.
3. Fill the form. Path must include executable (and extension on Windows). You can select an executable via the `File Explorer` (since v0.4).
4. Click `Run`.

You can edit entries and also have arguments applied.

Note: In the `File Explorer` you need to click on a file to select it (the selected file is registerd and can be seen in the window bottom next to `Selected:`) and finish your selection by clicking the `Confirm` button. Also directories are marked with `/`(Linux) or `\`(Windows) and you can change your drive(Windows) or root directory(Linux) from a dropdown by clicking `Drive or Root Directory` in the window top.

## Command Line Support
Command Line Support was added in v0.2
### Arguments
* `-s` or `-silence` to disable the GUI (cannot be used alone)
* `-a` or `-app` to run an app from your list by its name
* `-as` or `-appsilent` to run an app from your list without GUI
* `-r` or `-run` to run an app from the given path
* `-rs` or `-runsilent` to run an app from the given path without GUI
* `-h` or `-help` to list all possible arguments

### Start an app completely via Command Line
* Windows
    * To launch an app via CMD execute `start "" /d "Drive:\Path\to\SteamFolder" Steam.exe -applaunch <app-id of donor Game> -arg <App or Path>`
* Linux (Workaround)
    * To launch an app via Terminal execute `steam steam://run/<appID>//"-arg <App or Path>"/`
    * Unfortunately there is a warning that needs confirmation pops up because of the argument

### Playnite
To start an app directly via [Playnite](https://playnite.link/) do the following:
1. Click on the Controller icon in the upper left corner.
2. Click `Add Game`->`Manually...`.
3. Fill in the Name in the `General` Tab.
4. In the `Actions` Tab click `Add Action` and fill out the Path to Steam and put this in Arguments `-applaunch <app-id of donor Game> -arg <App or Path>`.
5. Click `Save` and you're ready to go.

### Windows Apps/Xbox Games
To use Windows Apps or Xbox games with RPD requires the use of `UWPHook`.
1. Follow the instructions to add your games to `Steam` with [UWPHook](https://github.com/BrianLima/UWPHook).
2. Right click the game in `Steam` and click `Properties`.
3. Take note of the the `UWPHook` file path and launch arguments.
4. Run RPD as normal and click `Import`.
5. Enter the exe location of `UWPHook` as the `file path`.
6. Provide the launch arguments given on the `Properties` page.
7. Run the game through `RPD`.

* Example: 

Name: Streets of Rage 4
Path: C:\Program Files (x86)\Briano\UWPHook\UWPHook.exe
Args: DotEmu.StreetsofRage4_map6zyh9ym1xy!App

*Warning: Users reported that UWPHook does not work for many games with RPD* 

### AppID
You can find the AppID of your donor game on [SteamDB](https://steamdb.info/).

## Building
### Linux
1. Install Go on your System.
2. Clone this git repo outside of your GOPATH or pull via go get with `go get -v github.com/smaTc/RemotePlayDetached.git`.
3. Navigate to that directory and execute `go mod download` if you pulled via git or `go get .` if you pulled via `go get` and let it pull the dependencies.
4. Run `go run main.go` to check if everything works.
5. Execute `go build -o YourPreferredName.DesiredExtension main.go`.

For Cross Compiling you can use [Fyne-Cross](https://fyne.io//develop/cross-compiling.html) and execute the `build.sh` File. (Linux only probably)

## Status
Tested on Windows and Linux. Working so far.

## Toubleshooting
* If you have problems with Streaming you should check the Steam Settings first. If there are still any problems join the Discord or open an issue here on GitHub
* If you encounter the error `Steam Error: application load error V:xxxxxxxxx` it means Steam is blocking the execution of that game if it is initiated by another Steam game
* Please beware if you have some games running and others don't it probably means that this is a game related issue. Many games just don't work because other launchers (through framework integration or similar) or the game itself block the correct execution via RPD and this is something that cannot be fixed.

### My friend can only see the RemotePlayDetached process and not the game
* Make sure if you can open the Steam Overlay in the game you started via RDP.
* If this is not the case, your SteamService might not be working correctly. To fix this do the following (for Windows, Linux users need to adapt it for their platform):
1. Close Steam completely (right-click the tray icon and choose "exit")
2. Press WIN + R
3. type `"PATH\TO\STEAM_Folder\bin\SteamService.exe" /repair` 
4. Restart Steam

## Feedback, suggestions and contribution
Feedback, suggestions and contributions are welcome. If you have improvements just open a pull request and I will have a look at your proposal. If there are any questions, feedback or suggestions open an issue or leave me a message.

## Contact
[smatcx@gmail.com](mailto:smatcx@gmail.com)
