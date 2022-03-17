from time import time

def timesince(ts):
    if not isinstance(ts, float) and not isinstance(ts, int):
        return 0
    else:
        return (time.time() * 1000) - ts