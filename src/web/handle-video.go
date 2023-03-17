package web

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

const bufferSize = 64 * 1024

func HandleVideo(res http.ResponseWriter, req *http.Request) {
	chunkSize := 2 * 1024 * 1024 // 2 MB
	filename := mux.Vars(req)["filename"]
	videoPath := fmt.Sprintf("assets/videos/%s.mp4", filename)
	stat, err := os.Stat(videoPath)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	header := res.Header()

	header.Set("Content-Type", "video/mp4")
	header.Set("Connection", "Keep-Alive")
	header.Set("Keep-Alive", "timeout=5")

	fileSize := stat.Size()
	videoRange := req.Header.Get("Range")
	if len(videoRange) == 0 {
		filenameToDownload := "2 girls, 1 cop.mp4"
		header.Set("Content-Length", fmt.Sprint(fileSize))
		header.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filenameToDownload))
		file, _ := os.Open(videoPath)
		defer file.Close()

		io.Copy(res, bufio.NewReaderSize(file, bufferSize))
		return
	}

	parts := strings.Split(strings.Replace(videoRange, "bytes=", "", 1), "-")
	start, _ := strconv.Atoi(parts[0])
	var end int
	if len(parts) == 2 && parts[1] != "" {
		end, _ = strconv.Atoi(parts[1])
	} else {
		end = int(math.Min(float64(fileSize)-1, float64(start+chunkSize)))
	}

	chunkSize = end - start + 1
	header.Set("Content-Length", fmt.Sprint(chunkSize))
	header.Set("Content-Range", fmt.Sprintf("bytes %v-%v/%v", start, end, fileSize))
	file, _ := os.Open(videoPath)
	defer file.Close()

	file.Seek(int64(start), 0)
	res.WriteHeader(206)

	io.Copy(res, io.LimitReader(bufio.NewReaderSize(file, bufferSize), int64(chunkSize)))
}
