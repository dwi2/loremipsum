package main

import (
    "fmt"
    "os"
    "github.com/urfave/cli"
    "github.com/atotto/clipboard"
)

func generate(command string) {
    text :=
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, " +
        "sed do eiusmod tempor incididunt ut labore et dolore magna " +
        "aliqua. Ut enim ad minim veniam, quis nostrud exercitation " +
        "ullamco laboris nisi ut aliquip ex ea commodo consequat. " +
        "Duis aute irure dolor in reprehenderit in voluptate velit " +
        "esse cillum dolore eu fugiat nulla pariatur. " +
        "Excepteur sint occaecat cupidatat non proident, sunt in culpa " +
        "qui officia deserunt mollit anim id est laborum."

    switch command {
    case "short":
        text = "Lorem ipsum dolor sit amet, " +
            "consectetur adipiscing elit, sed do eiusmod tempor " +
            "incididunt ut labore et dolore magna aliqua."
    case "tiny":
        text = "Lorem ipsum"
    default:
        // do nothing
    }
    clipboard.WriteAll(text)
    fmt.Println(text)
}

func main() {
    app := cli.NewApp()
    app.Name = "loremipsum"
    app.Version = "0.0.1"
    app.Usage = "Generate `Lorem ipsum` in your clipboard in no time"
    app.Commands = []cli.Command{
        {
            Name: "short",
            Aliases: []string{"s"},
            Usage: "Generate first sentence of `Lorem ipsum` in clipboard",
            Action: func(c *cli.Context) error {
                generate("short")
                return nil
            },
        },
        {
            Name: "tiny",
            Aliases: []string{"t"},
            Usage: "Generate only `Lorem ipsum` in clipboard",
            Action: func(c *cli.Context) error {
                generate("tiny")
                return nil
            },
        },
    }

    app.Action = func(c *cli.Context) error {
        generate("full")
        return nil;
    }

    app.Run(os.Args)
}
