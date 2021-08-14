Google License Classifier cgo binding
=====================================
**Note:** This library is a derivative work of [Google LicenseClassifier](https://github.com/google/licenseclassifier).

This is an extension of Google LicenseClassifier v2, which exposes the underlying functionality of the License Classifier in the form of shared library that can be integrated with other tools to provide license classification and copyright detection.

Installation
------------
_Note: Currently, this package only supports Linux Platform. Work is in progress for Windows and Mac._

Go ahead and setup your environment as
```sh
git clone https://github.com/AvishrantsSh/Ctypes_LicenseClassifier.git
go get .
```

Usage
-----
To build your own shared library, you can use the following `make` command
```sh
make build
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