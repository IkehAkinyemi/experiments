#!/bin/zsh

echo "Starting program at $(date)" # Date will be substituted

echo "Running program $0 with $# arguments with pid $$"

cd ../ || exit

for file in "$@"; do
  grep foobar "$file" > /dev/null 2> /dev/null
  # When a pattern is not found, grep has exit status 1
  # you redirect the STDOUT and STDERR to a null register since we don't care about them
  if [["$?" -n 0]]; then
  echo "File $file does not have any foobar, adding one"
  echo "# foobar" >> "$file"
  fi
done