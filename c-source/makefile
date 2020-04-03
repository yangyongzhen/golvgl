########################################
#makefile
########################################
#****************************************************************************
# Cross complie path
#****************************************************************************
#CHAIN_ROOT=/home/yang/imax283/ctools/gcc-4.4.4-glibc-2.11.1-multilib-1.0/arm-fsl-linux-gnueabi/bin
#CROSS_COMPILE=$(CHAIN_ROOT)/arm-none-linux-gnueabi-
CHAIN_ROOT= /home/yang/b503/ctools/gcc-linaro-arm-linux-gnueabihf-4.9-2014.09_linux/bin
CROSS_COMPILE=$(CHAIN_ROOT)/arm-linux-gnueabihf-

CROSS_COMPILE = 

#CC     := $(CROSS_COMPILE)gcc
#CXX    := $(CROSS_COMPILE)g++
#AS	   := $(CROSS_COMPILE)as
#AR     := $(CROSS_COMPILE)ar 
#LD     := $(CROSS_COMPILE)ld
#RANLIB := $(CROSS_COMPILE)ranlib
#OBJDUMP:= $(CROSS_COMPILE)objdump
#OBJCOPY:= $(CROSS_COMPILE)objcopy
#STRIP  := $(CROSS_COMPILE)strip
#编译主程序
BINARY  := liblvgl
OBJ_DIR := ./

INCS := -I ./   -I./lvgl/
CFLAGS= -Wall -g  -std=c99 

#****************************************************************************
# Source files
#****************************************************************************
SRC_C=$(shell find . -name "*.c")

OBJ_C=$(patsubst %.c, %.o, $(SRC_C))

SRCS := $(SRC_C) $(SRC_C)

OBJS := $(OBJ_C) 
LDSCRIPT=   
LDFLAGS= 
#LDSCRIPT= -lNC_FileSys
#LDFLAGS= -Llib 

#SRC  = $(wildcard *.c)
#DIR  = $(notdir $(SRC))
#OBJS = $(patsubst %.c,$(OBJ_DIR)%.o,$(DIR))
#OBJS=  main.o myutils.o  inirw.o  cmdpboc.o cputest.o bustcp.o ansrec.o m1cmd.o m1api.o m1test.o upcash.o myother.o getsys.o
#CFLAGS=-std=c99
#@echo Building lib...
#$(call make_subdir)
.PHONY: clean 
all:  prebuild  $(BINARY).so

prebuild:
	@echo Building shared lib...

$(BINARY).so : $(OBJS)
	@echo Generating ...
	$(CC)  -shared -fPIC -o  $(BINARY).so $(OBJS) $(LDFLAGS) $(LDSCRIPT) 
	@echo OK!

$(OBJ_DIR)%.o : %.c
	$(CC) -c -fPIC $(CFLAGS) $(INCS) $< -o  $@
	
clean:
	rm -f $(OBJ_DIR)*.o
	find . -name "*.[od]" |xargs rm
	@echo Removed!
