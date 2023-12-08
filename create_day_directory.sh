#!/bin/bash

# Check if the user provided an argument
if [ $# -eq 0 ]; then
  echo "Usage: $0 <new_suffix>"
  exit 1
fi

# Get the new suffix from the user input
new_suffix=$1

# Source directory
source_dir="day_template"

# Destination directory
dest_dir="day_${new_suffix}"

# Check if the source directory exists
if [ ! -d "$source_dir" ]; then
  echo "Error: Source directory '$source_dir' not found."
  exit 1
fi

# Check if the destination directory already exists
if [ -d "$dest_dir" ]; then
  echo "Error: Destination directory '$dest_dir' already exists. Please choose a different suffix."
  exit 1
fi

# Create the destination directory
mkdir "$dest_dir"

# Copy files from source to destination and rename with the new suffix
for file in "$source_dir"/*; do
  if [ -f "$file" ]; then
    filename=$(basename "$file")
    new_filename="${filename//X/$new_suffix}"
    cp "$file" "$dest_dir/$new_filename"
  fi
done

echo "Directory and files copied successfully with the new suffix '$new_suffix':"
echo " - Source directory: $source_dir"
echo " - Destination directory: $dest_dir"
