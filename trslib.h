
#ifndef _TRSLIB_H
#define _TRSLIB_H

#include <stdio.h>

typedef struct myStruct1 {
	char a[10];
} st1_t;

typedef struct myStruct2 {
	char a[8];
	char b[2];
} st2_t;

extern st1_t s1;
extern st2_t s2;

extern void myPrintFunction2(void); 

//void myPrintFunction(char *s);

#endif /* TRSLIB_H */
