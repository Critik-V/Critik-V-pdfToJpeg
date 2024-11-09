package utils

import "os"

const defaultPdfDir = "pdf" // Default directory where the PDFs are stored
const defaultImgDir = "img" // Default directory where the images will be stored

// IsProduction returns true if the application is running in production mode
func IsProduction() bool {
	return os.Getenv("GIN_MODE") == "production"
}

// GetPort returns the port where the server will be running
func GetPort() string {
	return ":5001"
}

// GetPdfDir returns the path to the directory where the PDFs are stored
func GetPdfDir() string {
	name := os.Getenv("PDF_STORAGE_PATH")
	if name == "" {
		return defaultPdfDir
	}
	return name
}

// GetImgDir returns the path to the directory where the images will be stored
func GetImgDir() string {
	name := os.Getenv("IMAGE_STORAGE_PATH")
	if name == "" {
		return defaultImgDir
	}
	return name
}

// GetCorsOrigin returns the origin to be allowed by the CORS middleware
func GetCorsOrigin() string {
	origin := os.Getenv("SERVER_SERVICE_ORIGIN")
	if origin == "" {
		return "*"
	}
	return origin
}
