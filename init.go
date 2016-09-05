package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/plugins/fs"
)

var (
	// Map from plugin name to object
	pluginMap = make(map[string]plugins.RikkaPlugin)

	// Command line arguments var
	argBindIPAddress *string
	argPort          *int
	argPassword      *string
	argMaxSizeByMB   *float64
	argPluginStr     *string
	argLogLevel      *int

	// concat socket from ip address and port
	socket string

	// The used plugin
	thePlugin plugins.RikkaPlugin
)

// --- Init and check functions ---

func createSignalHandler(handlerFunc func()) (func(), chan os.Signal) {
	signalChain := make(chan os.Signal, 1)

	return func() {
		for _ = range signalChain {
			handlerFunc()
		}
	}, signalChain
}

// registerSignalHandler register a handler for process Ctrl + C
func registerSignalHandler(handlerFunc func()) {
	signalHandler, channel := createSignalHandler(handlerFunc)
	signal.Notify(channel, os.Interrupt)
	go signalHandler()
}

func init() {

	registerSignalHandler(func() {
		l.Fatal("Rikka have to go to sleep, see you tomorrow")
	})

	initPluginList()

	initArgVars()

	flag.Parse()

	l.Info("Args bindIP =", *argBindIPAddress)
	l.Info("Args port =", *argPort)
	l.Info("Args password =", *argPassword)
	l.Info("Args maxFileSize =", *argMaxSizeByMB, "MB")
	l.Info("Args loggerLevel =", *argLogLevel)
	l.Info("Args.plugin =", *argPluginStr)

	if *argBindIPAddress == ":" {
		socket = *argBindIPAddress + strconv.Itoa(*argPort)
	} else {
		socket = *argBindIPAddress + ":" + strconv.Itoa(*argPort)
	}

	logger.SetLevel(*argLogLevel)

	runtimeEnvCheck()
}

func initPluginList() {
	pluginMap["fs"] = fs.FsPlugin
}

func initArgVars() {
	argBindIPAddress = flag.String("bind", ":", "bind ip address, use : for all address")
	argPort = flag.Int("port", 80, "server port")
	argPassword = flag.String("pwd", "rikka", "The password need provided when upload")
	argMaxSizeByMB = flag.Float64("size", 5, "Max file size by MB")
	argLogLevel = flag.Int(
		"level", logger.LevelDebug,
		fmt.Sprintf("logger level, from %d to %d", logger.LevelDebug, logger.LevelError),
	)

	// Get name array of all avaliable plugins, show in `rikka -h``
	pluginNames := make([]string, 0, len(pluginMap))
	for k := range pluginMap {
		pluginNames = append(pluginNames, k)
	}
	argPluginStr = flag.String(
		"plugin", "fs",
		"what plugin use to save file, selected from "+fmt.Sprintf("%v", pluginNames),
	)

}

func runtimeEnvCheck() {
	l.Info("Check runtime environment")

	l.Debug("Try to find plugin", *argPluginStr)

	// Make sure plugin be selected exist
	if plugin, ok := pluginMap[*argPluginStr]; ok {
		thePlugin = plugin
		l.Debug("Plugin", *argPluginStr, "found")
	} else {
		l.Fatal("Plugin", *argPluginStr, "not exist")
	}

	l.Info("All runtime environment check passed")
}