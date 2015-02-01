#include <stdio.h>
#include <stdlib.h>
#include <iostream>
#include <map>

using namespace std;


class Solution {
public:
    string fractionToDecimal(int numerator, int denominator) 
	{
		string strret;
		map<int, int> nummap;
		map<int, int>::iterator iter;
	
		int nindex = 0;

		if(numerator == 0)
			return "0";

		if((numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0))
		{
			strret += "-";
		}
	
		long long n = numerator;
		long long d = denominator;

		if(n < 0)
			n = -n;
		if(d < 0)
			d = -d;

		char buffer[32] = {0};
		sprintf(buffer, "%lld", n/d);
		strret += buffer;

		n %= d;

		if(n > 0)
			strret += '.';

		nindex = strret.size();

		while(n % d != 0)
		{
			iter = nummap.find(n);
			if(iter == nummap.end())
			{
				nummap[n] = nindex++;
				n *= 10;
				strret += n/d + '0';
				n %= d;
			}
			else
			{
				strret.insert(iter->second, "(");
				strret += ")";
				break;
			}
		}
       
		return strret;
    }
};

int main()
{
	int num, denom;
	while(scanf("%d, %d", &num, &denom) == 2)
	{
		Solution slt;
		string strret = slt.fractionToDecimal(num, denom);
		printf("%d div %d : result = %s\n", num, denom,strret.c_str());
	}

	return 0;
}
