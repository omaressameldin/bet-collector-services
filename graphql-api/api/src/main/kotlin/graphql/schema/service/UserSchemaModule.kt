package graphql.schema.service

import com.google.api.graphql.rejoiner.Query
import com.google.api.graphql.rejoiner.Mutation
import com.google.api.graphql.rejoiner.SchemaModule
import com.google.common.util.concurrent.ListenableFuture
import grpc.user.*
import io.github.vjames19.futures.jdk8.map

class UserSchemaModule : SchemaModule() {

    @Query("readUser")
    fun readUser(
      request: ReadRequest,
      client: UserServiceGrpc.UserServiceFutureStub
      ): ListenableFuture<ReadResponse> {
        return client.read(request)
    }

    @Mutation("loginUser")
    fun loginUser(
      request: LoginRequest,
      client: UserServiceGrpc.UserServiceFutureStub
      ): ListenableFuture<LoginResponse> {
        return client.login(request)
      }

      @Mutation("updateUser")
      fun updateUser(
        request: UpdateRequest,
        client: UserServiceGrpc.UserServiceFutureStub
      ): ListenableFuture<UpdateResponse> {
        return client.update(request)
      }

      @Mutation("deleteUser")
      fun deleteUser(
        request: DeleteRequest,
        client: UserServiceGrpc.UserServiceFutureStub
      ): ListenableFuture<DeleteResponse> {
        return client.delete(request)
      }
}