import datetime

# utils
def read_file(filePath):
    contentList = None
    f = open(filePath)
    if f:
        contentList = []
        line = f.readline()
        while line:
            contentList.append(line)
            line = f.readline()
        f.close()
    return contentList


def format_time(timeStr):
    bootTime = datetime.datetime.strptime(timeStr, '%Y-%m-%d %H:%M:%S')
    return bootTime 

