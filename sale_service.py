import sale_pb2
import sale_pb2_grpc
from concurrent import futures
import time
import grpc

class CalculatorSales(sale_pb2_grpc.CalculatorSaleServicer):

    def GetSale(self, request, context):
        response = sale_pb2.Cost()
        response.cost = self._calculate_sale(request.cost)
        return response

    @staticmethod
    def _calculate_sale(value):
        return value * 0.5


if __name__ == "__main__":
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    port = 8001

    sale_pb2_grpc.add_CalculatorSaleServicer_to_server(CalculatorSales(), server)
    server.add_insecure_port(f'[::]:{port}')
    server.start()

    try:
        while True:
            time.sleep(800)
    except KeyboardInterrupt:
        server.stop(0)