package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/doingodswork/deflix-stremio/pkg/realdebrid"
)

const (
	bigBuckBunnyMagnet = `magnet:?xt=urn:btih:dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c&dn=Big+Buck+Bunny&tr=udp%3A%2F%2Fexplodie.org%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.empire-js.us%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com&ws=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2F&xs=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2Fbig-buck-bunny.torrent`
)

var (
	apiToken = flag.String("apiToken", "", "RealDebrid API token")
)

func main() {
	ctx := context.Background()
	flag.Parse()

	// Precondition checks
	if *apiToken == "" {
		log.Fatal("apiToken CLI argument must not be empty")
	}

	rdClient := realdebrid.NewClient(ctx, 5*time.Second, nil, nil, time.Duration(0))
	streamURL, err := rdClient.GetStreamURL(ctx, bigBuckBunnyMagnet, *apiToken, false)
	if err != nil {
		log.Fatalf("Couldn't get stream URL: %v", err)
	}
	fmt.Printf("Stream URL retrieved successfully: %v\n", streamURL)
}