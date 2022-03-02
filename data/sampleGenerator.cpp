#include <iostream>
#include <cstdio>
#include <fstream>
#include <string>
#include <set>
#include<random>
using namespace std;
const int mod = 10000000;
#define inf 0x3f3f3f3f
const int maxn = 505;
int n,m;
int T;
int main() {
	mt19937_64 r(chrono::system_clock::now().time_since_epoch().count());
	n = r() % 100; m = r() % 100;
	cout << n << ' ' << m;
	return 0;
}