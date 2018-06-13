import socket
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
result = sock.connect_ex(('127.0.0.1',3000))
if result == 0:
   print ("Port [3000] is not free")
else:
   print ("Port [3000] is free")
