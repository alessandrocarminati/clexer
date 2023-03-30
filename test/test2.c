#include <stdio.h>

#define square(x) (x * x)

int add(int x, int y) {
	return x + y;
}


int main() {
	int a = 10;
	int b = 20;
	int sum = add(a, b);
	int square_sum = square(sum);

	printf("Sum: %d\n", sum);
	printf("Square of sum: %d\n", square_sum);
	// ed ora una stampa \importante\
	printf("stocazzo\n");
	return 0;
asd
}
