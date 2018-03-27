#!/bin/bash

echo "Recording 10 differnt files"

sudo ffmpeg -f v4l2 -framerate 25 -video_size 640x480 -i /dev/video0 \
	-s 640x480  output1.mp4 \
	-s 640x480  output2.mp4 \
	-s 640x480  output3.mp4 \
	-s 640x480  output4.mp4 \
	-s 640x480  output5.mp4 \
	-s 640x480  output6.mp4 \
	-s 640x480  output7.mp4 \
	-s 640x480  output8.mp4 \
	-s 640x480  output9.mp4 \
	-s 640x480  output10.mp4


