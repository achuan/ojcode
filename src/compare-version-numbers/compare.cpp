#include <stdio.h>
#include <stdlib.h>
#include <string>
#include <iostream>

using namespace std;

class Solution {
    public:

        int compareUnit(const string& string1, const string& string2)
        {
            int num1 =0 , num2 = 0;
            if(string1 == string2)
                return 0;
            sscanf(string1.c_str(), "%d", &num1);
            sscanf(string2.c_str(), "%d", &num2);

            if(num1 == num2 )
                return 0;

            return num1 > num2 ? 1 : -1;
        }

        int compareVersion(string version1, string version2) 
        {
            int nret = 0, pos1 = 0, pos2 = 0, posold1 = -1, posold2 = -1;

            if(version1 == version2)
                return 0;

            pos1 = version1.find('.');
            pos2 = version2.find('.');

            while(pos1 != string::npos && pos2 != string::npos)
            {
                nret = compareUnit(version1.substr(posold1+1, pos1-posold1-1), version2.substr(posold2+1, pos2 - posold2 - 1));
                if(nret != 0)
                    return nret;

                posold1 = pos1;
                posold2 = pos2;
                pos1 = version1.find('.', pos1+1);
                pos2 = version2.find('.', pos2+1);
            }

            if(pos1 == string::npos && pos2 == string::npos)
            {
                return compareUnit(version1.substr(posold1+1), version2.substr(posold2+1));
            }
            else if(pos1 == string::npos && pos2 != string::npos)
            {
                nret = compareUnit(version1.substr(posold1+1), version2.substr(posold2+1, pos2-posold2-1));
                if(nret != 0)
                    return nret;
                else
                    return compareUnit("0", version2.substr(pos2+1));
            }
            else
            {
                nret =  compareUnit(version1.substr(posold1+1), version2.substr(posold2+1, pos2-posold2-1));
                if(nret != 0)
                    return nret;
                else
                    return compareUnit(version1.substr(pos1+1), "0");
            }


            return nret;
        }
};

int main(int argc, char* argv[])
{
    if(argc > 2)
    {

        printf("text %s, %s\n", argv[1], argv[2]);
        string version1 = (char*)argv[1];
        string version2 = (char*)argv[2];

        Solution slt;
        printf("%s compare %s : result is:%d\n", version1.c_str(), version2.c_str(),slt.compareVersion(version1, version2));

    }

    return 0;
}
