#!/bin/bash
# This script is using for tests and requires VLC to be installed
vlc example.mp3 --sout '#transcode{acodec=mp3,ab=192}:standard{access=http,mux=ogg,dst=:8080}'