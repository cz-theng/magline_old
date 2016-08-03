
.DEFAULT:all

all : 
	cd maglined; make	
	cd magknot; make


.PHONY:clean

clean:
	cd maglined; make clean 
	cd magknot; make clean
