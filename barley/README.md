# byte-grind/barley

c++ test code and prototypes

Prerequest:
Install the vcpkg and set the `VCPKG_ROOT` env to your local vcpkg

Build steps:
1. create project files via cmake (add your boost library to cmake) 
   1.1. The cmake config is required CMake version 4.1.0+
   1.2. build the boost for static multithread runtime 
        (E.g. : b2 install --prefix=<install folder> toolset=msvc link=static address-model=64 threading=multi runtime-link=static --build-type=complete)
   1.3. add -DBoost_DIR to cmake 
       (e.g.: cmake <src dir>  -DBoost_DIR=<boost cmake config dir> )
2. build the project via the project files

