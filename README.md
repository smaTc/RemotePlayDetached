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

## Installation
1. Download a donor game for Remote Play on Steam (search on YouTube if you need help).
2. Delete or move the files in the folder of that game (you can just create a subfolder so you can still launch that game from RPD), but remember the original name of the main binary (file extension included!).
3. Copy the RPD binary into the folder and rename it to the name of the original binary.

Remember to set execute permissions on Linux!
Now just launch the original game in Steam and enjoy.

## Usage
1. Click `Import`
2. Fill the form. Path must include executable (and extension on Windows).
3. Click `Run`

You can edit entries and also have arguments applied.

## Testing
Tested on Windows and Linux. Working so far.

## Toubleshooting
If you have problems with Streaming you should check the Steam Settings first. If there are still any problems leave me a message or open an issue here on GitHub.

## Feedback and Contribution
Feedback and Contributions are welcome. If you have improvements just open a pull request and I will have a look at your proposal. If there are any questions or suggestions just leave me a message.

## Email
[smatcx@gmail.com](mailto:smatcx@gmail.com)
