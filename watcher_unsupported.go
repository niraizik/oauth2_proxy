// +build !go1.3 plan9 solaris

package main

import "github.com/oauth2-proxy/oauth2-proxy/pkg/logger"

func WatchForUpdates(filename string, done <-chan bool, action func()) {
	logger.Printf("file watching not implemented on this platform")
	go func() { <-done }()
}
