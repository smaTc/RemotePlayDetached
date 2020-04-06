# Remote Play Detached (RPD) 
If you like my work and want to sponsor some beers for me, feel free to donate.
[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=U6DXYQPU7JLJL)

## Discord
I set up a [Discord Server](https://discord.gg/U3zDs6N) for anyone who wants to give feedback, suggestions or contributions.

## About
A simple launcher that is able to launch any external game or application with Steam Remote Play Together Support (Steam Overlay works as soon as the Fullscreen is available).

I used [Fyne](https://fyne.io/) for the GUI.

## Features
* Enable Remote Play Together for non-Steam applications
* Enable Remote Play Together for Steam applications that normally don't support (e.g. added later via mod, thanks to [AkiraJkr](https://github.com/AkiraJkr) for pointing that out)
* CLI Support for integrations in other applications (e.g. Playnite, thanks to [darklinkpower](https://github.com/darklinkpower) for suggesting it)
* Importing/Editing your applications in a list with the option to append launch arguments


I created this Launcher to be able to easily start different non-Steam applications and use the Remote Play Together feature without copying files around all the time.
Works pretty well with applications like [Redream](https://redream.io/) and [RetroArch](https://www.retroarch.com/). Besides it seems that RPD also enables that feature for other Steam games if they are started with it. An example for this is Devil May Cry 4 with the [DDMK](https://github.com/serpentiem/ddmk)

## Building
1. Install Go on your System.
2. Clone the git repo.
3. Navigate to that directory and execute `go run main.go` and let it pull the dependencies.
4. Execute `go build -o YourPreferredName.DesiredExtension main.go`.

For Cross Compiling you can use [Fyne-Cross](https://fyne.io//develop/cross-compiling.html) and execute the `build.sh` File. (Linux only probably)

## Download
You can find prebuilt binaries for Windows and Linux here on [GitHub](https://github.com/smaTc/RemotePlayDetached/releases/).

## Installation
1. Download a donor game for Remote Play on Steam (search on YouTube if you need help).
2. Delete or move the files in the folder of that game (you can just create a subfolder so you can still launch that game from RPD), but remember the original name of the main binary (file extension included!).
3. Copy the RPD binary into the folder and rename it to the name of the original binary.

### Linux only
* Instead of doing the installation as stated above you can use the launch options to point directly to the RPD binary
* To do this add `/path/to/RPD_binary # %command%`

Remember to set execute permissions on Linux!
Now just launch the original game in Steam and enjoy.

## Usage
1. Start your donor game over Steam.
2. Click `Import`.
3. Fill the form. Path must include executable (and extension on Windows).
4. Click `Run`.

You can edit entries and also have arguments applied.

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
    * To launch an app via Terminal execute `steam steam://run//<appID>//"-arg <App or Path>"/`
    * Unfortunately there is a warning that needs confirmation pops up because of the argument

### Playnite
To start an app directly via [Playnite](https://playnite.link/) do the following:
1. Click on the Controller icon in the upper left corner.
2. Click `Add Game`->`Manually...`.
3. Fill in the Name in the `General` Tab.
4. In the `Actions` Tab click `Add Action` and fill out the Path to Steam and put this in Arguments `-applaunch <app-id of donor Game> -arg <App or Path>`.
5. Click `Save` and you're ready to go.

### AppID
You can find the AppID of your donor game on [SteamDB](https://steamdb.info/).

## Status
Tested on Windows and Linux. Working so far.

## Toubleshooting
* If you have problems with Streaming you should check the Steam Settings first. If there are still any problems leave me a message or open an issue here on GitHub
* If you encounter the error `Steam Error: application load error V:xxxxxxxxx` it means Steam is blocking the execution of that game if it is initiated by another Steam game

## Feedback and Contribution
Feedback and Contributions are welcome. If you have improvements just open a pull request and I will have a look at your proposal. If there are any questions, feedback or suggestions just leave me a message or join the [Discord Server](https://discord.gg/U3zDs6N).

## Email
[smatcx@gmail.com](mailto:smatcx@gmail.com)
