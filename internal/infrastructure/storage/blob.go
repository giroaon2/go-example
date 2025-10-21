package storage

import (
    "context"
    "io"
    "log"

    "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type BlobStorage struct {
    client        *azblob.ServiceClient
    containerName string
}

func NewBlobStorage() *BlobStorage {
    // ตัวอย่างสร้าง client Azure (สมมติมี endpoint + credential)
    client, err := azblob.NewServiceClientWithNoCredential("https://<your_account>.blob.core.windows.net/", nil)
    if err != nil {
        log.Fatalf("Failed to create blob client: %v", err)
    }

    return &BlobStorage{
        client:        client,
        containerName: "mycontainer",
    }
}

func (s *BlobStorage) Upload(file io.Reader, filename string) error {
    containerClient := s.client.NewContainerClient(s.containerName)
    blobClient := containerClient.NewBlockBlobClient(filename)

    _, err := blobClient.UploadStream(context.Background(), file, nil)
    return err
}
