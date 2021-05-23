from time import time
import math
start = time()
NUM = 1000000

def isPrime(num):
    if num < 2:
        return False
    for i in range(2, int(math.sqrt(num)) + 1):
        if num % i == 0:
            return False
    return True

def totalPrime():
    count = 0
    for i in range(2, NUM+1):
        if isPrime(i):
            count += 1

    return count

print(totalPrime())
print("Execution time : ", time() - start)

