#include "stdio.h"
#include <sys/time.h>

int main() {
  int sum = 0;
  struct timeval  tv1, tv2;
  gettimeofday(&tv1, NULL);

  printf("michael: 0:\n");
  for (int l = 0; l < 50; l++)
	for (int k = 0; k < 100; k++)
	  for (int i = 0; i < 100; i++) {
		sum = 0;
		for (int j = 0; j < 100; j++) 
		  sum *= j;
	  }
  gettimeofday(&tv2, NULL);

  printf("michael: 1: %ld\n", tv2.tv_usec - tv1.tv_usec + (tv2.tv_sec - tv1.tv_sec)*1000000);
  return 0;
}
