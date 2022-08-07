package com.armory.logsort;

import com.armory.logsort.generator.LogGenerator;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.io.File;
import java.time.Instant;
import java.util.Arrays;
import java.util.List;

@SpringBootApplication
public class LogSortApplication {

    public static void main(String[] args) {
        // generate log files
        new LogGenerator(100, 100).generate("./src/main/resources/generated/logs/");

        // read log files
        final List<File> logFiles = Arrays.asList(new File("./src/main/resources/generated/logs/").listFiles());

        final Instant start = Instant.now();

        new LogPrinter().printLogs(logFiles);

        final Instant finish = Instant.now();
        System.out.println("\nLog sorting and printing took: " + (finish.toEpochMilli() - start.toEpochMilli()) + "ms");
    }

}
