package utils

import "os"

func OpenFolderAndConsumeFiles(path string, consumer Consumer) {
	fileInfo, err := os.Stat(path)

	if err != nil || !fileInfo.IsDir() {
		consumer.OnError()
	}

	file, err := os.Open(path)

	if err != nil {
		consumer.OnError()
	}

	defer func() {
		_ = file.Close()
	}()

	fileNames, err := file.Readdirnames(0)

	if err != nil {
		consumer.OnError()
	}

	for _, name := range fileNames {
		consumer.OnFileName(name)
	}
}

type Consumer struct {
	OnError    func()
	OnFileName func(fileName string)
}
