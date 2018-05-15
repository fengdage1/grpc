#pip3 install grpcio
#pip3 install grpcio-tools
#python3 -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. grpct.proto
#server代码在golang的src/project/grpctest/server/server.go里面，这里跨语言通信
import grpc_pb2_grpc
import grpc_pb2
import grpc
if "__main__"==__name__:
    conn=grpc.insecure_channel("127.0.0.1" + ':' + "50052")
    client = grpc_pb2_grpc.HelloStub(channel=conn)
    r=client.SayHello(grpc_pb2.HelloRequest(name="Gamelife"))
    print(r.message)
