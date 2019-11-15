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
      ): ListenableFuture<User> {
        return client.read(request).map{ it.user }
    }

    @Mutation("createUser")
    fun createUser(
      request: CreateRequest,
      client: UserServiceGrpc.UserServiceFutureStub
      ): ListenableFuture<CreateResponse> {
        return client.create(request)
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