
import threading
from threading import Thread


i = 0
l = threading.Lock()


def minus():
	global i
	for j in range(1000000):
		l.acquire()
		i-=1
		l.release()
		#print "minus gives: " , i, "\n"
	


def plus():
	global i
	for j in range(1000000):
		l.acquire()
		i+=1
		l.release()
		#print "plus gives: " , i, "\n"
	
	

def main():
    mThread = Thread(target = minus, args = (),)
    pThread = Thread(target = plus, args = (),)

    pThread.start()
    mThread.start()

    pThread.join()
    mThread.join()

    print "done gives: " , i, "\n"


main()