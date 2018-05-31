from django.http import HttpResponse  
  
# http response content  
def output(request):  
    title = "<h1>When You Are Old</h1>"  
    author = "<h2>William Butler Yeats</h2>"  
    content = """ 
                 When you are old and grey and full of sleep,<br/> 
                 And nodding by the fire, take down this book,<br/> 
                 And slowly read, and dream of the soft look<br/> 
                 Your eyes had once, and of their shadows deep;<br/> 
                 How many loved your moments of glad grace,<br/> 
                 And loved your beauty with love false or true,<br/> 
                 But one man loved the pilgrim soul in you,<br/> 
                 And loved the sorrows of your changing face;<br/> 
                 And bending down beside the glowing bars,<br/> 
                 Murmur, a little sadly, how love fled<br/> 
                 And paced upon the mountains overhead<br/> 
                 And hid his face amid a crowd of stars.<br/> 
               """  
    return HttpResponse([title, author, content])  