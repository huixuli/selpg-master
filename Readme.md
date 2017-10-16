﻿# 使用 golang 开发 开发 Linux 命令行实用程序 中的 selpg

标签（空格分隔）： go语言学习

---

## 任务描述
使用go语言[开发Linux命令行实用程序][1]中的selpg命令行程序，该程序允许用户通过读取文件、标准输入或利用管道从另一个进程中输入一个文本，同时以参数指定要抽取输入文本的页的范围，该实用程序从标准输入或从作为命令行参数给出的文件名读取文本输入。它允许用户指定来自该输入并随后将被输出的页面范围。例如，如果输入含有100页，则用户可指定只打印第 35 至 65页。这种特性有实际价值，因为在打印机上打印选定的页面避免了浪费纸张。另一个示例是，原始文件很大而且以前已打印过，但某些页面由于打印机卡住或其它原因而没有被正确打印。在这样的情况下，则可用该工具来只打印需要打印的页面。

## 程序测试结果
**1  ./selpg -s=0 -e=1 text.txt**

![1](/图片/1.png)
![1.1](/图片/1.1.png)
 如图片所示，输出了text.txt文档的第0页，默认为72行。
 
 **2  ./selpg -s=0 -e=1 -l=5 text.txt**
 
 ![2](/图片/2.png)
如图片所示，规定的行数为5，所以打印了文档的第0页，共5行。

 **3  ./selpg -s=2 -e=1 -l=5 text.txt**
 
  ![3](/图片/3.png)
  如图片所示，如果开始的页数小于结束的页数，会提示参数错误。
  
  **4  ./selpg -s=0 -e=3 -l=5 -d=lp1 text.txt**
  
 ![4](/图片/4.png)
如图所示，页数正常打印，并且显示了它在lp1中的队列。

  **5  ./selpg -s=0 -e=3 <text.txt**
  
   ![5](/图片/5.png)
   ![5](/图片/5.1.png)
   正常打印
   
  **6  ./selpg -s=0 -e=3 text.txt >out.txt**
  
  ![6](/图片/6.png)
  ![6.1](/图片/6.1.png)
  如图所示，正常的将打印的结果输入到了out.txt文档中
  
  **7  ./selpg -s=0 -e=1 text.txt 2>error.txt**
  
   ![7](/图片/7.png)
   ![7.1](/图片/7.1.png)
   ![7.2](/图片/7.2.png)
   如图所示，当参数没有错误时，正常打印页数，而当参数出现错误，显示错误信息并将错误信息输入到error.txt文档中

  [1]: https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html
