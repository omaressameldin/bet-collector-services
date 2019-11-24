package graphql.schema

import com.google.api.graphql.rejoiner.Schema
import com.google.api.graphql.rejoiner.SchemaProviderModule
import com.google.inject.Guice
import com.google.inject.Key
import graphql.client.*
import graphql.schema.service.*

object Schema {

    val schema: GraphQLSchema = Guice
            .createInjector(
                    SchemaProviderModule(),
                    UserClientModule(),
                    UserSchemaModule(),
                    BetClientModule(),
                    BetSchemaModule()
            )
            .getInstance(Key.get(GraphQLSchema::class.java, Schema::class.java))
}