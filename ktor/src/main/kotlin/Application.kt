package pl.edu.uj

import dev.kord.core.Kord
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent
import io.github.cdimascio.dotenv.dotenv
import io.ktor.server.application.*
import kotlinx.coroutines.*
import kotlinx.coroutines.flow.any
import org.slf4j.LoggerFactory

private val logger = LoggerFactory.getLogger("Application")

data class Product(
    val name: String,
    val description: String,
    val category: String,
    val price: Double
)

private val products = listOf(
    Product("iPhone", "Telefon marki Apple", "Elektronika", 1000.0),
    Product("Never gonna give you up", "Piosenka autorstwa Ricka Astleya", "Muzyka", 30.0),
    Product("Jabłko", "Jabłko odmiany Granny Smith", "Owoce", 0.5)
)

private val categories = listOf(
    "Elektronika",
    "Muzyka",
    "Owoce"
)

@OptIn(ExperimentalCoroutinesApi::class, DelicateCoroutinesApi::class)
fun main(args: Array<String>): Unit = runBlocking {

    launch(newSingleThreadContext("BotThread")) {
        startBot()
    }

    launch { io.ktor.server.netty.EngineMain.main(args) }
}

suspend fun startBot() {
    logger.info("Starting Kord bot...")
    val token = loadToken()
    val bot = Kord(token)

    bot.on<MessageCreateEvent> {
        if (message.author?.isBot == true) return@on

        logger.info("Message from ${message.author?.username}: ${message.content}")

        val botMentioned = message.mentionedUsers.any { it.id == bot.selfId }

        if (botMentioned) {
            message.channel.createMessage("Hi there! You mentioned me? Say `!categories` to see available categories or `!category <category>` to see products in a specific category.")
        }

        if (message.content.startsWith("!categories")) {
            logger.info("Categories command received")
            message.channel.createMessage(categories.joinToString("\n") { "- $it" })
        }

        if (message.content.startsWith("!category")) {
            logger.info("Category command received")
            message.content.slice(10 until message.content.length).let { category ->
                val productsInCategory = products.filter { it.category == category }
                if (productsInCategory.isNotEmpty()) {
                    message.channel.createMessage(productsInCategory.joinToString("\n") { "- ${it.name}: ${it.description} (${it.price})" })
                } else {
                    message.channel.createMessage("No products found in category $category")
                }
            }
        }
    }

    bot.login {
        @OptIn(PrivilegedIntent::class)
        intents += Intent.MessageContent

        logger.info("Logged in as: ${bot.getSelf().username}")
    }
}

fun loadToken(): String {
    val dotenv = dotenv()
    return dotenv["DISCORD_TOKEN"] ?: error("DISCORD_TOKEN not found in .env")
}

fun Application.module() {
    configureRouting()
}
