build:
	@echo "Building..."
	go build -C ./src -o ../dist/parallel-array-summary -v
	@echo "Build complete."

demonstrate:
	@echo "Running with n = 5, t = 1"
	./dist/parallel-array-summary 5 1
	@echo "Running with n = 5, t = 4"
	./dist/parallel-array-summary 5 4
	@echo "Running with n = 5, t = 16"
	./dist/parallel-array-summary 5 16
	@echo "Running with n = 5, t = 64"
	./dist/parallel-array-summary 5 64
	@echo "Running with n = 5, t = 256"
	./dist/parallel-array-summary 5 256
	@echo "Running with n = 7, t = 1"
	./dist/parallel-array-summary 7 1
	@echo "Running with n = 7, t = 4"
	./dist/parallel-array-summary 7 4
	@echo "Running with n = 7, t = 16"
	./dist/parallel-array-summary 7 16
	@echo "Running with n = 7, t = 64"
	./dist/parallel-array-summary 7 64
	@echo "Running with n = 7, t = 256"
	./dist/parallel-array-summary 7 256
