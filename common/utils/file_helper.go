package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func FileSizeFormat(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

func ReadLastNLines(filename string, n int) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	var (
		lines     []string
		lineCount int
		buffer    []byte
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
		buffer = append(buffer, scanner.Bytes()...)
		buffer = append(buffer, '\n')
	}
	
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	startIndex := lineCount - n
	if startIndex < 0 {
		startIndex = 0
	}

	reader := bufio.NewReader(bytes.NewReader(buffer))
	for i := 0; i < lineCount; i++ {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			return nil, err
		}

		if i >= startIndex {
			lines = append(lines, line)
		}

		if err == io.EOF {
			break
		}
	}

	return lines, nil
}

func MatchSearchFileFromDir(dirPath string, search string) ([]string, error) {
	result := make([]string, 0)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return result, err
	}

	for _, entry := range entries {
		if strings.Contains(entry.Name(), search) {
			filePath := filepath.Join(dirPath, entry.Name())
			result = append(result, filePath)
		}
	}

	return result, nil
}
