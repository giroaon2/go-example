func (u *BlobUsecase) Upload(file *multipart.FileHeader) error {
    // แปลง file เป็น byte stream หรือเปิดไฟล์
    f, err := file.Open()
    if err != nil {
        return err
    }
    defer f.Close()

    // เรียก storage layer เพื่ออัปโหลดไป Azure Blob
    return u.storage.Upload(f, file.Filename)
}
