## Wszystkie dema zawarte są w folderze [demos](./demos) w pliku demos.7z

## Zadanie 1 Docker

✅ 3.0 Obraz ubuntu z Pythonem w wersji 3.10

✅ 3.5 Obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem

✅ 4.0 Do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC SQLite w ramach projektu na Gradle (build.gradle)

✅ 4.5 Stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji przez CMD oraz gradle

✅ 5.0 Dodać konfigurację docker-compose [Link do commita](https://github.com/krystian-sikora/e-biznes/commit/4699dc800e69e4ffa85948f83c317308041244bb)

Kod: [docker](./docker)

Dockerhub: [dockerhub](https://hub.docker.com/r/ksikora7183/e-biznes/tags)


## Zadanie 2 Scala

✅ 3.0 Należy stworzyć kontroler do Produktów

✅ 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy

✅ 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy zgodnie z CRUD [Link do commita](https://github.com/krystian-sikora/e-biznes/commit/27f84297b79c799fc447f1ff45508d2a8cbdf833)

:x: 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać skrypt uruchamiający aplikację via ngrok (nie podawać tokena ngroka w gotowym rozwiązaniu)

:x: 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD


Kod: [scala](./scala)

## Zadanie 3 Kotlin

✅ 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord 

✅ 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota) 

✅ 4.0 Zwróci listę kategorii na określone żądanie użytkownika 

✅ 4.5 Zwróci listę produktów wg żądanej kategorii [Link do commita 4](https://github.com/krystian-sikora/e-biznes/commit/5d8b5a12fcdba1813bacf3a203d66ca84ddd7782)

:x: 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger, Webex


Kod: [kotlin](./ktor)
