// Created by: Timothée MORANDEAU
// with the help of GitHub Copilot
// and Romain Houard, Pierre Frank-Papuchon

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <pthread.h>
#include <string.h>

#define NB_TOURS 10
#define NB_THREAD 100
int production = NB_TOURS*NB_THREAD;
int value = 0;
int *count = &value;
//on va ajouter des mutex et des cond
//pour gerer la concurrence
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t cond_empty = PTHREAD_COND_INITIALIZER;
pthread_cond_t cond_full = PTHREAD_COND_INITIALIZER;

typedef struct Paquet{
    char message[20];
} Paquet;

typedef struct FIFO {
    Paquet *buffer;
    int size;
    int head;
    int tail;
} FIFO;

void fifo_init(FIFO *fifo, int size){
    fifo->buffer = malloc(size * sizeof(char *));
    fifo->size = size;
    fifo->head = 0;
    fifo->tail = 0;
}

void fifo_pop(FIFO *fifo, Paquet *value){
    //usage des mutex
    pthread_mutex_lock(&mutex);
    while(*count == 0){
        pthread_cond_wait(&cond_empty, &mutex);
    }
    *value = fifo->buffer[fifo->head];
    fifo->head = (fifo->head + 1) % fifo->size;
    if (rand() % 2 == 0) {
        usleep(rand()%100);
    }
    (*count)--;
    pthread_cond_signal(&cond_full);
    pthread_mutex_unlock(&mutex);
}

void fifo_add(FIFO *fifo, Paquet value){
    //usage des mutex
    pthread_mutex_lock(&mutex);
    while(*count == fifo->size){
        pthread_cond_wait(&cond_full, &mutex);
    }
    fifo->buffer[fifo->tail] = value;
    fifo->tail = (fifo->tail + 1) % fifo->size;
    if (rand() % 2 == 0) {
        usleep(rand()%100); // Sleep for 100 microseconds
    }
    (*count)++;
    pthread_cond_signal(&cond_empty);
    pthread_mutex_unlock(&mutex);
}

void *thread_producteur(FIFO *fifo){
    Paquet p;
    for(int i = 0; i < NB_TOURS; i++){
        printf("Message envoyé : %d\n", i);
        sprintf(p.message, "Message %d", i);
        fifo_add(fifo, p);
    }
    pthread_exit(NULL);
}

void *thread_consommateur(FIFO *fifo){
    Paquet p;
    while(production > 0){
        fifo_pop(fifo, &p);
        printf("Message lu : %s\n", p.message);
        production--;
        printf("Production : %d\n", production);
        printf("Count : %d\n", *count);
    }
    pthread_exit(NULL);
}

int main(void)
{
    FIFO fifo;
    fifo_init(&fifo, 10);
    fifo.size = 10;
    //creer une dizaine de producteurs et une dizaine de consommateurs
    pthread_t producteur[NB_THREAD];
    pthread_t consommateur;

    pthread_create(&consommateur, NULL, (void *)thread_consommateur, &fifo);
    for(int i = 0; i < NB_THREAD; i++){
        pthread_create(&producteur[i], NULL, (void *)thread_producteur, &fifo);
    }


}
