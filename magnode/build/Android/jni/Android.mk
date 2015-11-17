LOCAL_PATH := $(call my-dir)

######################################
# libapollo_voice
include $(CLEAR_VARS)
SRC_PATH:=$(LOCAL_PATH)/../../../
LOCAL_MODULE := magnode 

LOCAL_CFLAGS := -D__gnu_linux__

LOCAL_CFLAGS += $(L_CFLAGS)
LOCAL_CFLAGS +=  -ftree-vectorize -ffast-math

LOCAL_ARM_MODE := arm

LOCAL_LDLIBS += -llog

LOCAL_C_INCLUDES := $(SRC_PATH)/include
LOCAL_C_INCLUDES += $(SRC_PATH)/src
LOCAL_C_INCLUDES += $(SRC_PATH)/src/net
LOCAL_C_INCLUDES += $(SRC_PATH)/src/proto
LOCAL_C_INCLUDES += $(SRC_PATH)/src/utils

MN_LOCAL_SRC_FILES :=$(wildcard $(SRC_PATH)/src/*.c)
MN_LOCAL_SRC_FILES +=$(wildcard $(SRC_PATH)/src/proto/*.c)
MN_LOCAL_SRC_FILES +=$(wildcard $(SRC_PATH)/src/net/*.c)
MN_LOCAL_SRC_FILES +=$(wildcard $(SRC_PATH)/src/utils/*.c)


LOCAL_SRC_FILES := $(patsubst $(LOCAL_PATH)/%,%,$(MN_LOCAL_SRC_FILES))

## Static Librarys' Dependence
#LOCAL_STATIC_LIBRARIES += av_utils

include $(BUILD_SHARED_LIBRARY)

## Include Files
### cdnvister
#include $(VOICE_SRC_PATH)/cdnvister/build/Android/jni/Android.mk 


