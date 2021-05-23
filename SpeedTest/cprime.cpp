#include <iostream>
#include <cmath>
int NUM = 1000000;

using namespace std;
bool isPrime(long num)
{
    if (num < 2)
        return false;

    for (int i = 2; i < int(sqrt(num)) + 1; i++)
    {
        if (num % i == 0)
            return false;
    }
    return true;
}

int totalPrime()
{
    int count = 0;
    for (int i = 2; i <= NUM; i++)
    {
        if (isPrime(i))
            count++;
    }
    return count;
}

int main()
{
    cout << totalPrime() << endl;
    return 0;
}