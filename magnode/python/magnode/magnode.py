__author__ = 'cz'


import socket
import struct
import time


class MagNode(object):
    def __init__(self):
        pass

    def connect(self, host, port):
        print("magnode connect with", host, port)
        addr = (host, port)
        sfd = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sfd.connect(addr)

        sfd.send(struct.pack("BBHII", 0x7f, 101, 0x01, 1, 6))
        time.sleep(1)
        sfd.send(struct.pack("HHH", 1, 2, 1))
        time.sleep(2)
        buf = sfd.recv(12)
        framehead = struct.unpack("BBHII", buf)
        print framehead
        time.sleep(1)
        buf = sfd.recv(framehead[4])
        ack  = struct.unpack("HH", buf)
        print ack

        sfd.send(struct.pack("BBHII", 0x7f, 101, 0x03, 2, 0))
        time.sleep(3)



