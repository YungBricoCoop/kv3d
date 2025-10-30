/*
Copyright © 2025 Elwan Mayencourt <mayencourt@elwan.ch>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kvd",
	Short: "A dumb Redis compatible key-value database using Docker as storage",
	Long: `
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""\___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
         \    \        __/
          \____\______/

         _  ____   _____
        | |/ /\ \ / /   \
        | ' <  \ V /| |) |
        |_|\_\  \_/ |___/

────────────────────────────────────────────────────────────────
 KVD — Key Value Docker - dumb docker-powered key-value database
───────────────────────────────────────────────────────────────
	Author:  Elwan Mayencourt <mayencourt@elwan.ch>

Description:
  kvd is a dumb key-value database that uses Docker containers
  as its storage engine. it speaks the RESP protocol, compatible with
  basic redis clients and commands. Why "dumb"? Because it's one of the worst
	possible way to implement a key-value store, very inefficient but fun.


Requirements:
  • Docker must be installed and running on the host machine

Examples:
  kvd serve --port 6379
  redis-cli -p 6379
    > SET mykey myvalue
    > GET mykey

`,
	Run: func(cmd *cobra.Command, args []string) {
		// Display the Long description when run without subcommands
		cmd.Println(cmd.Long)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Child commands are added in their respective files (e.g., serve.go)
}
