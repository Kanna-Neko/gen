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
char in[maxn];
int main() {
	cin >> n;
	for(int i = 1; i <= n; i++) cin >> in[i];
	for(int i = 1; i <= n; i++) cout << in[i];
	return 0;
}