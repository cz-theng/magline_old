APP_STL := gnustl_static
APP_CFLAGS := -fsigned-char
ifeq ($(NDK_DEBUG),1)
APP_CFLAGS 	 += -ggdb -O0
LOCAL_CFLAGS += -O0
APP_OPTIM := debug
else
APP_CFLAGS += -Os -ffreestanding
LOCAL_CPPFLAGS += -Os -ffreestanding 
LOCAL_CFLAGS += -Os -ffreestanding 
APP_CFLAGS += -D__RELEASE__ -DNDEBUG
APP_OPTIM := release
endif


APP_CPPFLAGS += -fexceptions -frtti
APP_ABI := armeabi armeabi-v7a
APP_SHORT_COMMANDS      := true
APP_CPPFLAGS += -Wno-error=format-security
APP_PLATFORM := android-17
