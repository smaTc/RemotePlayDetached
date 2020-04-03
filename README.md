# Remote Play Detached (RPD)

## About
A simple launcher that is able to launch any external game or application with Steam Remote Play Together Support (Steam Overlay works as soon as the Fullscreen is available).

I used [Fyne](https://fyne.io/) for the GUI.

I created this Launcher to be able to easily start different non-Steam applications and use the remote play feature without copying files around all the time.
Works pretty well with applications like [Redream](https://redream.io/) and [RetroArch](https://www.retroarch.com/)

## Building
1. Install Go on your System.
2. Clone the git repo.
3. Navigate to that directory and execute `go run main.go` and let it pull the dependencies.
4. Execute `go build -o YourPreferredName.DesiredExtension main.go`

For Cross Compiling you can use [Fyne-Cross](https://fyne.io//develop/cross-compiling.html) and execute the `build.sh` File. (Linux only probably)

## Download
You can find prebuilt binaries for Windows and Linux here on [GitHub](https://github.com/smaTc/RemotePlayDetached/releases/).

## Installation
1. Download a donor game for Remote Play on Steam (search on YouTube if you need help).
2. Delete or move the files in the folder of that game (you can just create a subfolder so you can still launch that game from RPD), but remember the original name of the main binary (file extension included!).
3. Copy the RPD binary into the folder and rename it to the name of the original binary.

Remember to set execute permissions on Linux!
Now just launch the original game in Steam and enjoy.

## Usage
1. Start your donor game over Steam
2. Click `Import`
3. Fill the form. Path must include executable (and extension on Windows).
4. Click `Run`

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
* Linux 
    * I don't know how to pass the arguments to the game at the moment. I will update this as soon as I know.

## Testing
Tested on Windows and Linux. Working so far.

## Toubleshooting
If you have problems with Streaming you should check the Steam Settings first. If there are still any problems leave me a message or open an issue here on GitHub.

## Feedback and Contribution
Feedback and Contributions are welcome. If you have improvements just open a pull request and I will have a look at your proposal. If there are any questions or suggestions just leave me a message.

## Email
[smatcx@gmail.com](mailto:smatcx@gmail.com)
