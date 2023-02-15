# Notion Control Box Hider

This is a program to hide the control box window of the Notion desktop application. The control box window is the white bar at the top of the application window that contains the close, minimize, and maximize buttons. This bar can be distracting in dark mode, as it remains white even when the rest of the application is dark.

This program modifies the file `WindowController.js` in the Notion application directory to remove the control box window. Specifically, it sets the `frame` property to `false` in the `windowCreationArgs` object, which removes the control box window and allows the application window to take up the entire screen.

## License

This program is released under the MIT License. See the [LICENSE](LICENSE) file for details.

## Requirements

- Golang 1.20 or later

## Usage

To use this program, follow these steps:

1. Clone this repository to your local machine.
2. Open a command prompt or terminal window and navigate to the directory where the repository is located.
3. Run the program with the command `go run notion-extension.go`. This will modify the Notion application file `WindowController.js` to remove the control box window.
4. If you want to revert the changes made by the program, run it with the command `go run notion-extension.go -reset`.

Note: This program modifies the Notion application file, so use it at your own risk. It's always a good idea to make a backup of the file before making any changes.

## Acknowledgements

This program is based on a solution provided by [mattn](https://github.com/mattn) in [this issue](https://github.com/microsoft/vscode/issues/10911#issuecomment-242130904) on the Visual Studio Code GitHub repository. Thanks to mattn for the solution and to the Notion team for creating such a great note-taking application.
