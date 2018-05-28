#include<stdio.h>
int main(){
	for(long i=0; i<10000000000; ++i){
		if(i%500000000 == 0){
			printf("Hello world %ld\n", i);
		}
	}
}
