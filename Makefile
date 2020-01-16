build:
	go build main.go

run:
	go run main.go

send:
	ffmpeg -v verbose -re -f lavfi -i testsrc -f lavfi -i "sine=frequency=440:sample_rate=44100:beep_factor=8" -strict -2 -c:a aac -b:a 128k -ar 44100 -r 30 -g 60 -keyint_min 60 -b:v 400000 -c:v libx264 -preset medium -bufsize 800k -maxrate 400k -f flv rtmp://localhost

play:
	ffplay rtmp://localhost/
