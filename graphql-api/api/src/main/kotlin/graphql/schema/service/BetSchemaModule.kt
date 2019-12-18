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
      UserReadRequest.newBuilder().setId(bet.betterId).build(),
      client
    ).map { it.user }
  }

  @SchemaModification(addField = "accepter", onType = Bet::class)
  fun accepterToUser(bet: Bet, client: UserServiceGrpc.UserServiceFutureStub): ListenableFuture<User> {
    return UserSchemaModule().readUser(
      UserReadRequest.newBuilder().setId(bet.accepterId).build(),
      client
    ).map { it.user }
  }

    @Query("readBet")
    fun readBet(
      request: ReadRequest,
      client: BetServiceGrpc.BetServiceFutureStub
      ): ListenableFuture<ReadResponse> {
        return client.read(request)
    }

    @Query("readAllBets")
    fun readAllBet(
      request: ReadAllRequest,
      client: BetServiceGrpc.BetServiceFutureStub
      ): ListenableFuture<ReadAllResponse> {
        return client.readAll(request)
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