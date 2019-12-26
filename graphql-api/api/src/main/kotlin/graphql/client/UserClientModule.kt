package graphql.client

import com.google.inject.AbstractModule
import grpc.user.UserServiceGrpc
import io.grpc.ManagedChannelBuilder

class UserClientModule : AbstractModule() {

    override fun configure() {
        var userGrpcServicePort: Int = System.getenv("USER_GRPC_SERVICE_PORT").toInt()
        var userGrpcService: String = System.getenv("USER_GRPC_SERVICE")

        val channel = ManagedChannelBuilder.forAddress(userGrpcService, userGrpcServicePort).usePlaintext(true).build()
        bind(UserServiceGrpc.UserServiceFutureStub::class.java).toInstance(UserServiceGrpc.newFutureStub(channel))
        bind(UserServiceGrpc.UserServiceBlockingStub::class.java).toInstance(UserServiceGrpc.newBlockingStub(channel))
    }
}