package main

import (
	_ "context"
	"flag"
	"fmt"
	"log"
	_ "net"
	nc "nvm-gocore"
	"os"
	"strings"
	tc "trident-core/config"

	"github.com/fatih/color"
	_ "google.golang.org/grpc"
)

type TridentFlags struct {
	Bind             string
	Port             int
	Config           string
	IsDisplayingHelp bool
}

func cli() nc.Option[TridentFlags] {
	// Define flags
	bind := flag.String("bind", "127.0.0.1", "Bind address for server listening, bind value from config will be ignored.")
	port := flag.Int("port", 9089, "Port to listen on, port value from config will be ignored.")
	config := flag.String("config", "trident-config.yaml", "Path to config file.")

	// Define a custom help flag
	help := flag.Bool("help", false, "Display help information")

	// Parse the flags
	flag.Parse()

	// Check if help was requested
	if *help {
		color.Green("Trident Server usage:")
		fmt.Println("")
		flag.PrintDefaults()
		color.Green("\nExample:\n")
		color.Cyan(" trident-server --bind=192.168.1.100 --port=8080 --config=myconfig.yaml\n")

		return nc.Some[TridentFlags](TridentFlags{
			IsDisplayingHelp: true,
		})
	}

	if len(os.Args) == 1 {
		return nc.None[TridentFlags]()
	} else {
		tf := TridentFlags{
			Bind:   *bind,
			Port:   *port,
			Config: *config,
		}

		return nc.Some[TridentFlags](tf)
	}

}

func printConfig() {
	color.Green("trident-config.yaml")
	fmt.Println("")

	trdConfigReader := tc.NewTridentConfig()

	binDir := trdConfigReader.GetBinaryDir()

	path := binDir.Unwrap() + "trident-conf.yaml"
	fileResult := trdConfigReader.OpenConfig(path)

	if fileResult.IsErr() {
		log.Fatalf("error open trident-config.yaml: %v", fileResult.Err)
	}

	configResult := trdConfigReader.ParseConfig(fileResult.Value)

	if configResult.IsErr() {
		log.Fatalf("trident-config.yaml parsed error: %v", fileResult.Err)
	}

	color.New(color.FgCyan).Printf("Config Version: ")
	fmt.Printf("%s\n", configResult.Value.Version)

	color.New(color.FgCyan).Printf("Server Address: ")
	fmt.Printf("%s\n", configResult.Value.Bind)

	color.New(color.FgCyan).Printf("Server Port: ")
	fmt.Printf("%v\n", configResult.Value.Port)

	color.New(color.FgCyan).Printf("Is use ext-auth: ")
	fmt.Printf("%v\n", configResult.Value.IsUseExtAuth)

	for _, service := range configResult.Value.Services {
		color.New(color.FgCyan).Printf("Service Name: ")
		fmt.Printf("Service Name: %s {base-url: %s}\n", service.ServiceName, service.BaseURL)

		allowedOperations := strings.Join(service.Allow, ", ")
		disallowedOperatios := strings.Join(service.Disallow, ", ")
		scopes := strings.Join(service.Scopes, ", ")

		color.New(color.FgCyan).Printf(" -scopes: ")
		fmt.Printf("   [%v]\n", scopes)

		color.New(color.FgCyan).Printf(" -allow: ")
		fmt.Printf("    [%v]\n", allowedOperations)

		color.New(color.FgCyan).Printf(" -disallow: ")
		fmt.Printf(" [%v]\n", disallowedOperatios)

	}

	fmt.Println("")
	color.Yellow("Starting Trident Identity Server on %s:%v...", configResult.Value.Bind, configResult.Value.Port)
	fmt.Println("")
	fmt.Println("")
}

func banner() {
	fmt.Println("")
	color.Yellow("Trident Server")
	color.Cyan("----------------")
	color.Cyan("v0.2.3")
	fmt.Println("")
}

func main() {
	banner()
	cliResult := cli()

	if cliResult.IsSome() {
		if cliResult.Unwrap().IsDisplayingHelp {
			return
		}
	}

	printConfig()
}
