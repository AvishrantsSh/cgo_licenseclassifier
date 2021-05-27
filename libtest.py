import ctypes
from time import time
from os import listdir
from os.path import isfile, join

def match(root):
    path = [root + f for f in listdir(root) if isfile(join(root, f))]
    # path = ["/home/avishrant/GitRepo/scancode.io/LICENSE"]

    # Shared Library Initialization
    so = ctypes.cdll.LoadLibrary('./compiled/libmatch.so')
    match = so.FindMatch
    custom = so.LoadCustomLicenses

    # Argument Data Type Initialization
    match.argtypes=[ctypes.c_char_p]
    match.restype = ctypes.c_char_p
    custom.argtypes=[ctypes.c_char_p]

    # Just for metrics :P
    start = time()

    path = '\n'.join(path)
    # print(path)
    custom('./classifier/alternate'.encode('utf-8'))
    res = match(path.encode('utf-8'))

    print(res.decode('utf-8'))
    print("Execution time : ", time() - start)

match('/home/avishrant/GitRepo/scancode.io/')