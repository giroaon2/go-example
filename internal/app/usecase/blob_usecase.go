package usecase

import (
    "mime/multipart"
    "phillip-cms-jobs/internal/infrastructure/storage"
)

type BlobUsecase struct {
    storage *storage.BlobStorage
}

func NewBlobUsecase() *BlobUsecase {
    return &BlobUsecase{
        storage: storage.NewBlobStorage(), // สร้าง storage
    }
}

func (u *BlobUsecase) Upload(file *multipart.FileHeader) error {
    f, err := file.Open()
    if err != nil {
        return err
    }
    defer f.Close()

    return u.storage.Upload(f, file.Filename)
}
