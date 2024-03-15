const MAX_BUFFER_SIZE = 1024 // Maximum size for a buffer
const PI = 3.14159           // Mathematical constant pi

func readData(buffer []byte) (int, error) {
    return os.Read(os.Stdin, buffer[:MAX_BUFFER_SIZE]) // Use constant
}
