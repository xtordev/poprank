package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"strings"
	"time"

	storage_go "github.com/supabase-community/storage-go"
	"github.com/supabase-community/supabase-go"
)

var SupabaseClient *supabase.Client

func SupabaseInit(API_KEY string, API_URL string) (*supabase.Client, error) {
	options := &supabase.ClientOptions{} // Initialize with default options or customize as needed
	client, err := supabase.NewClient(API_URL, API_KEY, options)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize client: %w", err)
	}
	return client, nil
}
func generateRandomFilename(fileExt string) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 10

	randomString := make([]byte, length)
	for i := range randomString {
		randomIndex := rand.Intn(len(letters)) // Generate random index
		randomString[i] = letters[randomIndex]
	}

	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s.%s", timestamp, string(randomString), fileExt)
}
func UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a buffer to read the file into
	buf := make([]byte, fileHeader.Size)
	_, err = file.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	fileExt := strings.Split(fileHeader.Filename, ".")[len(strings.Split(fileHeader.Filename, "."))-1]
	// Upload the file to Supabase Storage
	filePath := generateRandomFilename(fileExt)
	uploadedFile, err := SupabaseClient.Storage.UploadFile("files", filePath, bytes.NewReader(buf), storage_go.FileOptions{})
	if err != nil {
		return "", err
	}
	publicUrl := SupabaseClient.Storage.GetPublicUrl("files", strings.Split(uploadedFile.Key, "/")[1])
	return publicUrl.SignedURL, nil
}

func DeleteImage(fileLink string) {
	imageTarget := strings.Split(fileLink, "/")[len(strings.Split(fileLink, "/"))-1]
	SupabaseClient.Storage.RemoveFile("files", []string{imageTarget})
}
