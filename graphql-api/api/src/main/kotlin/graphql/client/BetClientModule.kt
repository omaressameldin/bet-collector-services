package graphql.client

import com.google.inject.AbstractModule
import grpc.bet.BetServiceGrpc
import io.grpc.ManagedChannelBuilder

class BetClientModule : AbstractModule() {

    override fun configure() {
        var betGrpcServicePort: Int = System.getenv("BET_GRPC_SERVICE_PORT").toInt()
        var betGrpcService: String = System.getenv("BET_GRPC_SERVICE")

        val channel = ManagedChannelBuilder.forAddress(betGrpcService, betGrpcServicePort).usePlaintext(true).build()
        bind(BetServiceGrpc.BetServiceFutureStub::class.java).toInstance(BetServiceGrpc.newFutureStub(channel))
        bind(BetServiceGrpc.BetServiceBlockingStub::class.java).toInstance(BetServiceGrpc.newBlockingStub(channel))
    }
}