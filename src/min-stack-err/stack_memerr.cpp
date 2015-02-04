#include<stdio.h>
#include<stdlib.h>
#include<iostream>

using namespace std;

struct node
{
    int value;
    int nmin;
    node *next;
};

#define min(a, b)  (((a)<(b))?(a):(b))

class MinStack {
    public:
        MinStack()
        {
            m_pnode = NULL;
        }

        void push(int x) 
        {
            node *pnew = (node*)malloc(sizeof(node));
            pnew->value = x;
            pnew->next = NULL;
            
            if(m_pnode == NULL)
            {
                pnew->nmin = x;
                m_pnode = pnew;
            }
            else
            {
                pnew->nmin = min(x, m_pnode->nmin);
                node * tmp = m_pnode;
                m_pnode = pnew;
                m_pnode->next = tmp;
                
            }
        }

        void pop()
        {
            node * tmp = m_pnode->next;
            free(m_pnode);
            m_pnode = tmp;
        }

        int top()
        {
            return m_pnode->value;
        }

        int getMin()
        {
            return m_pnode->nmin;
        }

        node* m_pnode;
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


        for(int i = start; i < n; ++i)
        {
            int top = slt.top();
            slt.pop();
            printf("pop element : %d \n", top);
        }

    }

    return 0;
}

