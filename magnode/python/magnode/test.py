__author__ = 'cz'
import magnode


def main():
    url = "tcp://127.0.0.1:8082"
    host = url.split(':')[1][2:]
    port = int(url.split(':')[2])
    node = magnode.MagNode()
    node.connect(host, port)

if __name__ == "__main__":
    main()