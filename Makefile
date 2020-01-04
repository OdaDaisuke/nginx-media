build:
	go build cmd/main.go

run:
	go run cmd/main.go

send:
	ffmpeg -re -i statics/test/sample.mp4 -c copy -f flv rtmp://localhost/

play:
	ffplay rtmp://localhost/
