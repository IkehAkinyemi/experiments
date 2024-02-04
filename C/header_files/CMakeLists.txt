cmake_minimum_required(VERSION 3.10)

# Project name
project(Vin)

# Specify the C standard
set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED True)

# Include directories
include_directories(include)

# Using globbing to include all .c files in lib
file(GLOB LIB_SOURCES "lib/*.c")
add_library(vinlib ${LIB_SOURCES})

# Add the executable
add_executable(vin main.c)

# Link the library with the executable
target_link_libraries(vin vinlib)