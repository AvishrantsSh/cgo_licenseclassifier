import ctypes
from error import PathIsDir
from os.path import join, dirname, isfile, isdir
from os import walk, listdir
from time import time
from error import *

class License:
    ROOT = dirname(__file__)
    # Shared Library
    so = ctypes.cdll.LoadLibrary(join(ROOT, 'compiled/libmatch.so'))
    match = so.FindMatch
    match.argtypes=[ctypes.c_char_p]
    match.restype = ctypes.c_char_p

    loadCustomLib = so.LoadCustomLicenses
    loadCustomLib.argtypes=[ctypes.c_char_p]
    loadCustomLib.restype = ctypes.c_int

    setThresh = so.SetThreshold
    setThresh.argtypes=[ctypes.c_int]
    setThresh.restype = ctypes.c_int

    def __init__(self):
        pass

    def findMatch(self, filepath):
        """Function to find a license match for file specified by `filepath`"""
        if isdir(filepath):
            raise PathIsDir
        exec_time = time()
        res = self.match(filepath.encode('utf-8'))
        exec_time = time() - exec_time
        return res.decode('utf-8'), exec_time
    
    def catalogueDir(self, root, searchSubDir = True):
        """Function to find a license match for all files present in `root`"""
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
        """Load a custom set of licenses. Licenses should be named as `<license>.txt`"""
        _ = self.loadCustomLib(libpath.encode('utf-8'))

    def setThreshold(self, thresh):
        """ Set a threshold between `0 - 100`. Default is `80`. Speed Degrades with lower threshold. """
        _ = self.setThresh(thresh)

l = License()
l.setThreshold(75)
res, exec_tm = l.catalogueDir('/home/avishrant/GitRepo/scancode.io/', searchSubDir=False)
print(res)
print(exec_tm)