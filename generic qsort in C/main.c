
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef struct grade
{
    char *name;
    float grade;
} grade;
int i;

grade students[5];

char *names[5] = {"John", "Eric", "Paul", "Carol", "Jim"};
float grades[5] = {90, 76, 80, 70, 60};
int compare1(const void *i1, const void *i2)
{
    const grade *r1 = i1;
    const grade *r2 = i2;
    return (*r2).grade-(*r1).grade;
}

int compare2(const void *s1, const void *s2) {
    const grade *r1 = s1;
    const grade *r2 = s2;
    return strcmp((*r1).name, (*r2).name);

}
int main()
{

    for (i = 0; i < 5; i++)
    {
        students[i].name = names[i];
        students[i].grade = grades[i];
    }

    qsort( students, 5, sizeof(students[0]), compare1);
    for (i = 0; i < 5; i++)
    {
        printf("(%s, %f) ", students[i].name, students[i].grade);
    }
    printf("\n");
    qsort( students, 5, sizeof(students[0]), compare2 ); //your code

    for (i = 0; i < 5; i++)
    {
        printf("(%s, %f)", students[i].name, students[i].grade);
    }
    printf("\n");
    return 0;
}

// main
//output:
// (1) (Jim, 60.000000) …  (John, 90.000000)
// (2)  (Carol, 70.00000) (Eric, 76.00000) … (Paul, 80.000000)