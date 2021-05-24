import ctypes
import sys

try:
    path = sys.argv[1]
except IndexError:
    print("No path specified")
    exit(-1)

so = ctypes.cdll.LoadLibrary('./compiled/shared/libprime.so')
match = so.FindMatch

match.argtypes=[ctypes.c_char_p]
match.restype=ctypes.c_char_p

res = match(path.encode('utf-8'))
print(res.decode('utf-8'))