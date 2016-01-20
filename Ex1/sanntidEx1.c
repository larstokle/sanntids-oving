
/*
'global shared int i = 0

main:
    spawn thread_1
    spawn thread_2
    join all threads (wait for them to finish)
    print i

thread_1:
    do 1_000_000 times:
        i++
thread_2:
    do 1_000_000 times:
        i--
*/

#include <pthread.h>
#include <stdio.h>

int i = 0;

void minus(){
	for (int j = 0; j< 1000000;j++){
		i--;
		//printf("minus gives: %i\n", i);
	}
}

void plus(){
	for (int j = 0; j< 1000000;j++){
		i++;
		//printf("plus gives: %i\n", i);
	}
}



int main(){
    pthread_t mThread;
    pthread_t pThread;

    pthread_create(&pThread, NULL, plus, NULL);
    pthread_create(&mThread, NULL, minus, NULL);
    
    pthread_join(pThread, NULL);
    pthread_join(mThread, NULL);
    printf("done gives: %i\n", i);
    return 0;
}