# C语言入门

##### Linux:

```bash
# 1.install dev env
yum install -y gcc
gcc -v

# 2.create hello.c
mkdir cdemo 
cd cdemo
vi hello.c
	#include<studio.h>
	int main()
	{
   	printf("Hello World\n");
   	return 0;
	}



# 3.gcc hello.c
gcc hello.c
ls -l
	hello.c
	a.out (defult)
	
or

gcc hello.c -o hello.out 
ls -l
	hello.c
	hello.out

# 4.execute a.out
./a.out  (or ./hello.out)
	Hello World

```



