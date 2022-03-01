#include <iostream>
#include <cstdio>
#include <fstream>
#include <string>
#include <set>
using namespace std;
const int mod = 10000000;
#define inf 0x3f3f3f3f
const int maxn = 505;
int n,m;
int T;
char cj[3] = {'R','G','B'};
int main() {
	srand(time(0));
	n = rand() % 100;
	cout << n << endl;
	for(int i = 1; i <= n; i++) cout << cj[rand() % 3];
	return 0;
}