#include <pthread.h>
#include <stdio.h>

int i = 0; //Husk i og ikke j
pthread_mutex_t k;

void minus(){
	for (int j = 0; j< 1000000;j++){
        pthread_mutex_lock(&k);
		i--;
        pthread_mutex_unlock(&k);
		//printf("minus gives: %i\n", i);
	}
}

void plus(){
	for (int j = 0; j< 1000000;j++){
		pthread_mutex_lock(&k);
        i++;
        pthread_mutex_unlock(&k);
		//printf("plus gives: %i\n", i);
	}
}



int main(){
    pthread_mutex_init(&k, NULL);
    pthread_t mThread;
    pthread_t pThread;

    pthread_create(&pThread, NULL, plus, NULL);
    pthread_create(&mThread, NULL, minus, NULL);
    
    pthread_join(pThread, NULL);
    pthread_join(mThread, NULL);
    printf("done gives: %i\n", i);
    return 0;
}
