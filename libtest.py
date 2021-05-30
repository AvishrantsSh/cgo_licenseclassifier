import ctypes
from error import *
from os.path import join, dirname, isfile, isdir
from os import walk, listdir
from time import time


class License:
    _ROOT = dirname(__file__)

    # Shared Library
    _so = ctypes.cdll.LoadLibrary(join(_ROOT, "compiled/libmatch.so"))
    _match = _so.FindMatch
    _match.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
    _match.restype = ctypes.c_char_p

    _loadCustomLib = _so.LoadCustomLicenses
    _loadCustomLib.argtypes = [ctypes.c_char_p]
    _loadCustomLib.restype = ctypes.c_int

    _setThresh = _so.SetThreshold
    _setThresh.argtypes = [ctypes.c_int]
    _setThresh.restype = ctypes.c_int

    def __init__(self):
        pass

    def findMatch(self, filepath):
        """Function to find a license match for file specified by `filepath`"""
        if isdir(filepath):
            raise PathIsDir

        exec_time = time()
        res = self._match(License._ROOT.encode("utf-8"), filepath.encode("utf-8"))
        exec_time = time() - exec_time
        return res.decode("utf-8"), exec_time

    def catalogueDir(self, root, searchSubDir=True):
        """Function to find a license match for all files present in `root`"""
        exec_time = time()
        filepath = list()
        if searchSubDir:
            for (dirpath, _, filenames) in walk(root):
                filepath += [dirpath + f for f in filenames]
        else:
            filepath = [root + f for f in listdir(root) if isfile(join(root, f))]

        filepath = "\n".join(filepath)
        res = self._match(License._ROOT.encode("utf-8"), filepath.encode("utf-8"))
        exec_time = time() - exec_time
        return res.decode("utf-8"), exec_time

    def loadCustom(self, libpath):
        """Load a custom set of licenses"""
        _ = self._loadCustomLib(libpath.encode("utf-8"))

    def setThreshold(self, thresh):
        """Set a threshold between `0 - 100`. Default is `80`. Speed Degrades with lower threshold"""
        _ = self._setThresh(thresh)


l = License()
l.setThreshold(75)
res, exec_tm = l.catalogueDir(
    "/home/avishrant/GitRepo/scancode.io/", searchSubDir=False
)
print(res)
print(exec_tm)
