package main

func chunk[T any](s []T, chunkSize int) [][]T {
	var chunks [][]T

	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize

		if end > len(s) {
			end = len(s)
		}

		chunks = append(chunks, s[i:end])
	}

	return chunks
}
