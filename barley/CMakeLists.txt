# CMake build script
# Chen - My Network protocol tests
cmake_minimum_required(VERSION 3.11)

project(barley)
set_property(GLOBAL PROPERTY USE_FOLDERS ON)
set(CMAKE_EXPORT_COMPILE_COMMAND ON)

if (MSVC)
  set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /W4")
  # Enable the static libraries on Windows (/MT)
  foreach (flag_var
    CMAKE_CXX_FLAGS CMAKE_CXX_FLAGS_DEBUG CMAKE_CXX_FLAGS_RELEASE
    CMAKE_CXX_FLAGS_MINSIZEREL CMAKE_CXX_FLAGS_RELWITHDEBINFO
    CMAKE_C_FLAGS CMAKE_C_FLAGS_DEBUG CMAKE_C_FLAGS_RELEASE
    CMAKE_C_FLAGS_MINSIZEREL CMAKE_C_FLAGS_RELWITHDEBINFO)
    # update to the static version of the run time library
    string(REPLACE "/MD" "/MT" ${flag_var} "${${flag_var}}")
  endforeach()
  set(CMAKE_C_STANDARD_LIBRARIES "")
  set(CMAKE_CXX_STANDARD_LIBRARIES "")
endif()

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# enable boost static flag
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

target_link_libraries(barley libbarley
  Boost::unit_test_framework)

target_link_libraries(barleytest libbarley
  Boost::unit_test_framework)

