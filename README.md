Pasta-Cutter image cloner
===================

Idea
---------------

[![Original video](https://img.youtube.com/vi/f1fXCRtSUWU/0.jpg)](https://www.youtube.com/watch?v=f1fXCRtSUWU)


Given an image, it creates 4 cloned versions made with pieces from the original. 
Original image: 
![enter image description here](https://github.com/xfreakart/pasta_cutter/blob/master/gopher.jpg?raw=true)
we get:
![enter image description here](https://github.com/xfreakart/pasta_cutter/blob/master/output_images/output_final.png?raw=true)

Useful?¿ 
Not at all, but it was **fun**.


Params
-------------
```
go run main.go -debug -file=gopher.jpg
```

```shell
-debug creates all steps images, nice for debugging, don't add the flag for disable debugging images.
-file=FILENAME   route of image (JPG) to be processed.
```

Dependencies
---------
```
$ go get github.com/sirupsen/logrus
$ go get github.com/disintegration/imaging
```
