# CMake build script
cmake_minimum_required(VERSION 4.1.0)

project(barley)
set_property(GLOBAL PROPERTY USE_FOLDERS ON)
set(CMAKE_EXPORT_COMPILE_COMMAND ON)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_MSVC_RUNTIME_LIBRARY "MultiThreaded$<$<CONFIG:Debug>:Debug>")

# enable boost static flag
unset(Boost_LIBRARIES)
set(Boost_USE_STATIC ON)
set(Boost_USE_STATIC_LIBS ON)
set(Boost_USE_MULTITHREADED ON)
set(Boost_USE_STATIC_RUNTIME ON)
find_package(Boost 1.88.0 REQUIRED COMPONENTS unit_test_framework)

include_directories("${Boost_INCLUDE_DIRS}"
  "${PROJECT_SOURCE_DIR}/src")

file(GLOB SOURCES_GROUP_MISC
  "src/barley_util.*")

file(GLOB SOURCES_GROUP_CFG
  ".editorconfig")

source_group("misc" FILES ${SOURCES_GROUP_MISC})
source_group("conf" FILES ${SOURCES_GROUP_CONF})

add_library(libbarley STATIC
  "${PROJECT_SOURCE_DIR}/src/barley_util.cxx"
  "${PROJECT_SOURCE_DIR}/src/barley_util.hxx")

add_executable(barley
  ".editorconfig"
  "${PROJECT_SOURCE_DIR}/src/main.cxx")

add_executable(barleytest
  "${PROJECT_SOURCE_DIR}/test/barley_test_util.cxx"
  "${PROJECT_SOURCE_DIR}/test/barley_test.cxx")

target_link_libraries(barley libbarley)
target_link_libraries(barleytest libbarley Boost::unit_test_framework)

