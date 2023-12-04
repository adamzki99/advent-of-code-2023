#!/bin/bash

# Check if the user provided an argument
if [ $# -eq 0 ]; then
  echo "Usage: $0 <day_number>"
  exit 1
fi

# Get the day number from the user input
day_number=$1

# Create the directory
dir_name="day_${day_number}"
mkdir "$dir_name"

# Create the Go source file
go_source_file="${dir_name}/day${day_number}.go"
touch "$go_source_file"

# Create the Go test file
go_test_file="${dir_name}/day${day_number}_test.go"
touch "$go_test_file"

# Create the puzzle input file
puzzle_input_file="${dir_name}/puzzle_input.txt"
touch "$puzzle_input_file"

# Display success message
echo "Directory and files created successfully:"
echo " - Directory: $dir_name"
echo " - Go source file: $go_source_file"
echo " - Go test file: $go_test_file"
echo " - Puzzle input file: $puzzle_input_file"
