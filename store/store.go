package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

// FileMeta: 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

type FileMetaStore interface {
	Create(*FileMeta) error
	Update(*FileMeta) error
	Get(string) (FileMeta, error)
	GetAll() ([]FileMeta, error)
	Delete(string) error
}
