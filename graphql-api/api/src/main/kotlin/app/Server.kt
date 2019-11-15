package app

import graphql.GraphQLServer

class Server {

    private var graphql = GraphQLServer()

    fun start() {
        var port: Int = System.getenv("GRAPHQL_API_PORT").toInt()

        graphql.start(port)
    }
}

fun main(args: Array<String>) {
    Server().start()
}