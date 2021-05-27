import ctypes
from os.path import join, dirname, isfile
from os import walk, listdir
from time import time

class License:
    ROOT = dirname(__file__)
    # Shared Library
    so = ctypes.cdll.LoadLibrary(join(ROOT, 'compiled/libmatch.so'))

    def __init__(self):
        self.match = License.so.FindMatch
        self.match.argtypes=[ctypes.c_char_p]
        self.match.restype = ctypes.c_char_p

        self.custom = License.so.LoadCustomLicenses
        self.custom.argtypes=[ctypes.c_char_p]

        self.thresh = License.so.SetThreshold
        self.thresh.argtypes=[ctypes.c_float]

    def findMatch(self, root, searchSubDir = True):
        exec_time = time()
        filepath = list()
        if searchSubDir:
            for (dirpath, _, filenames) in walk(root):
                filepath += [dirpath + f for f in filenames]
        else:
            filepath = [root + f for f in listdir(root) if isfile(join(root, f))]
            
        filepath = '\n'.join(filepath)
        res = self.match(filepath.encode('utf-8'))
        exec_time = time() - exec_time
        return res.decode('utf-8'), exec_time
    
    def loadCustom(self, libpath):
        self.custom(libpath.encode('utf-8'))

    def setThreshold(self, thresh):
        print(self.thresh(thresh))

l = License()
res, exec_tm = l.findMatch('/home/avishrant/GitRepo/avishrantssh.github.io/')
print(res)
print(exec_tm)