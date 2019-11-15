package graphql.client

import com.google.inject.AbstractModule
import grpc.user.UserServiceGrpc
import io.grpc.ManagedChannelBuilder

class UserClientModule : AbstractModule() {

    override fun configure() {
        val channel = ManagedChannelBuilder.forAddress("bets-user-server", 7500).usePlaintext(true).build()
        bind(UserServiceGrpc.UserServiceFutureStub::class.java).toInstance(UserServiceGrpc.newFutureStub(channel))
        bind(UserServiceGrpc.UserServiceBlockingStub::class.java).toInstance(UserServiceGrpc.newBlockingStub(channel))
    }
}