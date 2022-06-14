package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	DefaultPort = "8080"
)

func serveFiles(port, dir string) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	port := "0.0.0.0:" + cmd.Flag("port").Value.String()
	dir := "./"
	if len(args) > 0 {
		dir = args[0]
	}
	fmt.Println("Serving files from ", dir, "on port ", port)
	fmt.Println("Press Ctrl+C to stop")
	serveFiles(port, dir)
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "httpdir",
		Short: "Serve files from the current directory",
		Long:  "Simple http server that serves files from directory and sub directories",
		Run:   run,
	}
	rootCmd.Flags().StringP("port", "p", DefaultPort, "port to listen on")
	rootCmd.Execute()
}
