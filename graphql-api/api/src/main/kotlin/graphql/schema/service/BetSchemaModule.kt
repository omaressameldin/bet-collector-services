package graphql.schema.service

import com.google.api.graphql.rejoiner.Query
import com.google.api.graphql.rejoiner.Mutation
import com.google.api.graphql.rejoiner.SchemaModification
import com.google.api.graphql.rejoiner.SchemaModule
import com.google.common.util.concurrent.ListenableFuture
import grpc.bet.*
import grpc.user.ReadRequest as UserReadRequest
import grpc.user.User
import grpc.user.UserServiceGrpc
import io.github.vjames19.futures.jdk8.map

class BetSchemaModule : SchemaModule() {

  @SchemaModification(addField = "better", onType = Bet::class)
  fun betterToUser(bet: Bet, client: UserServiceGrpc.UserServiceFutureStub): ListenableFuture<User> {
    return UserSchemaModule().readUser(
      UserReadRequest.newBuilder().setAuthId(bet.betterId).build(),
      client
    )
  }

  @SchemaModification(addField = "accepter", onType = Bet::class)
  fun accepterToUser(bet: Bet, client: UserServiceGrpc.UserServiceFutureStub): ListenableFuture<User> {
    return UserSchemaModule().readUser(
      UserReadRequest.newBuilder().setAuthId(bet.accepterId).build(),
      client
    )
  }

    @Query("readBet")
    fun readBet(
      request: ReadRequest,
      client: BetServiceGrpc.BetServiceFutureStub
      ): ListenableFuture<Bet> {
        return client.read(request).map{ it.bet }
    }

    @Mutation("createBet")
    fun createBet(
      request: CreateRequest,
      client: BetServiceGrpc.BetServiceFutureStub
      ): ListenableFuture<CreateResponse> {
        return client.create(request)
      }

      @Mutation("updateBet")
      fun updateBet(
        request: UpdateRequest,
        client: BetServiceGrpc.BetServiceFutureStub
      ): ListenableFuture<UpdateResponse> {
        return client.update(request)
      }

      @Mutation("deleteBet")
      fun deleteBet(
        request: DeleteRequest,
        client: BetServiceGrpc.BetServiceFutureStub
      ): ListenableFuture<DeleteResponse> {
        return client.delete(request)
      }
}