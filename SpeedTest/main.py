import ctypes
from time import time

so = ctypes.cdll.LoadLibrary('./_libprime.so')
total = so.totalPrime
total.argtypes = [ctypes.c_int64]
print("Starting Execution")
start = time()
print(total(1000000))
print("Execution Time : ", time() - start)