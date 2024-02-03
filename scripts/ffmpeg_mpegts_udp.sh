ffmpeg -hide_banner -loglevel verbose \
    -re -f lavfi -i testsrc2=size=1280x720:rate=30,format=yuv420p \
    -f lavfi -i sine=frequency=1000:sample_rate=44100 \
    -c:v libx264 -preset veryfast -tune zerolatency -profile:v baseline \
    -b:v 1000k -bufsize 2000k -x264opts keyint=30:min-keyint=30:scenecut=-1 \
    -c:a aac -b:a 128k \
    -f mpegts "udp://${SRT_INPUT_HOST}:${SRT_INPUT_PORT}?pkt_size=${PKT_SIZE}"