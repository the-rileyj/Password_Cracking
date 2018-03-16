
#!/usr/bin/python
import hashlib

hash = "5f4dcc3b5aa765d61d8327deb882cf99"
#wordlist = "/usr/share/john/password.lst"
wordlist = "lists/rockyou.txt"

num = 0

for line in open(wordlist, "r"):
    line = line.replace("\n", "")
    
    num = num + 1
    
    if hashlib.md5(line).hexdigest() == hash:
        print "WOO! Got one on line " + str(num) + ". The password is: " + line
        quit()
        
        
