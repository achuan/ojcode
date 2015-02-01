#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <vector>
#include <map>
#include <string>

using namespace std;

class Solution {
public:
    int majorityElement(vector<int> &num) 
	{
		int nret = 0, size = num.size()/2;
		map<int, int> map_record;
		vector<int>::iterator iter = num.begin();
		for(; iter != num.end(); ++iter)
		{
			if(map_record.find(*iter) != map_record.end())
				map_record[*iter] += 1;
			else
				map_record[*iter] = 1;
		}

		map<int, int>::iterator it = map_record.begin();
		for(; it != map_record.end(); ++it)
		{
			if(it->second > size)
			{
				nret = it->first;
				break;
			}
		}
        
		return nret;
    }
};

int main()
{
	int n = 0,num = 0;
	vector<int> vecnum;

	while(scanf("%d", &n) == 1)
	{
		vecnum.clear();
		for(int i = 0; i < n; ++i)
		{
			scanf("%d", &num);
			vecnum.push_back(num);
		}
		Solution slt;
		int nret = slt.majorityElement(vecnum);
		printf("majorityElement is :%d\n", nret);
	}

	return 0;
}
