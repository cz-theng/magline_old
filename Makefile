
.DEFAULT:all

all : 
	cd maglined/app; make	
	cd magknot; make


.PHONY:clean

clean:
	cd maglined/app; make clean 
	cd magknot; make clean
