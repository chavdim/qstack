[![Go Report Card](https://goreportcard.com/badge/github.com/chavdim/qstack)](https://goreportcard.com/report/github.com/chavdim/qstack)
# qstack
quickly find the first stackoverflow answer from a yahoo search 
#Uses
>qstack export bin mac
```
Searching for: export+bin+mac
Source: http://stackoverflow.com/questions/11025980/how-to-add-usr-local-bin-in-path-on-mac
ANSWER##################################################
########################################################
CODE---------------------------------------
export PATH=$PATH:/usr/local/git/bin:/usr/local/bin
ENDCODE------------------------------------
One note: you don't need quotation marks here because it's on the right hand side of an assignment, but in general, and especially on Macs with their tradition of spacy pathnames, expansions like 
CODE---------------------------------------
$PATH
ENDCODE------------------------------------
 should be double-quoted as 
CODE---------------------------------------
"$PATH"
ENDCODE------------------------------------
.
QUESTION################################################
########################################################
How to add /usr/local/bin in $PATH on Mac
```
