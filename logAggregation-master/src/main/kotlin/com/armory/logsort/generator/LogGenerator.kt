package com.armory.logsort.generator

import java.io.File
import java.time.Instant
import kotlin.random.Random

class LogGenerator(val fileCount: Int, val avgLineCount: Long) {

    fun generate(filePath: String) {
        File(filePath).mkdirs()
        val files = (0..fileCount).map { File("${filePath}log_$it.log") }
        files.forEach { it.createNewFile() }

        var currentTs = Instant.ofEpochSecond(925543800L)
        val linesCountToGenerate = fileCount*avgLineCount

        var file = files.get(Random.nextInt(fileCount))
        println("Generating log files...")
        (1..(linesCountToGenerate)).forEach {
            if (Random.nextLong(3) == 0L) {
                file = files.get(Random.nextInt(fileCount))
            }
            val increase = Random.nextLong(100)
            currentTs = currentTs.plusSeconds(increase)
            file.appendText("${currentTs}, generic log text $it.\n")
            println("${(it.div(linesCountToGenerate.toFloat()))*100}%")
        }
        println("Log files generated.")
    }

}
