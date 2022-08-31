package utils

var (
	// validation size uploud file
	MAX_SIZE_FILE = 2 * 1024 * 1024 // 2MB

	// Extension file
	JPG  = ".jpg"
	JPEG = ".jpeg"
	PNG  = ".png"

	// validation extension uploud file
	EXTENSION_FILE = map[string]bool{
		JPG:  true,
		JPEG: true,
		PNG:  true,
	}
)
