# selpg

## detail
[开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)

## example


### 1
```
./selpg -h
```

```
Usage of ./selpg:

./selpg is a tool to select pages from what you want.

Usage:

        selpg -s=Number -e=Number [options] [filename]

The arguments are:

        -s=Number       Start from Page <number>.
        -e=Number       End to Page <number>.
        -l=Number       [options]Specify the number of line per page.Default is 72.
        -d=lp number    [options]Using cat to test.
        -f              [options]Specify that the pages are sperated by \f.
        [filename]      [options]Read input from the file.

If no file specified, ./selpg will read input from stdin. Control-D to end.

```

Page number start from 0.

### 2
```
./selpg -s=1 -e=3 -l=5 test.txt
```


```
Line-1
Line-1
Line-1
Line-1
Line-1
Line-1
Line-1
Line-1
Line-1
Line-1
```

### 3
using cat to test the function.

```
 ./selpg -s=0 -e=1 -l=2 -d=lp1 test.txt
```

```
     1  Line-1
     2  Line-1
```

### 4

```
./selpg -s=1 -e=3 -l=2 
```

Using terminal to input.


