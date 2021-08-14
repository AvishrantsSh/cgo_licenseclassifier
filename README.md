Google License Classifier cgo binding
=====================================
**Note:** This library is a derivative work of [Google LicenseClassifier](https://github.com/google/licenseclassifier).

This is an extension of Google LicenseClassifier v2, which exposes the underlying functionality of the License Classifier in the form of shared library, which can further be integrated with other applications and tools to provide license classification and copyright detection.

Installation
------------
_**Note:** Currently, this shared library has been built only for Linux Platform. Work is in progress for other platforms._

Go ahead and setup your environment as
```sh
git clone https://github.com/AvishrantsSh/Ctypes_LicenseClassifier.git
go get .
```

Usage
-----
The [main.go](main.go) file contains several functions to detect license expressions in various files depending upon the requirements.

1. `ScanFile`

    This function scans the given file and returns the license expression and copyright notifications found in the file.

2. `BuffScanFile`

    This function is similar to `ScanFile` but it scans the file in a buffer. Suitable for high memory applications.

3. `CopyrightInfo`

    CopyrightInfo scans a given text/string and finds valid copyright notifications. Migrated from Google License Classifier v1.

Building Shared Library
-----------------------
To build your own shared library, you can use the following `make` command
```sh
make build
```

Or if you use an OS other than linux, you can use the following `build` command
```sh
go build -o compiled/libclassifier.so -buildmode=c-shared
``` 

Interfacing Example
-------------------
Now that you have the shared library, you can use it to develop your own wrapper to expose the underlying functionality. An example of how to use the shared library in Python is as follows:
```python
import ctypes
import os

ROOT = os.path.dirname(__file__)

# Shared Library Initialization
so = ctypes.cdll.LoadLibrary(os.path.join(ROOT, "compiled/libmatch.so"))
init = so.CreateClassifier
init.argtypes = [ctypes.c_char_p, ctypes.c_double]

scanfile = so.ScanFile
scanfile.argtypes = [ctypes.c_char_p, ctypes.c_int, ctypes.c_bool]
scanfile.restype = ctypes.c_char_p
...
scanfile(os.fsencode(location), max_size=100, use_buffer=False)
```