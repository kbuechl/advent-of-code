BASE_PATH := ./2024
SOURCE_FOLDER := $(BASE_PATH)/0

.PHONY: day

# This target works for any day passed directly in the command
day:
	@DAY=$(word 2, $(MAKECMDGOALS)); \
	if [ -z $$DAY ]; then \
		echo "Please specify a day number (e.g., make day 4)"; \
		exit 1; \
	fi; \
	if [ $$DAY -lt 1 ] || [ $$DAY -gt 25 ] 2>/dev/null; then \
		echo "Invalid day number. Please enter a number between 1 and 25."; \
		exit 1; \
	fi; \
	DAY_DIR=$(BASE_PATH)/`printf "%02d" $$DAY`; \
	if [ -d $$DAY_DIR ]; then \
		echo "Day folder already exists: $$DAY_DIR"; \
		exit 1; \
	fi; \
	cp -r $(SOURCE_FOLDER) $$DAY_DIR; \
	echo "Copied $(SOURCE_FOLDER) to $$DAY_DIR."

# This is needed to prevent make from failing with 'No rule to make target'
%:
	@: