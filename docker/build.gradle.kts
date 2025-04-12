plugins {
    kotlin("jvm") version "2.0.20"
    application
}

group = "pl.edu.uj"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation(kotlin("test"))
    implementation("org.xerial:sqlite-jdbc:3.7.2")
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    jvmToolchain(8)
}

application {
    mainClass = "pl.edu.uj.MainKt"
}