#include<stdio.h>
#include<stdlib.h>
#include<iostream>
#include<vector>



using namespace std;


class Solution {
public:
    int majorityElement(vector<int> &num) 
	{
		int size = num.size();
		int count = 0;
		int candicate = 0;
		vector<int>::iterator iter = num.begin();

		for(; iter != num.end(); ++iter)
		{
			if(count == 0)
			{
				candicate = *iter;
				count = 1;
			}
			else
			{
				candicate == *iter ? ++count : --count;
				if(count > size/2)
					break;
			}
		}
		return candicate;
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
