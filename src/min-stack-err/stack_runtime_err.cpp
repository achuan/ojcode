#include<stdio.h>
#include<stdlib.h>
#include<string.h>

#include <stdio.h>
#include <stdlib.h>
#include <string>
#include <iostream>

using namespace std;

class MinStack {
public:
    MinStack()
    {
        m_length = 0;
        m_min = 0;
        m_capacity = 50;

        int nLen = sizeof(long long);
        m_pstack = (long long*)malloc(nLen * m_capacity);
    }

    ~MinStack()
    {
        if(m_pstack != NULL)
        {
            free(m_pstack);
            m_pstack = NULL;
        }
    }

    void push(int x) 
    {
        long long element = 0, tmp = 0;

        if(m_length >= m_capacity && m_capacity > 0)
        {
            int nLen = sizeof(long long) * (m_capacity * 2) ;
            long long * more_capacity = (long long*)realloc(m_pstack, nLen);
            if(more_capacity != NULL)
            {
                m_capacity *= 2;
                m_pstack = more_capacity; 
            }
            else
            {
                free(m_pstack);
                m_pstack = 0;
                m_capacity = 0;
                m_length = 0;
                return ;
            }
        }

        if(m_length == 0 || x < m_min)
            m_min = x;

        tmp = m_min;
        element = (tmp << 32) + x;
        m_pstack[m_length++] = element;
    }

    void pop() 
    {  
        if(m_length > 0)
        {
            m_pstack[--m_length] = 0;
            m_min = getMin();
        }
    }

    int top() 
    {
        return m_pstack[m_length-1] & 0xffffffff;
    }

    int getMin() 
    {
        long long tmp = m_pstack[m_length-1];
        return (tmp >> 32) & 0xffffffff;
    }

    void showElement()
    {
        printf("ShowElement...\n");
        for(int i = 0; i < m_length; ++i)
        {
            long long tmp = m_pstack[i];
            int small = (tmp >> 32 ) & 0xffffffff;
            int value = tmp & 0xffffffff;

            printf("min=%d, value=%d \n", small, value);
        }
        printf("\n");
    }

    int m_min;
    int m_capacity;
    int m_length;
    long long *m_pstack;
};



int main()
{
    int n = 0, min = 0, start = 0;

    while(scanf("%d %d %d", &min, &start, &n) == 3)
    {
        MinStack slt;
        for(int i = start; i < n; ++i)
        {
            slt.push(i);
        }

        slt.showElement();

        for(int i = start; i < n; ++i)
        {
            int top = slt.top();
            slt.pop();
            printf("pop element : %d \n", top);
        }

    }

    return 0;
}

