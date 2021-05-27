class PathIsDir(Exception):
    """Exception raised when given path does not correspond to file"""
    def __str__(self):
        return "The given path does not correspond to a file"