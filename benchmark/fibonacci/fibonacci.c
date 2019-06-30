#include <stdio.h>
#include <stdint.h>
#include <sys/time.h>

static uint64_t fib(uint64_t n) {
  if (n <= 1) return 1;
  return fib(n - 1) + fib(n - 2);
}

long long current_timestamp() {
    struct timeval te;
    gettimeofday(&te, NULL); // get current time
    long long milliseconds = te.tv_sec*1000LL + te.tv_usec/1000; // calculate milliseconds
    // printf("milliseconds: %lld\n", milliseconds);
    return milliseconds;
}

int main(void) {

  long start = current_timestamp();
  printf("%llu \n", fib(45));
  long end = current_timestamp();
  printf("took %dms\n", end-start);
  return 0;
}
