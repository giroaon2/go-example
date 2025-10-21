func (s *BlobStorage) Upload(file io.Reader, filename string) error {
    // ใช้ Azure SDK upload ไป blob storage
    containerClient := s.client.NewContainerClient(s.containerName)
    blobClient := containerClient.NewBlockBlobClient(filename)

    _, err := blobClient.UploadStream(context.Background(), file, nil)
    return err
}
