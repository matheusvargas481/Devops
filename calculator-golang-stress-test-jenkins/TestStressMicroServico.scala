package computerdatabase

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._
import scala.util.Random

class TestStressMicroServico extends Simulation {
    val count = Iterator.continually(Map("numeroUm" -> Random.nextInt(1000), "numeroDois" -> Random.nextInt(1000)))

    val httpConf = http
    .baseUrl("http://localhost:9000")

    val scan = scenario("Simple Stress Test")
            .feed(count)
            .exec(http("Http Request Sum")
            .get("/calc/sum/" + "${numeroUm}" + "/" + "${numeroDois}")
            .check(status.is(200))
            )
            .exec(http("Http Request Sub")
            .get("/calc/sub/" + "${numeroUm}" + "/" + "${numeroDois}")
            .check(status.is(200))
            )
            .exec(http("Http Request Mul")
            .get("/calc/mul/" + "${numeroUm}" + "/" + "${numeroDois}")
            .check(status.is(200))
            )
            .exec(http("Http Request Div")
            .get("/calc/div/" + "${numeroUm}" + "/" + "${numeroDois}")
            .check(status.is(200))
            )
            .exec(http("Http Request Historico")
            .get("/calc/historico")
            .check(status.is(200))
            )
    setUp(
        scan.inject(rampUsersPerSec(1) to(10) during(60 seconds))
    ).protocols(httpConf)
}