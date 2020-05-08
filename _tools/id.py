import datetime

def getID():
    x = datetime.datetime.now()
    return x.strftime("%Y%m%d%H%M%S")