build:
	@echo "Building..."
	go build -C ./src -o ../dist/parallel-array-summary -v
	@echo "Build complete."

demonstrate:
	@echo "Running with n = 5, t = 16, e = 0"
	./dist/parallel-array-summary 5 16 0
	@echo "Running with n = 5, t = 16, e = 40"
	./dist/parallel-array-summary 5 16 40
	@echo "Running with n = 5, t = 256, e = 40"
	./dist/parallel-array-summary 5 256 40
	@echo "Running with n = 7, t = 256, e = 0"
	./dist/parallel-array-summary 7 256 0
	@echo "Running with n = 7, t = 256, e = 40"
	./dist/parallel-array-summary 7 256 40
	@echo "Running with n = 9, t = 256, e = 0"
	./dist/parallel-array-summary 9 256 0
	@echo "Running with n = 9, t = 256, e = 40"
	./dist/parallel-array-summary 9 256 40
