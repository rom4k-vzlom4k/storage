package storage

import (
	"fmt"
	"rom4k-vzlom4k/storage/internal/file"

	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {

	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err // при ошибке, лучше постоянно возвращатиь nil
	}

	s.Files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetById(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[fileID]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileID)
	}

	return foundFile, nil
}
