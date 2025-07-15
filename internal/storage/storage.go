package storage

import "rom4k-vzlom4k/storage/internal/file"

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	return file.NewFile(filename, blob)
	// if err != nil {
	// 	return nil, err // При ошибке, лучше постоянно возвращатиь nil
	// }

	// return newFile
}
