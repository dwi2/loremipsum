package main

import (
    "os"
    "github.com/urfave/cli"
    "github.com/atotto/clipboard"
)

func main() {

    app := cli.NewApp()
    app.Name = "loremipsum"
    app.Usage = "copy lorem ipsum text to your clipboard"
    app.Run(os.Args)

    clipboard.WriteAll("Lorem ipsum dolor sit amet, consectetur adipiscing elit," +
" sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad" +
" minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea" +
" commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit" +
" esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat" +
" non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
}
